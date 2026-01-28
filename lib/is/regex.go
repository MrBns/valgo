package is

import "regexp"

// Pre-compiled regex patterns for performance
var (
	// UUID: 8-4-4-4-12 hex format, case-insensitive (handled by ToLower in function)
	uuidRegex = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)
	// UUID v1: time-based, version=1, variant bits=[89ab]
	uuidv1Regex = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-1[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)
	// UUID v3: MD5 hash-based, version=3, variant bits=[89ab]
	uuidv3Regex = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-3[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)
	// UUID v4: random, version=4, variant bits=[89ab]
	uuidv4Regex = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)
	// UUID v5: SHA-1 hash-based, version=5, variant bits=[89ab]
	uuidv5Regex = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-5[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)
	// EVM/Ethereum address: 0x prefix + 40 hex chars, case-insensitive
	evmRegex = regexp.MustCompile(`(?i)^0x[0-9a-f]{40}$`)
	// Bitcoin address: P2PKH (1...), P2SH (3...), Bech32 (bc1q...), Bech32m/Taproot (bc1p...)
	btcRegex = regexp.MustCompile(`^(bc1[qpzry9x8gf2tvdw0s3jn54khce6mua7l]{39,59}|[13][a-km-zA-HJ-NP-Z1-9]{25,34})$`)
	// Decimal number: optional sign, integer or float, no trailing dot alone
	decimalRegex = regexp.MustCompile(`^-?(?:0|[1-9]\d*)(?:\.\d+)?$|^-?\.\d+$`)
	// Base58: Bitcoin alphabet (no 0, O, I, l)
	base58Regex = regexp.MustCompile(`^[123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]+$`)
	// Hexadecimal: optional 0x prefix, at least one hex digit
	hexRegex = regexp.MustCompile(`^(?:0[xX])?[0-9a-fA-F]+$`)
	// HTML: matches opening/closing/self-closing tags with attributes
	htmlRegex = regexp.MustCompile(`(?i)<\/?[a-z][a-z0-9]*(?:\s+[a-z][a-z0-9\-]*(?:\s*=\s*(?:"[^"]*"|'[^']*'|[^\s>]+))?)*\s*\/?>`)
	// File path: Unix/Windows paths with common chars, spaces allowed
	pathRegex = regexp.MustCompile(`^(?:[a-zA-Z]:)?[a-zA-Z0-9._\-\/\\:\s~@]+$`)
	// RGB color: rgb(r, g, b) or rgba(r, g, b, a) with 0-255 values or percentages
	rgbRegex = regexp.MustCompile(`(?i)^rgba?\(\s*(\d{1,3}%?)\s*[,\s]\s*(\d{1,3}%?)\s*[,\s]\s*(\d{1,3}%?)(?:\s*[,\/]\s*([\d.]+%?))?\s*\)$`)
	// Hex color: #RGB, #RGBA, #RRGGBB, #RRGGBBAA formats
	hexColorRegex = regexp.MustCompile(`(?i)^#(?:[0-9a-f]{3}|[0-9a-f]{4}|[0-9a-f]{6}|[0-9a-f]{8})$`)
	// HSL color: hsl(h, s%, l%) or hsla(h, s%, l%, a) with degree/turn/rad support
	hslRegex = regexp.MustCompile(`(?i)^hsla?\(\s*(\d{1,3}(?:\.\d+)?(?:deg|rad|turn)?)\s*[,\s]\s*(\d{1,3}(?:\.\d+)?%?)\s*[,\s]\s*(\d{1,3}(?:\.\d+)?%?)(?:\s*[,\/]\s*([\d.]+%?))?\s*\)$`)
	// ULID: Crockford's Base32 alphabet (excludes I, L, O, U), 26 chars
	ulidRegex = regexp.MustCompile(`^[0-9A-HJKMNP-TV-Z]{26}$`)
	// Data URI: data:[<mediatype>][;charset=...][;base64],<data>
	dataURIRegex = regexp.MustCompile(`^data:(?:[a-zA-Z0-9]+\/[a-zA-Z0-9\-+.]+)?(?:;[a-zA-Z0-9\-]+=(?:[a-zA-Z0-9\-]+|"[^"]*"))*(?:;base64)?,[a-zA-Z0-9+/\-_=%]*$`)
	// Date: YYYY-MM-DD with valid month (01-12) and day (01-31) ranges
	dateRegex = regexp.MustCompile(`^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01])$`)
)
