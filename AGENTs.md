# AGENTs Guide For valgo

This file is for LLMs and other automated contributors working in this repository.

## Purpose

`valgo` is a Go validation library built around small, typed validation pipelines. The public API is centered in `lib/v`, while reusable low-level predicate helpers live in `lib/is`.

If you need to understand how the library works, start with the public contracts in `lib/v/types.go` and then move outward into the typed pipes, actions, parser helpers, and tests.

## Core Features

These are the main library features and the most likely implementation locations:

- Typed validation pipes for `string`, `int`, `float64`, and `time.Time`: `lib/v/string_pipe.go`, `lib/v/int_pipe.go`, `lib/v/float_pipe.go`, `lib/v/time_pipe.go`
- Composable validation actions per type: `lib/v/string_actions.go`, `lib/v/int_actions.go`, `lib/v/float_actions.go`, `lib/v/time_actions.go`
- Schema validation across multiple fields: `lib/v/builder.go`, `lib/v/pipe_map.go`, `lib/v/parser.go`
- Builder-style schema composition with `Entry(...)`: `lib/v/builder.go`
- Parse-and-validate helpers for JSON payloads: `lib/v/parser.go`
- Structured errors for single-field, multi-field, and parse lifecycle failures: `lib/v/errors.go`
- Custom validators and custom pipes: `lib/v/string_actions.go`, `lib/v/int_actions.go`, `lib/v/float_actions.go`, `lib/v/time_actions.go`, `lib/v/customPipe.go`
- Custom error message support through `ErrMsg(...)`: `lib/v/builder.go`
- Lower-level reusable string/date/regex checks: `lib/is/string.go`, `lib/is/date_string.go`, `lib/is/regex.go`

## Repo Map

### Public API Surface

- `lib/v/types.go`: core interfaces such as `Schema`, `PipeSet`, `PipeFace`, and option/error-message contracts
- `lib/v/builder.go`: schema builders, `Entry(...)`, `NewPipesBuilder(...)`, `NewPipesMap(...)`, and `ErrMsg(...)`
- `lib/v/parser.go`: `Parse`, `ParseFull`, `ParseBytes`, `ParseBytesFull`, `Validate`, and `ValidateAll`
- `lib/v/errors.go`: `PipeError`, `ValidationErrors`, and `ParseError`

### Pipe Implementations

- `lib/v/string_pipe.go`: executes string actions sequentially until first failure
- `lib/v/int_pipe.go`: executes int actions sequentially until first failure
- `lib/v/float_pipe.go`: executes float actions sequentially until first failure
- `lib/v/time_pipe.go`: executes time actions sequentially until first failure
- `lib/v/customPipe.go`: generic custom pipe for user-defined validation functions

### Validator Families

- `lib/v/string_actions.go`: string validators including content, regex, format, path, IP, UUID, RFC date/time string, and custom string checks
- `lib/v/int_actions.go`: numeric validators such as min/max, sign, bounds, custom integer rules
- `lib/v/float_actions.go`: float validators such as min/max, sign, bounds, custom float rules
- `lib/v/time_actions.go`: time validators such as before/after, same day/month/year, age checks, weekday, timezone checks

### Helper Layer

- `lib/is/string.go`: helper predicates used by public string validators
- `lib/is/date_string.go`: date and time string helpers
- `lib/is/regex.go`: shared regex patterns backing validators

### Tests And Behavior Locks

- `tests/string_test.go`: strongest coverage for string validators and action composition
- `tests/parser_test.go`: parse-and-validate flow from bytes and readers
- `tests/errors_system_test.go`: error formatting, wrapping, and `errors.Is` / `errors.As` behavior
- `tests/customPipe_test.go`: custom pipe behavior in schema parsing flow
- `tests/validate_all_bench_test.go`: benchmarks around schema validation

### Non-Critical / Orientation Files

- `README.md`: high-level feature list and examples
- `cmd/main.go`: usage/demo entrypoint, useful for quick API shape only
- `internal/pipe/`: currently empty

## Mental Model

The repository is easiest to understand in this order:

1. A user type implements `Rules() (PipeSet, error)`.
2. A `PipeSet` groups multiple `PipeFace` values.
3. Each `PipeFace` validates one typed value by running actions in order.
4. Validation stops at the first failing action inside a pipe.
5. Schema validation either returns the first error or aggregates all field errors.
6. Parse helpers decode JSON first, build rules second, and validate last.

In practice, most end-user behavior is defined by `lib/v`, not `lib/is`.

## Where To Start For Common Tasks

### I need to add a new validator

1. Find the correct type family in `lib/v/*_actions.go`.
2. If the logic is reusable and low-level, add a helper in `lib/is/*`.
3. Keep the exported user-facing constructor in `lib/v`.
4. Add tests in the nearest relevant test file.

### I need to trace a validation failure

1. Check the typed pipe in `lib/v/*_pipe.go`.
2. Inspect the action constructor in `lib/v/*_actions.go`.
3. If it delegates to helper logic, inspect `lib/is/*`.
4. If the error is schema-wide or parse-related, inspect `lib/v/errors.go` and `lib/v/parser.go`.

### I need to understand schema construction

1. Read `lib/v/types.go` for contracts.
2. Read `lib/v/builder.go` for `Entry(...)`, `NewPipesBuilder(...)`, and `NewPipesMap(...)`.
3. Read `lib/v/pipe_map.go` for direct map-backed validation semantics.

## Best-Practice Snippets

### Snippet: Add A New Public Validator

```text
Preferred path:
1. Add the exported constructor in lib/v/<type>_actions.go
2. Reuse lib/is helpers if the predicate can be shared
3. Return errors in the same style as existing validators
4. Add tests before touching README examples
```

### Snippet: Choose The Right Test File

```text
String validator change      -> tests/string_test.go
Parse lifecycle change       -> tests/parser_test.go
Error wrapping/message change -> tests/errors_system_test.go
Custom pipe change           -> tests/customPipe_test.go
Performance investigation    -> tests/validate_all_bench_test.go
```

### Snippet: Preferred Reading Order

```text
lib/v/types.go
lib/v/builder.go
lib/v/<type>_pipe.go
lib/v/<type>_actions.go
lib/v/parser.go
lib/v/errors.go
tests/*.go
```

### Snippet: When Docs And Code Disagree

```text
Trust order:
1. lib/v source
2. tests
3. lib/is helpers
4. README examples
5. cmd/main.go demo
```

## Important Repo-Specific Behaviors

- Validation inside a pipe is sequential and short-circuits on first failure.
- `Validate()` returns the first schema error it encounters.
- `ValidateAll()` accumulates all field errors into `ValidationErrors`.
- `PipeMap` iteration order follows Go map semantics, so field order is not stable when map-backed validation is used.
- `NewPipesBuilder(...)` is preferable when deterministic pipe order matters.
- Parse helpers wrap failures into `ParseError` rather than returning raw decode/validation errors directly.
- Embedding `Include` in a schema allows parse helpers to work even when validation rules are intentionally omitted.

## Known Gotchas

- `ValidateAllParallel` currently behaves the same as sequential validation. Do not assume real parallel execution.
- `Pattern(...)` uses `regexp.MustCompile`, so invalid regex input can panic at construction time.
- Some README wording is broader than the actual tests. When uncertain, verify behavior in source and tests.
- `cmd/main.go` is only a demo. It is not the authoritative specification of library behavior.

## Editing Guidance For Future LLMs

- Prefer minimal edits in `lib/v` unless the task is explicitly about helper predicates.
- Preserve the existing public API shape unless the request is specifically a breaking-change task.
- Add or update tests with every behavior change.
- If you are debugging a validator bug, inspect both the action constructor and any helper it delegates to.
- If you are changing parse or error behavior, validate that `errors.Is` and `errors.As` semantics still make sense.
- Do not treat `README.md` as the source of truth when implementation or tests disagree.

## Quick Orientation Summary

If another LLM has only a minute to orient itself, the shortest reliable path is:

1. Read `lib/v/types.go`.
2. Read `lib/v/builder.go` and `lib/v/parser.go`.
3. Open the relevant typed action file in `lib/v`.
4. Use `tests/string_test.go`, `tests/parser_test.go`, and `tests/errors_system_test.go` as behavior references.
5. Only then consult `README.md` for user-facing examples.
