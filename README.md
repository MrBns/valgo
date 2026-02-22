# Valigo - Simple Validation Library for Go

A flexible and easy-to-use validation framework for Go that supports strings, integers, floats, and time values through a clean pipeline pattern.

## ✨ Features

- **Simple API** - Build validations with small, composable actions
- **Pipeline Pattern** - Chain multiple checks in a single pipe
- **Type Support** - String, int, float64, and time.Time validators
- **Schema Validation** - Validate multiple fields together
- **Custom Validators** - Define your own rules per type
- **Custom Error Messages** - Use `ErrMsg("...")` with `{VALUE}` placeholder
- **Parse + Validate** - Decode JSON and validate in one step
- **Structured Errors** - `PipeError`, `ValidationErrors`, and `ParseError`

## 📦 Installation

```bash
go get github.com/mrbns/valgo
```

## 🚀 Quick Start

```go
package main

import (
	"fmt"

	"github.com/mrbns/valgo/lib/v"
)

func main() {
	pipe := v.StringPipe("user@example.com", v.NotEmpty(), v.IsEmail())

	if err := pipe.Validate(); err != nil {
		fmt.Println("Validation failed:", err)
		return
	}

	fmt.Println("Email is valid!")
}
```

## 📖 Basic Concepts

### Pipe
A **pipe** validates one value by running actions in order.

### Action
An **action** is a single rule (for example, `IsEmail()` or `Min(18)`).

### Schema
A **schema** groups multiple pipes and validates a struct or payload.

## 💡 Usage Examples

### Validate a Single String

```go
pipe := v.StringPipe(
	"test@example.com",
	v.NotEmpty(),
	v.IsEmail(),
	v.MaxLength(50),
)

if err := pipe.Validate(); err != nil {
	fmt.Println(err)
}
```

### Validate a Number

```go
pipe := v.IntPipe(
	25,
	v.Min(18),
	v.Max(100),
	v.IsPositive(),
)

if err := pipe.Validate(); err != nil {
	fmt.Println(err)
}
```

### Validate a Float

```go
pipe := v.FloatPipe(
	99.99,
	v.MinFloat(0.0),
	v.MaxFloat(1000.0),
)

if err := pipe.Validate(); err != nil {
	fmt.Println(err)
}
```

### Validate Time

```go
pipe := v.TimePipe(
	time.Now().Add(-2*time.Hour),
	v.BeforeNow(),
	v.IsWeekday(),
)

if err := pipe.Validate(); err != nil {
	fmt.Println(err)
}
```

### Validate Multiple Fields (Schema)

```go
schema := v.NewPipesMap(v.PipeMap{
	"email": v.StringPipe("user@example.com", v.NotEmpty(), v.IsEmail()),
	"age":   v.IntPipe(25, v.Min(18)),
	"price": v.FloatPipe(19.99, v.MinFloat(0)),
})

// Stop at first error
if err := schema.Validate(); err != nil {
	fmt.Println(err)
}

// Return all field errors
if err := schema.ValidateAll(); err != nil {
	if errs, ok := err.(v.ValidationErrors); ok {
		for _, fieldErr := range errs {
			fmt.Printf("%s: %v\n", fieldErr.Key, fieldErr.Err)
		}
	}
}
```

### Build Schema with Entries

```go
schema := v.NewPipesBuilder(
	v.Entry("email").StringPipe("user@example.com", v.NotEmpty(), v.IsEmail()),
	v.Entry("age").IntPipe(20, v.Min(18)),
)
```

### Parse JSON and Validate in One Step

```go
type UserSchema struct {
	v.Include
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u *UserSchema) Rules() (v.PipeSet, error) {
	return v.NewPipesMap(v.PipeMap{
		"name": v.StringPipe(u.Name, v.NotEmpty(), v.MaxLength(40)),
		"age":  v.IntPipe(u.Age, v.NonZero(), v.Min(18)),
	}), nil
}

func parseAndValidate(data []byte) error {
	var payload UserSchema
	return v.ParseBytesFull(data, &payload)
}
```

### Custom Error Messages

```go
pipe := v.IntPipe(
	15,
	v.Min(18, v.ErrMsg("You must be at least 18, but got {VALUE}")),
	v.Max(65, v.ErrMsg("{VALUE} exceeds maximum age of 65")),
)
```

### Custom Validators

```go
nameRule := v.CustomString(func(value string) bool {
	return strings.HasPrefix(value, "user_")
}, v.ErrMsg("username must start with user_"))

evenRule := v.CustomNumber(func(value int) bool {
	return value%2 == 0
}, v.ErrMsg("{VALUE} must be an even number"))

_ = v.StringPipe("user_john", nameRule)
_ = v.IntPipe(10, evenRule)
```

### Custom Pipe

```go
pipe := v.CustomPipe("secret", func(value string) error {
	if value == "" {
		return errors.New("cannot be empty")
	}
	return nil
})

schema := v.NewPipesMap(v.PipeMap{"token": pipe})
```

## ⚠️ Error Types

- `*v.PipeError` - Single field validation error (`key` + `error`)
- `v.ValidationErrors` - Multiple field errors from `ValidateAll()`
- `*v.ParseError` - Parse/Rules/Validation lifecycle errors from Parse helpers

All error types support `errors.Is` / `errors.As` through `Unwrap`.

## 📋 Available String Validators

| Validator | Description |
|-----------|-------------|
| `NotEmpty()` | String must not be empty |
| `Enum([]string)` | Value must be one of allowed values |
| `Pattern(regex)` | Value must match regex |
| `MinLength(n)` | Minimum length |
| `MaxLength(n)` | Maximum length |
| `HasPrefix(prefix)` | Must start with prefix |
| `HasSuffix(suffix)` | Must end with suffix |
| `EqualFold(target)` | Case-insensitive equality |
| `Contains(substr)` | Must contain substring |
| `CustomString(fn)` | Custom string validator |
| `IsAlpha()` | Letters only |
| `IsAlphaNumeric()` | Letters and numbers only |
| `IsAscii()` | ASCII only |
| `IsBase32()` | Base32 format |
| `IsBase58()` | Base58 format |
| `IsBase64()` | Base64 format |
| `IsBitcoinAddress()` | Bitcoin address |
| `IsCreditCard()` | Credit card number |
| `IsDate()` | Date string |
| `IsDataURI()` | Data URI format |
| `IsDecimal()` | Decimal number string |
| `IsEmail()` | Email format |
| `IsEvmAddress()` | Ethereum address |
| `IsHTML()` | HTML content |
| `IsHexColor()` | Hex color format |
| `IsHexDecimal()` | Hexadecimal format |
| `IsHSL()` | HSL color format |
| `IsIPV4()` | IPv4 address |
| `IsIPV6()` | IPv6 address |
| `IsJSON()` | JSON string |
| `IsRGB()` | RGB/RGBA format |
| `IsULID()` | ULID format |
| `IsURL()` | URL format |
| `IsUUID()` | UUID (any supported version) |
| `IsUUIDV1()` | UUID v1 |
| `IsUUIDV3()` | UUID v3 |
| `IsUUIDV4()` | UUID v4 |
| `IsUUIDV5()` | UUID v5 |
| `IsValidPath()` | Filesystem path |
| `IsValidPort()` | Network port (0-65535) |
| `IsXML()` | XML string |
| `IsANSIC()` | ANSI C date/time format |
| `IsUnixDate()` | UnixDate format |
| `IsRubyDate()` | RubyDate format |
| `IsRFC822()` | RFC822 date format |
| `IsRFC822Z()` | RFC822Z date format |
| `IsRFC850()` | RFC850 date format |
| `IsRFC1123()` | RFC1123 date format |
| `IsRFC1123Z()` | RFC1123Z date format |
| `IsRFC3339()` | RFC3339 date format |
| `IsRFC3339Nano()` | RFC3339Nano date format |
| `IsKitchen()` | Kitchen time format |
| `IsStamp()` | Stamp format |
| `IsStampMilli()` | StampMilli format |
| `IsStampMicro()` | StampMicro format |
| `IsStampNano()` | StampNano format |
| `IsDateTime()` | DateTime format |
| `IsTimeOnly()` | TimeOnly format |

## 🔢 Available Integer Validators

| Validator | Description |
|-----------|-------------|
| `CustomNumber(fn)` | Custom int validator |
| `Min(n)` | Value must be `>= n` |
| `Max(n)` | Value must be `<= n` |
| `Gt(n)` | Value must be `> n` |
| `Gte(n)` | Value must be `>= n` |
| `Lt(n)` | Value must be `< n` |
| `Lte(n)` | Value must be `<= n` |
| `IsPositive()` | Value must be `> 0` |
| `IsNegative()` | Value must be `< 0` |
| `NonZero()` | Value must be `!= 0` |
| `IsIntString()` | Placeholder validator (currently always true) |

## 📊 Available Float Validators

| Validator | Description |
|-----------|-------------|
| `CustomFloat(fn)` | Custom float validator |
| `MinFloat(n)` | Value must be `>= n` |
| `MaxFloat(n)` | Value must be `<= n` |
| `GtFloat(n)` | Value must be `> n` |
| `GteFloat(n)` | Value must be `>= n` |
| `LtFloat(n)` | Value must be `< n` |
| `LteFloat(n)` | Value must be `<= n` |
| `IsPositiveFloat()` | Value must be `>= 0` |
| `IsNegativeFloat()` | Value must be `< 0` |

## 🕒 Available Time Validators

| Validator | Description |
|-----------|-------------|
| `CustomTime(fn)` | Custom time validator |
| `Before(t)` | Must be before `t` |
| `After(t)` | Must be after `t` |
| `Between(start, end)` | Must be strictly between start and end |
| `BeforeNow()` | Must be in the past |
| `AfterNow()` | Must be in the future |
| `NotEmptyDate()` | Must not be zero `time.Time` |
| `SameDay(t)` | Same year/month/day as `t` |
| `SameMonth(t)` | Same year/month as `t` |
| `SameYear(t)` | Same year as `t` |
| `MinDate(t)` | Must be on or after `t` |
| `MaxDate(t)` | Must be on or before `t` |
| `Equal(t)` | Must equal `t` |
| `NotEqual(t)` | Must not equal `t` |
| `OldOfDays(n)` | Must be at least `n` days old |
| `OldOf(d)` | Must be at least duration `d` old |
| `NewOf(n)` | Must be at least `n` days in the future |
| `SameWeek(t)` | Same ISO week as `t` |
| `IsWeekday()` | Must be Monday-Friday |
| `IsTimezone()` | Must have timezone offset in valid range |

## 🧠 Validation Helpers

- `v.Validate(schema)` - Validate schema and stop at first error
- `v.ValidateAll(schema)` - Validate schema and return all errors
- `v.ValidateAllParallel(schema)` - Current behavior matches `ValidateAll` (sequential)

## 📝 Notes

- `Validate()` stops on the first failed action in each pipe.
- `ValidateAll()` collects all failed fields and returns `v.ValidationErrors`.
- `PipeMap` iteration order is Go map iteration order.

## 🤝 Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## 📚 Source Reference

For more detailed information about each validator, check:
- [`lib/v/string_actions.go`](lib/v/string_actions.go) - String validators
- [`lib/v/int_actions.go`](lib/v/int_actions.go) - Integer validators
- [`lib/v/float_actions.go`](lib/v/float_actions.go) - Float validators
- [`lib/v/time_actions.go`](lib/v/time_actions.go) - Time validators
- [`lib/v/parser.go`](lib/v/parser.go) - Parse and schema validation flow
- [`lib/v/errors.go`](lib/v/errors.go) - Error types
- [`lib/is/string.go`](lib/is/string.go) - Low-level validation functions

## 📝 License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## author
Mr Binary Sniper. @mrbns.