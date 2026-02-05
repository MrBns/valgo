package tests_test

import (
	"testing"

	"github.com/mrbns/valgo/lib/v"
)

func TestEmail(t *testing.T) {
	err := v.IsEmail().Run("hi@mrbns.dev")
	if err != nil {
		t.Error(err)
	}

	err = v.IsEmail().Run("invalid-email")
	if err == nil {
		t.Errorf("invalid email should throws error")
	}
}

func TestLength(t *testing.T) {
	err := v.MinLength(5).Run("hello there")
	if err != nil {
		t.Error(err)
	}

	err = v.MinLength(50).Run("invalid text length")
	if err == nil {
		t.Errorf("should throw error with invalid lenght")
	}

	err = v.MaxLength(50).Run("hello there")
	if err != nil {
		t.Error(err)
	}

	err = v.MaxLength(5).Run("invalid maximum content")
	if err == nil {
		t.Errorf("should throw for out of max length")
	}

}

func TestNotEmpty(t *testing.T) {
	err := v.NotEmpty().Run("hello")
	if err != nil {
		t.Error(err)
	}

	err = v.NotEmpty().Run("")
	if err == nil {
		t.Errorf("empty string should throw error")
	}
}

func TestPattern(t *testing.T) {
	err := v.Pattern(`^[a-z]+$`).Run("hello")
	if err != nil {
		t.Error(err)
	}

	err = v.Pattern(`^[a-z]+$`).Run("hello123")
	if err == nil {
		t.Errorf("pattern mismatch should throw error")
	}

	err = v.Pattern(`^\d{3}-\d{2}-\d{4}$`).Run("123-45-6789")
	if err != nil {
		t.Error(err)
	}
}

func TestEnum(t *testing.T) {
	allowedStatuses := []string{"active", "inactive", "pending"}

	err := v.Enum(allowedStatuses).Run("active")
	if err != nil {
		t.Error(err)
	}

	err = v.Enum(allowedStatuses).Run("rejected")
	if err == nil {
		t.Errorf("invalid enum value should throw error")
	}
}

func TestHasPrefix(t *testing.T) {
	err := v.HasPrefix("test_").Run("test_function")
	if err != nil {
		t.Error(err)
	}

	err = v.HasPrefix("test_").Run("function_test")
	if err == nil {
		t.Errorf("wrong prefix should throw error")
	}
}

func TestHasSuffix(t *testing.T) {
	err := v.HasSuffix("_test").Run("function_test")
	if err != nil {
		t.Error(err)
	}

	err = v.HasSuffix("_test").Run("test_function")
	if err == nil {
		t.Errorf("wrong suffix should throw error")
	}
}

func TestEqualFold(t *testing.T) {
	err := v.EqualFold("Hello").Run("hello")
	if err != nil {
		t.Error(err)
	}

	err = v.EqualFold("Hello").Run("HELLO")
	if err != nil {
		t.Error(err)
	}

	err = v.EqualFold("Hello").Run("world")
	if err == nil {
		t.Errorf("unequal strings should throw error")
	}
}

func TestContains(t *testing.T) {
	err := v.Contains("world").Run("hello world")
	if err != nil {
		t.Error(err)
	}

	err = v.Contains("world").Run("hello earth")
	if err == nil {
		t.Errorf("substring not found should throw error")
	}
}

func TestIsAlpha(t *testing.T) {
	err := v.IsAlpha().Run("abcDEF")
	if err != nil {
		t.Error(err)
	}

	err = v.IsAlpha().Run("abc123")
	if err == nil {
		t.Errorf("alphanumeric should throw error")
	}

	err = v.IsAlpha().Run("abc-def")
	if err == nil {
		t.Errorf("string with special chars should throw error")
	}
}

func TestIsAlphaNumeric(t *testing.T) {
	err := v.IsAlphaNumeric().Run("abcDEF123")
	if err != nil {
		t.Error(err)
	}

	err = v.IsAlphaNumeric().Run("abc-123")
	if err == nil {
		t.Errorf("string with special chars should throw error")
	}
}

func TestIsAscii(t *testing.T) {
	err := v.IsAscii().Run("hello123!@#")
	if err != nil {
		t.Error(err)
	}

	err = v.IsAscii().Run("hello€world")
	if err == nil {
		t.Errorf("non-ascii chars should throw error")
	}
}

func TestIsBase64(t *testing.T) {
	err := v.IsBase64().Run("aGVsbG8gd29ybGQ=")
	if err != nil {
		t.Error(err)
	}

	err = v.IsBase64().Run("not_valid_base64!!!")
	if err == nil {
		t.Errorf("invalid base64 should throw error")
	}
}

func TestIsDate(t *testing.T) {
	err := v.IsDate().Run("2024-01-15")
	if err != nil {
		t.Error(err)
	}

	err = v.IsDate().Run("01/15/2024")
	if err == nil {
		t.Errorf("invalid date format should throw error")
	}

	err = v.IsDate().Run("2024-13-01")
	if err == nil {
		t.Errorf("invalid month should throw error")
	}
}

func TestIsDecimal(t *testing.T) {
	err := v.IsDecimal().Run("123.45")
	if err != nil {
		t.Error(err)
	}

	err = v.IsDecimal().Run("-456.78")
	if err != nil {
		t.Error(err)
	}

	err = v.IsDecimal().Run("abc.def")
	if err == nil {
		t.Errorf("non-numeric string should throw error")
	}
}

func TestIsHexColor(t *testing.T) {
	err := v.IsHexColor().Run("#FF5733")
	if err != nil {
		t.Error(err)
	}

	err = v.IsHexColor().Run("#fff")
	if err != nil {
		t.Error(err)
	}

	err = v.IsHexColor().Run("FF5733")
	if err == nil {
		t.Errorf("hex without # should throw error")
	}

	err = v.IsHexColor().Run("#GGGGGG")
	if err == nil {
		t.Errorf("invalid hex chars should throw error")
	}
}

func TestIsHexDecimal(t *testing.T) {
	err := v.IsHexDecimal().Run("ABCDEF123")
	if err != nil {
		t.Error(err)
	}

	err = v.IsHexDecimal().Run("0x1A2B")
	if err != nil {
		t.Error(err)
	}

	err = v.IsHexDecimal().Run("0xFF")
	if err != nil {
		t.Error(err)
	}

	err = v.IsHexDecimal().Run("0x")
	if err == nil {
		t.Errorf("0x without hex digits should throw error")
	}

	err = v.IsHexDecimal().Run("0xGGG")
	if err == nil {
		t.Errorf("invalid hex chars after 0x should throw error")
	}
}

func TestIsIPV4(t *testing.T) {
	err := v.IsIPV4().Run("192.168.1.1")
	if err != nil {
		t.Error(err)
	}

	err = v.IsIPV4().Run("256.168.1.1")
	if err == nil {
		t.Errorf("invalid IPv4 should throw error")
	}

	err = v.IsIPV4().Run("192.168.1")
	if err == nil {
		t.Errorf("incomplete IPv4 should throw error")
	}
}

func TestIsIPV6(t *testing.T) {
	err := v.IsIPV6().Run("2001:0db8:85a3:0000:0000:8a2e:0370:7334")
	if err != nil {
		t.Error(err)
	}

	err = v.IsIPV6().Run("::1")
	if err != nil {
		t.Error(err)
	}

	err = v.IsIPV6().Run("192.168.1.1")
	if err == nil {
		t.Errorf("IPv4 should not pass IPv6 validation")
	}
}

func TestIsJSON(t *testing.T) {
	err := v.IsJSON().Run(`{"name":"John","age":30}`)
	if err != nil {
		t.Error(err)
	}

	err = v.IsJSON().Run(`["item1","item2"]`)
	if err != nil {
		t.Error(err)
	}

	err = v.IsJSON().Run(`{invalid json}`)
	if err == nil {
		t.Errorf("invalid JSON should throw error")
	}
}

func TestIsURL(t *testing.T) {
	err := v.IsURL().Run("https://example.com")
	if err != nil {
		t.Error(err)
	}

	err = v.IsURL().Run("http://localhost:8080/path")
	if err != nil {
		t.Error(err)
	}

	err = v.IsURL().Run("not a url")
	if err == nil {
		t.Errorf("invalid URL should throw error")
	}
}

func TestIsUUID(t *testing.T) {
	err := v.IsUUID().Run("550e8400-e29b-41d4-a716-446655440000")
	if err != nil {
		t.Error(err)
	}

	err = v.IsUUID().Run("not-a-uuid")
	if err == nil {
		t.Errorf("invalid UUID should throw error")
	}
}

func TestIsValidPort(t *testing.T) {
	err := v.IsValidPort().Run("8080")
	if err != nil {
		t.Error(err)
	}

	err = v.IsValidPort().Run("443")
	if err != nil {
		t.Error(err)
	}

	err = v.IsValidPort().Run("65535")
	if err != nil {
		t.Error(err)
	}

	err = v.IsValidPort().Run("65536")
	if err == nil {
		t.Errorf("port > 65535 should throw error")
	}

	err = v.IsValidPort().Run("abc")
	if err == nil {
		t.Errorf("non-numeric port should throw error")
	}
}

func TestCustomString(t *testing.T) {
	customValidator := v.CustomString(func(value string) bool {
		return len(value) > 5 && len(value) < 20
	})

	err := customValidator.Run("hello world")
	if err != nil {
		t.Error(err)
	}

	err = customValidator.Run("hi")
	if err == nil {
		t.Errorf("short string should throw error")
	}

	err = customValidator.Run("this is a very long string that exceeds limit")
	if err == nil {
		t.Errorf("long string should throw error")
	}
}

func TestStringPipeMultipleActions(t *testing.T) {
	pipe := v.StringPipe("user@example.com", v.NotEmpty(), v.IsEmail())
	err := pipe.Validate()
	if err != nil {
		t.Error(err)
	}

	pipe = v.StringPipe("", v.NotEmpty(), v.IsEmail())
	err = pipe.Validate()
	if err == nil {
		t.Errorf("empty string should fail validation")
	}

	pipe = v.StringPipe("invalid-email", v.NotEmpty(), v.IsEmail())
	err = pipe.Validate()
	if err == nil {
		t.Errorf("invalid email should fail validation")
	}
}

func TestStringPipeWithKey(t *testing.T) {
	pipe := v.StringPipe("test123", v.IsAlpha())
	pipe.Validate()

	if pipe.Key() == "" {
		t.Log("Key is initially empty - expected behavior")
	}
}

func TestStringPipeComplexValidation(t *testing.T) {
	pipe := v.StringPipe("test_user_123", v.NotEmpty(), v.MinLength(8), v.MaxLength(30), v.Contains("user"))
	err := pipe.Validate()
	if err != nil {
		t.Error(err)
	}

	pipe = v.StringPipe("hi", v.NotEmpty(), v.MinLength(8), v.MaxLength(30))
	err = pipe.Validate()
	if err == nil {
		t.Errorf("string too short should fail validation")
	}
}

func TestStringPipeEarlyExit(t *testing.T) {
	pipe := v.StringPipe("", v.NotEmpty(), v.IsEmail())
	err := pipe.Validate()
	if err == nil {
		t.Errorf("should fail on first action (NotEmpty)")
	}
	if err != nil && err.Key != "" {
		// Key should be set if setKey was called
		t.Logf("Error key: %s", err.Key)
	}
}

func TestIsHTML(t *testing.T) {
	err := v.IsHTML().Run("<html><body>test</body></html>")
	if err != nil {
		t.Error(err)
	}

	err = v.IsHTML().Run("plain text")
	if err == nil {
		t.Errorf("plain text should throw error")
	}
}

func TestIsRGB(t *testing.T) {
	err := v.IsRGB().Run("rgb(255, 100, 50)")
	if err != nil {
		t.Error(err)
	}

	err = v.IsRGB().Run("rgb(256, 100, 50)")
	if err == nil {
		t.Errorf("invalid RGB values should throw error")
	}

	err = v.IsRGB().Run("rgba(255, 100, 50, 50)")
	if err != nil {
		t.Error(err)
	}

	err = v.IsRGB().Run("rgba(255%, 100%, 50%, 0.5)")
	if err != nil {
		t.Error(err)
	}

	err = v.IsRGB().Run("rgba(255, 100, 500, 2.3)")
	if err == nil {
		t.Errorf("invalid alpha channel throw error")
	}

	err = v.IsRGB().Run("rgba(256, 100, 400, 101)")
	if err == nil {
		t.Errorf("invalid rgba should throw errors")
	}

}

func TestIsHSL(t *testing.T) {
	err := v.IsHSL().Run("hsl(200, 50%, 50%)")
	if err != nil {
		t.Error(err)
	}

	err = v.IsHSL().Run("hsl(400, 50%, 50%")
	if err == nil {
		t.Errorf("invalid HSL values should throw error")
	}
}

func TestEdgeCases(t *testing.T) {
	tests := []struct {
		name      string
		action    v.StringPipeAction
		input     string
		shouldErr bool
	}{
		{"Empty string with NotEmpty", v.NotEmpty(), "", true},
		{"Single char with MinLength 1", v.MinLength(1), "a", false},
		{"Unicode with IsAlpha", v.IsAlpha(), "café", false},
		{"Zero length MaxLength", v.MaxLength(0), "", false},
		{"Pattern empty string", v.Pattern(`^.*$`), "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.action.Run(tt.input)
			if (err != nil) != tt.shouldErr {
				t.Errorf("Expected error: %v, got error: %v", tt.shouldErr, err != nil)
			}
		})
	}
}
