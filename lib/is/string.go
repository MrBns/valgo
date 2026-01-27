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

func IsUUIDV1(v string) bool {
	uuidv1Regex := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-1[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)
	return uuidv1Regex.MatchString(strings.ToLower(v))
}

func IsUUIDV3(v string) bool {
	uuidv3Regex := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-3[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)
	return uuidv3Regex.MatchString(strings.ToLower(v))
}

func IsUUIDV4(v string) bool {
	uuidv4Regex := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)
	return uuidv4Regex.MatchString(strings.ToLower(v))
}

func IsUUIDV5(v string) bool {
	uuidv5Regex := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-5[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)
	return uuidv5Regex.MatchString(strings.ToLower(v))
}

// IP validation
func IsIPV4(v string) bool {
	ip := net.ParseIP(v)
	return ip != nil && strings.Contains(v, ".")
}

func IsIPV6(v string) bool {
	ip := net.ParseIP(v)
	return ip != nil && strings.Contains(v, ":")
}

// Length validation
func IsMinLength(v string, min int) bool {
	return len(v) >= min
}

func IsMaxLength(v string, max int) bool {
	return len(v) <= max
}

// Address validation
func IsEvmAddress(v string) bool {
	evmRegex := regexp.MustCompile(`^0x[a-fA-F0-9]{40}$`)
	return evmRegex.MatchString(v)
}

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

func IsAlphaNumeric(v string) bool {
	for _, ch := range v {
		if !unicode.IsLetter(ch) && !unicode.IsDigit(ch) {
			return false
		}
	}
	return true
}

func IsAscii(v string) bool {
	for _, ch := range v {
		if ch > 127 {
			return false
		}
	}
	return true
}

func IsDecimal(v string) bool {
	decimalRegex := regexp.MustCompile(`^-?(\d+\.?\d*|\.\d+)$`)
	return decimalRegex.MatchString(v)
}

// Encoding validation
func IsBase64(v string) bool {
	_, err := base64.StdEncoding.DecodeString(v)
	return err == nil
}

func IsBase32(v string) bool {
	_, err := base32.StdEncoding.DecodeString(v)
	return err == nil
}

func IsBase58(v string) bool {
	base58Regex := regexp.MustCompile(`^[123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]+$`)
	return base58Regex.MatchString(v)
}

func IsHexDecimal(v string) bool {
	hexRegex := regexp.MustCompile(`^[0-9a-fA-F]+$`)
	return hexRegex.MatchString(v)
}

// Data format validation
func IsJSON(v string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(v), &js) == nil
}

func IsXML(v string) bool {
	return xml.Unmarshal([]byte(v), &struct{}{}) == nil
}

func IsHTML(v string) bool {
	htmlRegex := regexp.MustCompile(`(?i)<[a-z][\s\S]*>`)
	return htmlRegex.MatchString(v)
}

// Path validation
func IsValidPath(v string) bool {
	pathRegex := regexp.MustCompile(`^[a-zA-Z0-9._\-\/\\:]+$`)
	return pathRegex.MatchString(v)
}

// Credit card validation
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

// Color validation
func IsRGB(v string) bool {
	rgbRegex := regexp.MustCompile(`^rgb\(\s*([0-9]{1,3})\s*,\s*([0-9]{1,3})\s*,\s*([0-9]{1,3})\s*\)$`)
	return len(rgbRegex.FindStringSubmatch(v)) > 0
}

func IsHexColor(v string) bool {
	hexColorRegex := regexp.MustCompile(`^#[0-9a-fA-F]{6}$`)
	return hexColorRegex.MatchString(v)
}

func IsHSL(v string) bool {
	hslRegex := regexp.MustCompile(`^hsl\(\s*([0-9]{1,3})\s*,\s*([0-9]{1,3})%\s*,\s*([0-9]{1,3})%\s*\)$`)
	return len(hslRegex.FindStringSubmatch(v)) > 0
}

// Port validation
func IsValidPort(v string) bool {
	port, err := strconv.Atoi(v)
	return err == nil && port >= 0 && port <= 65535
}

// ULID validation
func IsULID(v string) bool {
	ulidRegex := regexp.MustCompile(`^[0-9A-Z]{26}$`)
	return ulidRegex.MatchString(v)
}

// Data URI validation
func IsDataURI(v string) bool {
	dataURIRegex := regexp.MustCompile(`^data:[a-zA-Z0-9][a-zA-Z0-9!#$&\-\^_]*\/[a-zA-Z0-9][a-zA-Z0-9!#$&\-\^_]*(;[a-zA-Z0-9\-]+=(([a-zA-Z0-9\-]+)|(['\"][^\"]*['\"])))*;base64,[a-zA-Z0-9+/]*={0,2}$`)
	return dataURIRegex.MatchString(v)
}

// Date validation
func IsDate(v string) bool {
	dateRegex := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	return dateRegex.MatchString(v)
}

// Empty validation
func IsEmpty(v string) bool {
	return strings.TrimSpace(v) == ""
}
