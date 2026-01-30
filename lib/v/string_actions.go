package v

import (
	"fmt"
	"net/mail"
	"regexp"
	"slices"
	"strings"

	"github.com/mrbns/valgo/lib/is"
)

// stringAction implements StringPipeAction for string validation.
type stringAction struct {
	errorMsg func(v string) string
	validate func(v string) bool
}

// Run executes the validation function on the given string value.
// Returns an error if validation fails.
func (action *stringAction) Run(value string) error {
	if !action.validate(value) {
		return fmt.Errorf("%s", action.errorMsg(value))
	}
	return nil
}

// extractMsg extracts a custom error message from ActionOptions or returns the default message.
func extractMsg(defaultMsg string, value any, option ...ActionOptionFace) string {
	errMsg := defaultMsg
	for _, op := range option {
		if errInterface, ok := op.(CustomErrFace); ok {
			errMsg = errInterface.Msg(value)
		}
	}
	return errMsg
}

// CustomString creates a custom string validator using the provided validation function.
// The optional ActionOptions parameter can be used to customize the error message.
//
// Example:
//
//	CustomString(func(v string) bool { return strings.HasPrefix(v, "test_") })
func CustomString(fn func(value string) bool, option ...ActionOptionFace) StringPipeAction {
	return &stringAction{
		errorMsg: func(v string) string {
			return extractMsg("invalid string", v, option...)
		},
		validate: fn,
	}
}

// NotEmpty validates that a string is not empty.
// The optional ActionOptions parameter can be used to customize the error message.
func NotEmpty(option ...ActionOptionFace) StringPipeAction {
	return &stringAction{
		errorMsg: func(v string) string {
			return extractMsg("cannot be empty", v, option...)
		},
		validate: func(v string) bool {
			return v != ""
		},
	}
}

// Enum validate that a string includes from a set of string.
func Enum(slice []string, option ...ActionOptionFace) StringPipeAction {
	return &stringAction{
		errorMsg: func(v string) string {
			return extractMsg("value is not allowed", v, option...)
		},
		validate: func(v string) bool {
			return slices.Contains(slice, v)
		},
	}
}

// Pattern validates that a string matches the provided regular expression pattern.
// The regex is compiled once during validator creation for efficiency.
// The optional ActionOptions parameter can be used to customize the error message.
//
// Example:
//
//	Pattern(`^\d{3}-\d{2}-\d{4}$`) // SSN format
func Pattern(regexStr string, option ...ActionOptionFace) StringPipeAction {
	regex := regexp.MustCompile(regexStr)
	return &stringAction{
		errorMsg: func(v string) string {
			return extractMsg("string doesn't follow the pattern "+regexStr, v, option...)
		},
		validate: func(v string) bool {
			return regex.MatchString(v)
		},
	}
}

// MaxLength validates that a string does not exceed the specified maximum length.
// The optional ActionOptions parameter can be used to customize the error message.
func MaxLength(max int, option ...ActionOptionFace) StringPipeAction {
	return &stringAction{
		errorMsg: func(v string) string {
			return extractMsg("string length exceeds maximum", v, option...)
		},
		validate: func(v string) bool {
			return is.IsMaxLength(v, max)
		},
	}
}

// MinLength validates that a string has at least the specified minimum length.
// The optional ActionOptions parameter can be used to customize the error message.
func MinLength(min int, option ...ActionOptionFace) StringPipeAction {
	return &stringAction{
		errorMsg: func(v string) string {
			return extractMsg("string length must be at least specified minimum", v, option...)
		},
		validate: func(v string) bool {
			return is.IsMinLength(v, min)
		},
	}
}

// HasPrefix validates that a string starts with the provided prefix.
// The optional ActionOptions parameter can be used to customize the error message.
func HasPrefix(prefix string, option ...ActionOptionFace) StringPipeAction {
	return &stringAction{
		errorMsg: func(v string) string {
			return extractMsg("must start with "+prefix, v, option...)
		},
		validate: func(v string) bool {
			return strings.HasPrefix(v, prefix)
		},
	}
}

// HasSuffix validates that a string end with the provided prefix.
// The optional ActionOptions parameter can be used to customize the error message.
func HasSuffix(suffix string, option ...ActionOptionFace) StringPipeAction {
	return &stringAction{
		errorMsg: func(v string) string {
			return extractMsg("must end with "+suffix, v, option...)
		},
		validate: func(v string) bool {
			return strings.HasSuffix(v, suffix)
		},
	}
}

// EqualFold validates that a string is equal to the provided string under Unicode case-folding.
// This comparison is case-insensitive and handles Unicode correctly.
// The optional ActionOptions parameter can be used to customize the error message.
func EqualFold(target string, option ...ActionOptionFace) StringPipeAction {
	return &stringAction{
		errorMsg: func(v string) string {
			return extractMsg("must be equal to "+target+" (case-insensitive)", v, option...)
		},
		validate: func(v string) bool {
			return strings.EqualFold(v, target)
		},
	}
}

// Contains validates that a string contains the provided substring.
// The optional ActionOptions parameter can be used to customize the error message.
func Contains(substr string, option ...ActionOptionFace) StringPipeAction {
	return &stringAction{
		errorMsg: func(v string) string {
			return extractMsg("must contain "+substr, v, option...)
		},
		validate: func(v string) bool {
			return strings.Contains(v, substr)
		},
	}
}

// IsAlpha validates that a string contains only alphabetic characters.
// The optional ActionOptions parameter can be used to customize the error message.
func IsAlpha(option ...ActionOptionFace) StringPipeAction {
	return &stringAction{
		errorMsg: func(v string) string {
			return extractMsg("must contain only alphabetic characters", v, option...)
		},
		validate: is.IsAlpha,
	}
}

// IsAlphaNumeric validates that a string contains only alphanumeric characters.
// The optional ActionOptions parameter can be used to customize the error message.
func IsAlphaNumeric(option ...ActionOptionFace) StringPipeAction {
	return &stringAction{
		errorMsg: func(v string) string {
			return extractMsg("must contain only alphanumeric characters", v, option...)
		},
		validate: is.IsAlphaNumeric,
	}
}

// IsAscii validates that a string contains only ASCII characters.
// The optional ActionOptions parameter can be used to customize the error message.
func IsAscii(option ...ActionOptionFace) StringPipeAction {
	return &stringAction{
		errorMsg: func(v string) string {
			return extractMsg("must contain only ASCII characters", v, option...)
		},
		validate: is.IsAscii,
	}
}

// IsBase32 validates that a string is valid base32-encoded data.
// The optional ActionOptions parameter can be used to customize the error message.
func IsBase32(option ...ActionOptionFace) StringPipeAction {
	return &stringAction{
		errorMsg: func(v string) string {
			return extractMsg("not a valid base32 string", v, option...)
		},
		validate: is.IsBase32,
	}
}

// IsBase58 validates that a string uses base58 encoding.
// The optional ActionOptions parameter can be used to customize the error message.
func IsBase58(option ...ActionOptionFace) StringPipeAction {
	return &stringAction{
		errorMsg: func(v string) string {
			return extractMsg("not a valid base58 string", v, option...)
		},
		validate: is.IsBase58,
	}
}

// IsBase64 validates that a string is valid base64-encoded data.
// The optional ActionOptions parameter can be used to customize the error message.
func IsBase64(option ...ActionOptionFace) StringPipeAction {
	return &stringAction{
		errorMsg: func(v string) string {
			return extractMsg("not a valid base64 string", v, option...)
		},
		validate: is.IsBase64,
	}
}

// IsBitcoinAddress validates that a string is a valid Bitcoin address.
// The optional ActionOptions parameter can be used to customize the error message.
func IsBitcoinAddress(option ...ActionOptionFace) StringPipeAction {
	return &stringAction{
		errorMsg: func(v string) string {
			return extractMsg("not a valid Bitcoin address", v, option...)
		},
		validate: is.IsBitcoinAddress,
	}
}

// IsCreditCard validates that a string is a valid credit card number using the Luhn algorithm.
// The optional ActionOptions parameter can be used to customize the error message.
func IsCreditCard(option ...ActionOptionFace) StringPipeAction {
	return &stringAction{
		errorMsg: func(v string) string {
			return extractMsg("not a valid credit card number", v, option...)
		},
		validate: is.IsCreditCard,
	}
}

// IsDate validates that a string represents a date in YYYY-MM-DD format.
// The optional ActionOptions parameter can be used to customize the error message.
func IsDate(option ...ActionOptionFace) StringPipeAction {
	return &stringAction{
		errorMsg: func(v string) string {
			return extractMsg("not a valid date", v, option...)
		},
		validate: is.IsDate,
	}
}

// IsDataURI validates that a string is a valid data URI.
// The optional ActionOptions parameter can be used to customize the error message.
func IsDataURI(option ...ActionOptionFace) StringPipeAction {
	return &stringAction{
		errorMsg: func(v string) string {
			return extractMsg("not a valid data URI", v, option...)
		},
		validate: is.IsDataURI,
	}
}

// IsDecimal validates that a string represents a valid decimal number.
// The optional ActionOptions parameter can be used to customize the error message.
func IsDecimal(option ...ActionOptionFace) StringPipeAction {
	return &stringAction{
		errorMsg: func(v string) string {
			return extractMsg("not a valid decimal number", v, option...)
		},
		validate: is.IsDecimal,
	}
}

// IsEmail validates that a string is a valid email address.
// Uses Go's standard mail.ParseAddress for validation.
// The optional ActionOptions parameter can be used to customize the error message.
func IsEmail(option ...ActionOptionFace) StringPipeAction {
	return &stringAction{
		errorMsg: func(v string) string {
			return extractMsg("not a valid email", v, option...)
		},
		validate: func(v string) bool {
			_, err := mail.ParseAddress(v)
			return err == nil
		},
	}
}

// IsEvmAddress validates that a string is a valid Ethereum Virtual Machine address.
// The optional ActionOptions parameter can be used to customize the error message.
func IsEvmAddress(option ...ActionOptionFace) StringPipeAction {
	return &stringAction{
		errorMsg: func(v string) string {
			return extractMsg("not a valid EVM address", v, option...)
		},
		validate: is.IsEvmAddress,
	}
}

// IsHTML validates that a string contains HTML tags.
// The optional ActionOptions parameter can be used to customize the error message.
func IsHTML(option ...ActionOptionFace) StringPipeAction {
	return &stringAction{
		errorMsg: func(v string) string {
			return extractMsg("not a valid HTML string", v, option...)
		},
		validate: is.IsHTML,
	}
}

// IsHexColor validates that a string is a valid hexadecimal color code.
// The optional ActionOptions parameter can be used to customize the error message.
func IsHexColor(option ...ActionOptionFace) StringPipeAction {
	return &stringAction{
		errorMsg: func(v string) string {
			return extractMsg("not a valid hex color", v, option...)
		},
		validate: is.IsHexColor,
	}
}

// IsHexDecimal validates that a string contains only hexadecimal characters.
// The optional ActionOptions parameter can be used to customize the error message.
func IsHexDecimal(option ...ActionOptionFace) StringPipeAction {
	return &stringAction{
		errorMsg: func(v string) string {
			return extractMsg("not a valid hexadecimal string", v, option...)
		},
		validate: is.IsHexDecimal,
	}
}

// IsHSL validates that a string represents a valid HSL color.
// The optional ActionOptions parameter can be used to customize the error message.
func IsHSL(option ...ActionOptionFace) StringPipeAction {
	return &stringAction{
		errorMsg: func(v string) string {
			return extractMsg("not a valid HSL color", v, option...)
		},
		validate: is.IsHSL,
	}
}

// IsIPV4 validates that a string is a valid IPv4 address.
// The optional ActionOptions parameter can be used to customize the error message.
func IsIPV4(option ...ActionOptionFace) StringPipeAction {
	return &stringAction{
		errorMsg: func(v string) string {
			return extractMsg("not a valid IPv4 address", v, option...)
		},
		validate: is.IsIPV4,
	}
}

// IsIPV6 validates that a string is a valid IPv6 address.
// The optional ActionOptions parameter can be used to customize the error message.
func IsIPV6(option ...ActionOptionFace) StringPipeAction {
	return &stringAction{
		errorMsg: func(v string) string {
			return extractMsg("not a valid IPv6 address", v, option...)
		},
		validate: is.IsIPV6,
	}
}

// IsJSON validates that a string is valid JSON data.
// The optional ActionOptions parameter can be used to customize the error message.
func IsJSON(option ...ActionOptionFace) StringPipeAction {
	return &stringAction{
		errorMsg: func(v string) string {
			return extractMsg("not a valid JSON string", v, option...)
		},
		validate: is.IsJSON,
	}
}

// IsRGB validates that a string represents a valid RGB color.
// The optional ActionOptions parameter can be used to customize the error message.
func IsRGB(option ...ActionOptionFace) StringPipeAction {
	return &stringAction{
		errorMsg: func(v string) string {
			return extractMsg("not a valid RGB color", v, option...)
		},
		validate: is.IsRGB,
	}
}

// IsULID validates that a string is a valid ULID (Universally Unique Lexicographically Sortable Identifier).
// The optional ActionOptions parameter can be used to customize the error message.
func IsULID(option ...ActionOptionFace) StringPipeAction {
	return &stringAction{
		errorMsg: func(v string) string {
			return extractMsg("not a valid ULID", v, option...)
		},
		validate: is.IsULID,
	}
}

// IsURL validates that a string is a valid URL.
// The optional ActionOptions parameter can be used to customize the error message.
func IsURL(option ...ActionOptionFace) StringPipeAction {
	return &stringAction{
		errorMsg: func(v string) string {
			return extractMsg("not a valid URL", v, option...)
		},
		validate: is.IsURL,
	}
}

// IsUUID validates that a string is a valid UUID in any version.
// The optional ActionOptions parameter can be used to customize the error message.
func IsUUID(option ...ActionOptionFace) StringPipeAction {
	return &stringAction{
		errorMsg: func(v string) string {
			return extractMsg("not a valid UUID", v, option...)
		},
		validate: is.IsUUID,
	}
}

// IsUUIDV1 validates that a string is a valid UUID version 1.
// The optional ActionOptions parameter can be used to customize the error message.
func IsUUIDV1(option ...ActionOptionFace) StringPipeAction {
	return &stringAction{
		errorMsg: func(v string) string {
			return extractMsg("not a valid UUIDv1", v, option...)
		},
		validate: is.IsUUIDV1,
	}
}

// IsUUIDV3 validates that a string is a valid UUID version 3.
// The optional ActionOptions parameter can be used to customize the error message.
func IsUUIDV3(option ...ActionOptionFace) StringPipeAction {
	return &stringAction{
		errorMsg: func(v string) string {
			return extractMsg("not a valid UUIDv3", v, option...)
		},
		validate: is.IsUUIDV3,
	}
}

// IsUUIDV4 validates that a string is a valid UUID version 4.
// The optional ActionOptions parameter can be used to customize the error message.
func IsUUIDV4(option ...ActionOptionFace) StringPipeAction {
	return &stringAction{
		errorMsg: func(v string) string {
			return extractMsg("not a valid UUIDv4", v, option...)
		},
		validate: is.IsUUIDV4,
	}
}

// IsUUIDV5 validates that a string is a valid UUID version 5.
// The optional ActionOptions parameter can be used to customize the error message.
func IsUUIDV5(option ...ActionOptionFace) StringPipeAction {
	return &stringAction{
		errorMsg: func(v string) string {
			return extractMsg("not a valid UUIDv5", v, option...)
		},
		validate: is.IsUUIDV5,
	}
}

// IsValidPath validates that a string is a valid file system path.
// The optional ActionOptions parameter can be used to customize the error message.
func IsValidPath(option ...ActionOptionFace) StringPipeAction {
	return &stringAction{
		errorMsg: func(v string) string {
			return extractMsg("not a valid path", v, option...)
		},
		validate: is.IsValidPath,
	}
}

// IsValidPort validates that a string represents a valid port number (0-65535).
// The optional ActionOptions parameter can be used to customize the error message.
func IsValidPort(option ...ActionOptionFace) StringPipeAction {
	return &stringAction{
		errorMsg: func(v string) string {
			return extractMsg("not a valid port number", v, option...)
		},
		validate: is.IsValidPort,
	}
}

// IsXML validates that a string is valid XML data.
// The optional ActionOptions parameter can be used to customize the error message.
func IsXML(option ...ActionOptionFace) StringPipeAction {
	return &stringAction{
		errorMsg: func(v string) string {
			return extractMsg("not a valid XML string", v, option...)
		},
		validate: is.IsXML,
	}
}
