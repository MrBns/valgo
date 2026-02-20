package tests_test

import (
	"testing"

	"github.com/mrbns/valgo/lib/v"
)

func buildBenchIntPipeSet(fieldCount int) v.PipeSet {
	pipes := make([]v.PipeFace, 0, fieldCount)

	for i := 0; i < fieldCount; i++ {
		value := i + 100
		pipes = append(pipes, v.IntPipe(
			value,
			v.IsPositive(),
			v.Gt(0),
			v.Gte(10),
			v.Min(10),
			v.Max(1000000),
			v.Lte(1000000),
			v.Lt(2000000),
		))
	}

	return v.NewPipesBuilder(pipes...)
}

func buildBenchStringPipeSet(fieldCount int) v.PipeSet {
	pipes := make([]v.PipeFace, 0, fieldCount)
	value := "bench.user+valgo@example.com"

	for i := 0; i < fieldCount; i++ {
		pipes = append(pipes, v.StringPipe(
			value,
			v.NotEmpty(),
			v.MinLength(10),
			v.MaxLength(64),
			v.Contains("@"),
			v.HasSuffix(".com"),
			v.IsEmail(),
		))
	}

	return v.NewPipesBuilder(pipes...)
}

func benchmarkValidateAllIntSerial(b *testing.B, fieldCount int) {
	b.ReportAllocs()
	pipeSet := buildBenchIntPipeSet(fieldCount)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		errs := pipeSet.ValidateAll()
		if errs != nil {
			b.Fatalf("expected nil errors, got %d", len(errs.Errors))
		}
	}
}

func benchmarkValidateAllStringSerial(b *testing.B, fieldCount int) {
	b.ReportAllocs()
	pipeSet := buildBenchStringPipeSet(fieldCount)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		errs := pipeSet.ValidateAll()
		if errs != nil {
			b.Fatalf("expected nil errors, got %d", len(errs.Errors))
		}
	}
}

func BenchmarkValidateAllIntSerialSmall(b *testing.B) {
	benchmarkValidateAllIntSerial(b, 8)
}

func BenchmarkValidateAllIntSerialLarge(b *testing.B) {
	benchmarkValidateAllIntSerial(b, 512)
}

func BenchmarkValidateAllStringSerialSmall(b *testing.B) {
	benchmarkValidateAllStringSerial(b, 8)
}

func BenchmarkValidateAllStringSerialLarge(b *testing.B) {
	benchmarkValidateAllStringSerial(b, 512)
}
