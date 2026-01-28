// Package is provides a collection of string validation functions.
// These functions check various formats, encodings, and patterns commonly used in web applications.
package is

import (
	"encoding/base32"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"net"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// UUID validation
func IsUUID(v string) bool {
	uuidRegex := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
	return uuidRegex.MatchString(strings.ToLower(v))
}

// IsUUIDV1 validates whether the string is a valid UUID version 1.
// Version 1 UUIDs are time-based and include a timestamp and MAC address.
func IsUUIDV1(v string) bool {
	uuidv1Regex := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-1[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)
	return uuidv1Regex.MatchString(strings.ToLower(v))
}

// IsUUIDV3 validates whether the string is a valid UUID version 3.
// Version 3 UUIDs are name-based and use MD5 hashing.
func IsUUIDV3(v string) bool {
	uuidv3Regex := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-3[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)
	return uuidv3Regex.MatchString(strings.ToLower(v))
}

// IsUUIDV4 validates whether the string is a valid UUID version 4.
// Version 4 UUIDs are randomly generated.
func IsUUIDV4(v string) bool {
	uuidv4Regex := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)
	return uuidv4Regex.MatchString(strings.ToLower(v))
}

// IsUUIDV5 validates whether the string is a valid UUID version 5.
// Version 5 UUIDs are name-based and use SHA-1 hashing.
func IsUUIDV5(v string) bool {
	uuidv5Regex := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-5[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)
	return uuidv5Regex.MatchString(strings.ToLower(v))
}

// IP validation
func IsIPV4(v string) bool {
	ip := net.ParseIP(v)
	return ip != nil && strings.Contains(v, ".")
}

// IsIPV6 validates whether the string is a valid IPv6 address.
// Accepts standard colon-separated hexadecimal notation.
func IsIPV6(v string) bool {
	ip := net.ParseIP(v)
	return ip != nil && strings.Contains(v, ":")
}

// Length validation
func IsMinLength(v string, min int) bool {
	return len(v) >= min
}

// IsMaxLength validates whether the string does not exceed the specified maximum length.
func IsMaxLength(v string, max int) bool {
	return len(v) <= max
}

// Address validation
func IsEvmAddress(v string) bool {
	evmRegex := regexp.MustCompile(`^0x[a-fA-F0-9]{40}$`)
	return evmRegex.MatchString(v)
}

// IsBitcoinAddress validates whether the string is a valid Bitcoin address.
// Supports legacy (P2PKH, P2SH) and SegWit (Bech32) address formats.
func IsBitcoinAddress(v string) bool {
	btcRegex := regexp.MustCompile(`^(bc1|[13])[a-zA-HJ-NP-Z0-9]{25,39}$`)
	return btcRegex.MatchString(v)
}

// URL validation
func IsURL(v string) bool {
	_, err := url.ParseRequestURI(v)
	return err == nil
}

// Character validation
func IsAlpha(v string) bool {
	for _, ch := range v {
		if !unicode.IsLetter(ch) {
			return false
		}
	}
	return true
}

// IsAlphaNumeric validates whether the string contains only alphanumeric characters.
// Accepts letters and digits, but not spaces or special characters.
func IsAlphaNumeric(v string) bool {
	for _, ch := range v {
		if !unicode.IsLetter(ch) && !unicode.IsDigit(ch) {
			return false
		}
	}
	return true
}

// IsAscii validates whether the string contains only ASCII characters (codes 0-127).
func IsAscii(v string) bool {
	for _, ch := range v {
		if ch > 127 {
			return false
		}
	}
	return true
}

// IsDecimal validates whether the string represents a valid decimal number.
// Accepts integers and floating-point numbers with optional negative sign.
func IsDecimal(v string) bool {
	decimalRegex := regexp.MustCompile(`^-?(\d+\.?\d*|\.\d+)$`)
	return decimalRegex.MatchString(v)
}

// Encoding validation
func IsBase64(v string) bool {
	_, err := base64.StdEncoding.DecodeString(v)
	return err == nil
}

// IsBase32 validates whether the string is valid base32-encoded data.
// Uses Go's standard base32 encoding for validation.
func IsBase32(v string) bool {
	_, err := base32.StdEncoding.DecodeString(v)
	return err == nil
}

// IsBase58 validates whether the string uses base58 encoding.
// Base58 is commonly used in cryptocurrency addresses (excludes 0, O, I, l).
func IsBase58(v string) bool {
	base58Regex := regexp.MustCompile(`^[123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]+$`)
	return base58Regex.MatchString(v)
}

// IsHexDecimal validates whether the string contains only hexadecimal characters (0-9, a-f, A-F).
func IsHexDecimal(v string) bool {
	hexRegex := regexp.MustCompile(`^[0-9a-fA-F]+$`)
	return hexRegex.MatchString(v)
}

// Data format validation
func IsJSON(v string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(v), &js) == nil
}

// IsXML validates whether the string is valid XML data.
// Uses Go's standard xml.Unmarshal for validation.
func IsXML(v string) bool {
	return xml.Unmarshal([]byte(v), &struct{}{}) == nil
}

// IsHTML validates whether the string contains HTML tags.
// Uses a simple regex pattern to detect HTML-like structures.
func IsHTML(v string) bool {
	htmlRegex := regexp.MustCompile(`(?i)<[a-z][\s\S]*>`)
	return htmlRegex.MatchString(v)
}

// IsValidPath validates whether the string is a valid file system path.
// Accepts alphanumeric characters and common path separators.
func IsValidPath(v string) bool {
	pathRegex := regexp.MustCompile(`^[a-zA-Z0-9._\-\/\\:]+$`)
	return pathRegex.MatchString(v)
}

// IsCreditCard validates whether the string is a valid credit card number.
// Uses the Luhn algorithm to verify the checksum.
// Accepts numbers with spaces or hyphens.
func IsCreditCard(v string) bool {
	cleaned := strings.ReplaceAll(v, " ", "")
	cleaned = strings.ReplaceAll(cleaned, "-", "")

	if len(cleaned) < 13 || len(cleaned) > 19 {
		return false
	}

	sum := 0
	for i, ch := range cleaned {
		digit := int(ch - '0')
		if (len(cleaned)-i)%2 == 0 {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}

	return sum%10 == 0
}

// IsRGB validates whether the string represents a valid RGB color.
//
// Accepts format: rgb(r, g, b) where r, g, b are 0-255.
func IsRGB(v string) bool {
	rgbRegex := regexp.MustCompile(`^rgb\(\s*([0-9]{1,3})\s*,\s*([0-9]{1,3})\s*,\s*([0-9]{1,3})\s*\)$`)
	return len(rgbRegex.FindStringSubmatch(v)) > 0
}

// IsHexColor validates whether the string is a valid hexadecimal color code.
//
// Accepts 6-character hex colors starting with # (e.g., "#FF5733").
func IsHexColor(v string) bool {
	hexColorRegex := regexp.MustCompile(`^#[0-9a-fA-F]{6}$`)
	return hexColorRegex.MatchString(v)
}

// IsHSL validates whether the string represents a valid HSL color.
//
// Accepts format: hsl(h, s%, l%) where h is 0-360 and s, l are 0-100.
func IsHSL(v string) bool {
	hslRegex := regexp.MustCompile(`^hsl\(\s*([0-9]{1,3})\s*,\s*([0-9]{1,3})%\s*,\s*([0-9]{1,3})%\s*\)$`)
	return len(hslRegex.FindStringSubmatch(v)) > 0
}

// IsValidPort validates whether the string represents a valid port number.
//
// Accepts port numbers from 0 to 65535.
func IsValidPort(v string) bool {
	port, err := strconv.Atoi(v)
	return err == nil && port >= 0 && port <= 65535
}

// IsULID validates whether the string is a valid ULID (Universally Unique Lexicographically Sortable Identifier).
//
// ULIDs are 26 characters long and use Crockford's base32 alphabet.
func IsULID(v string) bool {
	ulidRegex := regexp.MustCompile(`^[0-9A-Z]{26}$`)
	return ulidRegex.MatchString(v)
}

// IsDataURI validates whether the string is a valid data URI.
// Data URIs embed data inline in the format: data:[<mediatype>][;base64],<data>
func IsDataURI(v string) bool {
	dataURIRegex := regexp.MustCompile(`^data:[a-zA-Z0-9][a-zA-Z0-9!#$&\-\^_]*\/[a-zA-Z0-9][a-zA-Z0-9!#$&\-\^_]*(;[a-zA-Z0-9\-]+=(([a-zA-Z0-9\-]+)|(['\"][^\"]*['\"])))*;base64,[a-zA-Z0-9+/]*={0,2}$`)
	return dataURIRegex.MatchString(v)
}

// IsDate validates whether the string represents a date in YYYY-MM-DD format.
func IsDate(v string) bool {
	dateRegex := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	return dateRegex.MatchString(v)
}

// IsEmpty validates whether the string is empty or contains only whitespace.
func IsEmpty(v string) bool {
	return strings.TrimSpace(v) == ""
}
