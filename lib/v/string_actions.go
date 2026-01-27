package v

import (
	"fmt"
	"net/mail"

	"github.com/mrbns/valgo/lib/is"
)

type stringAction struct {
	errorMsg string
	validate func(v string) bool
}

func (acttion *stringAction) Run(value string) error {
	if !acttion.validate(value) {
		return fmt.Errorf("%s", acttion.errorMsg)
	}
	return nil
}

func extractMsg(defaultMsg string, option ...ActionOptions) string {
	errMsg := defaultMsg
	for _, op := range option {
		if val, ok := op.(ErrMsg); ok {
			errMsg = val.Msg()
		}
	}
	return errMsg
}

func Empty(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("cannot be empty", option...),
		validate: func(v string) bool {
			return v != ""
		},
	}
}

func IsEmail(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid email", option...),
		validate: func(v string) bool {
			_, err := mail.ParseAddress(v)
			return err == nil
		},
	}
}

// UUID validators
func IsUUID(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid UUID", option...),
		validate: is.IsUUID,
	}
}

func IsUUIDV1(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid UUIDv1", option...),
		validate: is.IsUUIDV1,
	}
}

func IsUUIDV3(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid UUIDv3", option...),
		validate: is.IsUUIDV3,
	}
}

func IsUUIDV4(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid UUIDv4", option...),
		validate: is.IsUUIDV4,
	}
}

func IsUUIDV5(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid UUIDv5", option...),
		validate: is.IsUUIDV5,
	}
}

// IP validators
func IsIPV4(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid IPv4 address", option...),
		validate: is.IsIPV4,
	}
}

func IsIPV6(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid IPv6 address", option...),
		validate: is.IsIPV6,
	}
}

// Length validators
func IsMinLength(min int, option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("string length must be at least specified minimum", option...),
		validate: func(v string) bool {
			return is.IsMinLength(v, min)
		},
	}
}

func IsMaxLength(max int, option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("string length exceeds maximum", option...),
		validate: func(v string) bool {
			return is.IsMaxLength(v, max)
		},
	}
}

// Address validators
func IsEvmAddress(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid EVM address", option...),
		validate: is.IsEvmAddress,
	}
}

func IsBitcoinAddress(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid Bitcoin address", option...),
		validate: is.IsBitcoinAddress,
	}
}

// URL validator
func IsURL(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid URL", option...),
		validate: is.IsURL,
	}
}

// Character validators
func IsAlpha(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("must contain only alphabetic characters", option...),
		validate: is.IsAlpha,
	}
}

func IsAlphaNumeric(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("must contain only alphanumeric characters", option...),
		validate: is.IsAlphaNumeric,
	}
}

func IsAscii(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("must contain only ASCII characters", option...),
		validate: is.IsAscii,
	}
}

// Number validator
func IsDecimal(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid decimal number", option...),
		validate: is.IsDecimal,
	}
}

// Encoding validators
func IsBase64(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid base64 string", option...),
		validate: is.IsBase64,
	}
}

func IsBase32(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid base32 string", option...),
		validate: is.IsBase32,
	}
}

func IsBase58(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid base58 string", option...),
		validate: is.IsBase58,
	}
}

func IsHexDecimal(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid hexadecimal string", option...),
		validate: is.IsHexDecimal,
	}
}

// Data format validators
func IsJSON(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid JSON string", option...),
		validate: is.IsJSON,
	}
}

func IsXML(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid XML string", option...),
		validate: is.IsXML,
	}
}

func IsHTML(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid HTML string", option...),
		validate: is.IsHTML,
	}
}

// Path validator
func IsValidPath(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid path", option...),
		validate: is.IsValidPath,
	}
}

// Credit card validator
func IsCreditCard(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid credit card number", option...),
		validate: is.IsCreditCard,
	}
}

// Color validators
func IsRGB(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid RGB color", option...),
		validate: is.IsRGB,
	}
}

func IsHexColor(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid hex color", option...),
		validate: is.IsHexColor,
	}
}

func IsHSL(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid HSL color", option...),
		validate: is.IsHSL,
	}
}

// Port validator
func IsValidPort(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid port number", option...),
		validate: is.IsValidPort,
	}
}

// ULID validator
func IsULID(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid ULID", option...),
		validate: is.IsULID,
	}
}

// Data URI validator
func IsDataURI(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid data URI", option...),
		validate: is.IsDataURI,
	}
}

// Date validator
func IsDate(option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("not a valid date", option...),
		validate: is.IsDate,
	}
}

// string custom

func CustomString(fn func(value string) bool, option ...ActionOptions) StringPipeAction {
	return &stringAction{
		errorMsg: extractMsg("invalid string", option...),
		validate: fn,
	}
}
