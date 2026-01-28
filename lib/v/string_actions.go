package v

import (
	"fmt"
	"net/mail"

	"github.com/mrbns/valgo/lib/is"
)

// stringAction implements StringPipeAction for string validation.
type stringAction struct {
	errorMsg string
	validate func(v string) bool
}

// Run executes the validation function on the given string value.
// Returns an error if validation fails.
func (acttion *stringAction) Run(value string) error {
	if !acttion.validate(value) {
		return fmt.Errorf("%s", acttion.errorMsg)
	}
	return nil
}

// extractMsg extracts a custom error message from ActionOptions or returns the default message.
func extractMsg(defaultMsg string, option ...ActionOptions) string {
	errMsg := defaultMsg
	for _, op := range option {
		if val, ok := op.(ErrMsg); ok {
			errMsg = val.Msg()
		}
	}
	return errMsg
}

// Empty validates that a string is not empty.
// The optional ActionOptions parameter can be used to customize the error message.
func Empty(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("cannot be empty", option...),
		validate: func(v string) bool {
			return v != ""
		},
	}
}

// IsEmail validates that a string is a valid email address.
// Uses Go's standard mail.ParseAddress for validation.
// The optional ActionOptions parameter can be used to customize the error message.
func IsEmail(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid email", option...),
		validate: func(v string) bool {
			_, err := mail.ParseAddress(v)
			return err == nil
		},
	}
}

// IsUUID validates that a string is a valid UUID in any version.
// The optional ActionOptions parameter can be used to customize the error message.
func IsUUID(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid UUID", option...),
		validate: is.IsUUID,
	}
}

// IsUUIDV1 validates that a string is a valid UUID version 1.
// The optional ActionOptions parameter can be used to customize the error message.
func IsUUIDV1(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid UUIDv1", option...),
		validate: is.IsUUIDV1,
	}
}

// IsUUIDV3 validates that a string is a valid UUID version 3.
// The optional ActionOptions parameter can be used to customize the error message.
func IsUUIDV3(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid UUIDv3", option...),
		validate: is.IsUUIDV3,
	}
}

// IsUUIDV4 validates that a string is a valid UUID version 4.
// The optional ActionOptions parameter can be used to customize the error message.
func IsUUIDV4(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid UUIDv4", option...),
		validate: is.IsUUIDV4,
	}
}

// IsUUIDV5 validates that a string is a valid UUID version 5.
// The optional ActionOptions parameter can be used to customize the error message.
func IsUUIDV5(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid UUIDv5", option...),
		validate: is.IsUUIDV5,
	}
}

// IsIPV4 validates that a string is a valid IPv4 address.
// The optional ActionOptions parameter can be used to customize the error message.
func IsIPV4(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid IPv4 address", option...),
		validate: is.IsIPV4,
	}
}

// IsIPV6 validates that a string is a valid IPv6 address.
// The optional ActionOptions parameter can be used to customize the error message.
func IsIPV6(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid IPv6 address", option...),
		validate: is.IsIPV6,
	}
}

// IsMinLength validates that a string has at least the specified minimum length.
// The optional ActionOptions parameter can be used to customize the error message.
func IsMinLength(min int, option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("string length must be at least specified minimum", option...),
		validate: func(v string) bool {
			return is.IsMinLength(v, min)
		},
	}
}

// IsMaxLength validates that a string does not exceed the specified maximum length.
// The optional ActionOptions parameter can be used to customize the error message.
func IsMaxLength(max int, option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("string length exceeds maximum", option...),
		validate: func(v string) bool {
			return is.IsMaxLength(v, max)
		},
	}
}

// IsEvmAddress validates that a string is a valid Ethereum Virtual Machine address.
// The optional ActionOptions parameter can be used to customize the error message.
func IsEvmAddress(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid EVM address", option...),
		validate: is.IsEvmAddress,
	}
}

// IsBitcoinAddress validates that a string is a valid Bitcoin address.
// The optional ActionOptions parameter can be used to customize the error message.
func IsBitcoinAddress(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid Bitcoin address", option...),
		validate: is.IsBitcoinAddress,
	}
}

// IsURL validates that a string is a valid URL.
// The optional ActionOptions parameter can be used to customize the error message.
func IsURL(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid URL", option...),
		validate: is.IsURL,
	}
}

// IsAlpha validates that a string contains only alphabetic characters.
// The optional ActionOptions parameter can be used to customize the error message.
func IsAlpha(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("must contain only alphabetic characters", option...),
		validate: is.IsAlpha,
	}
}

// IsAlphaNumeric validates that a string contains only alphanumeric characters.
// The optional ActionOptions parameter can be used to customize the error message.
func IsAlphaNumeric(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("must contain only alphanumeric characters", option...),
		validate: is.IsAlphaNumeric,
	}
}

// IsAscii validates that a string contains only ASCII characters.
// The optional ActionOptions parameter can be used to customize the error message.
func IsAscii(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("must contain only ASCII characters", option...),
		validate: is.IsAscii,
	}
}

// IsDecimal validates that a string represents a valid decimal number.
// The optional ActionOptions parameter can be used to customize the error message.
func IsDecimal(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid decimal number", option...),
		validate: is.IsDecimal,
	}
}

// IsBase64 validates that a string is valid base64-encoded data.
// The optional ActionOptions parameter can be used to customize the error message.
func IsBase64(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid base64 string", option...),
		validate: is.IsBase64,
	}
}

// IsBase32 validates that a string is valid base32-encoded data.
// The optional ActionOptions parameter can be used to customize the error message.
func IsBase32(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid base32 string", option...),
		validate: is.IsBase32,
	}
}

// IsBase58 validates that a string uses base58 encoding.
// The optional ActionOptions parameter can be used to customize the error message.
func IsBase58(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid base58 string", option...),
		validate: is.IsBase58,
	}
}

// IsHexDecimal validates that a string contains only hexadecimal characters.
// The optional ActionOptions parameter can be used to customize the error message.
func IsHexDecimal(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid hexadecimal string", option...),
		validate: is.IsHexDecimal,
	}
}

// IsJSON validates that a string is valid JSON data.
// The optional ActionOptions parameter can be used to customize the error message.
func IsJSON(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid JSON string", option...),
		validate: is.IsJSON,
	}
}

// IsXML validates that a string is valid XML data.
// The optional ActionOptions parameter can be used to customize the error message.
func IsXML(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid XML string", option...),
		validate: is.IsXML,
	}
}

// IsHTML validates that a string contains HTML tags.
// The optional ActionOptions parameter can be used to customize the error message.
func IsHTML(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid HTML string", option...),
		validate: is.IsHTML,
	}
}

// IsValidPath validates that a string is a valid file system path.
// The optional ActionOptions parameter can be used to customize the error message.
func IsValidPath(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid path", option...),
		validate: is.IsValidPath,
	}
}

// IsCreditCard validates that a string is a valid credit card number using the Luhn algorithm.
// The optional ActionOptions parameter can be used to customize the error message.
func IsCreditCard(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid credit card number", option...),
		validate: is.IsCreditCard,
	}
}

// IsRGB validates that a string represents a valid RGB color.
// The optional ActionOptions parameter can be used to customize the error message.
func IsRGB(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid RGB color", option...),
		validate: is.IsRGB,
	}
}

// IsHexColor validates that a string is a valid hexadecimal color code.
// The optional ActionOptions parameter can be used to customize the error message.
func IsHexColor(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid hex color", option...),
		validate: is.IsHexColor,
	}
}

// IsHSL validates that a string represents a valid HSL color.
// The optional ActionOptions parameter can be used to customize the error message.
func IsHSL(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid HSL color", option...),
		validate: is.IsHSL,
	}
}

// IsValidPort validates that a string represents a valid port number (0-65535).
// The optional ActionOptions parameter can be used to customize the error message.
func IsValidPort(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid port number", option...),
		validate: is.IsValidPort,
	}
}

// IsULID validates that a string is a valid ULID (Universally Unique Lexicographically Sortable Identifier).
// The optional ActionOptions parameter can be used to customize the error message.
func IsULID(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid ULID", option...),
		validate: is.IsULID,
	}
}

// IsDataURI validates that a string is a valid data URI.
// The optional ActionOptions parameter can be used to customize the error message.
func IsDataURI(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid data URI", option...),
		validate: is.IsDataURI,
	}
}

// IsDate validates that a string represents a date in YYYY-MM-DD format.
// The optional ActionOptions parameter can be used to customize the error message.
func IsDate(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid date", option...),
		validate: is.IsDate,
	}
}

// CustomString creates a custom string validator using the provided validation function.
// The optional ActionOptions parameter can be used to customize the error message.
//
// Example:
//
//	CustomString(func(v string) bool { return strings.HasPrefix(v, "test_") })
func CustomString(fn func(value string) bool, option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("invalid string", option...),
		validate: fn,
	}
}
