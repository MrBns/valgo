# Valigo - Simple Validation Library for Go

A flexible and easy-to-use validation framework for Go that supports strings, integers, and floats through a clean pipeline pattern.

## ✨ Features

- **Simple API** - Easy to learn and use
- **Pipeline Pattern** - Chain multiple validations together
- **Type Support** - Validate strings, integers, and floats
- **Custom Validators** - Create your own validation rules
- **Custom Error Messages** - Personalize error messages with `{VALUE}` placeholder
- **Concurrent Validation** - Validate multiple fields at once
- **Rich String Validators** - 40+ built-in string validation functions

## 📦 Installation

```bash
go get github.com/mrbns/valgo
```

## 🚀 Quick Start

Here's a simple example to validate an email:

```go
package main

import (
    "fmt"
    "github.com/mrbns/valgo/lib/v"
)

func main() {
    // Create a validation pipe for an email
    pipe := v.NewStringPipe("user@example.com", v.IsEmail())
    
    // Validate
    if err := pipe.Validate(); err != nil {
        fmt.Println("Validation failed:", err.Err)
    } else {
        fmt.Println("Email is valid!")
    }
}
```

## 📖 Basic Concepts

### Pipe
A **pipe** is a sequence of validation checks. Think of it as a series of tests that your data must pass.

### Action
An **action** is a single validation rule (like "must be an email" or "must be at least 5 characters").

### Schema
A **schema** is a collection of pipes for validating multiple fields together.

## 💡 Usage Examples

### Validating a Single String

```go
pipe := v.NewStringPipe(
    "test@example.com",
    v.Empty(),           // Must not be empty
    v.IsEmail(),         // Must be a valid email
    v.MaxLength(50),     // Cannot exceed 50 characters
)

err := pipe.Validate()
if err != nil {
    fmt.Println(err.Err)
}
```

### Validating a Number

```go
pipe := v.NewNumberPipe(
    25,
    v.Min(18),           // Must be at least 18
    v.Max(100),          // Cannot exceed 100
    v.IsPositive(),      // Must be positive
)

err := pipe.Validate()
```

### Validating a Float

```go
pipe := v.NewFloatPipe(
    99.99,
    v.MinFloat(0.0),     // Must be at least 0
    v.MaxFloat(1000.0),  // Cannot exceed 1000
)

err := pipe.Validate()
```

### Validating Multiple Fields (Schema)

```go
schema := v.NewPipesMap(map[string]v.Pipe{
    "email": v.NewStringPipe("user@example.com", v.IsEmail()),
    "age":   v.NewNumberPipe(25, v.Min(18)),
    "price": v.NewFloatPipe(19.99, v.MinFloat(0)),
})

// Validate and stop at first error
if err := schema.Validate(); err != nil {
    fmt.Printf("Field '%s' failed: %v\n", err.Key, err.Err)
}

// Or validate all fields and get all errors
if errs := schema.ValidateAll(); errs != nil {
    for _, err := range errs {
        fmt.Printf("Field '%s' failed: %v\n", err.Key, err.Err)
    }
}
```

### Custom Error Messages

You can customize error messages using the `ErrMsg` function with `{VALUE}` placeholder:

```go
pipe := v.NewNumberPipe(
    15,
    v.Min(18, v.ErrMsg("You must be at least 18, but you are {VALUE}")),
    v.Max(65, v.ErrMsg("{VALUE} exceeds the maximum age of 65")),
)
```

### Custom Validators

Create your own validation logic:

```go
// Custom string validator
customCheck := v.CustomString(func(value string) bool {
    return strings.HasPrefix(value, "user_")
}, v.ErrMsg("Username must start with 'user_'"))

pipe := v.NewStringPipe("user_john", customCheck)

// Custom number validator
evenNumber := v.CustomNumber(func(value int) bool {
    return value%2 == 0
}, v.ErrMsg("{VALUE} must be an even number"))

numberPipe := v.NewNumberPipe(10, evenNumber)
```

## 📋 Available String Validators

| Validator | Description | Example |
|-----------|-------------|---------|
| `Empty()` | Checks if string is not empty | `v.Empty()` |
| `MinLength(n)` | Minimum length required | `v.MinLength(5)` |
| `MaxLength(n)` | Maximum length allowed | `v.MaxLength(100)` |
| `Pattern(regex)` | Match regex pattern | `v.Pattern("^[A-Z]")` |
| `IsEmail()` | Valid email address | `v.IsEmail()` |
| `IsURL()` | Valid URL | `v.IsURL()` |
| `IsUUID()` | Valid UUID (any version) | `v.IsUUID()` |
| `IsUUIDV4()` | Valid UUID v4 | `v.IsUUIDV4()` |
| `IsAlpha()` | Only letters (a-z, A-Z) | `v.IsAlpha()` |
| `IsAlphaNumeric()` | Letters and numbers only | `v.IsAlphaNumeric()` |
| `IsAscii()` | Only ASCII characters | `v.IsAscii()` |
| `IsDecimal()` | Valid decimal number | `v.IsDecimal()` |
| `IsJSON()` | Valid JSON string | `v.IsJSON()` |
| `IsXML()` | Valid XML string | `v.IsXML()` |
| `IsHTML()` | Contains HTML tags | `v.IsHTML()` |
| `IsBase64()` | Valid base64 encoding | `v.IsBase64()` |
| `IsBase32()` | Valid base32 encoding | `v.IsBase32()` |
| `IsBase58()` | Valid base58 encoding | `v.IsBase58()` |
| `IsHexDecimal()` | Hexadecimal string | `v.IsHexDecimal()` |
| `IsIPV4()` | Valid IPv4 address | `v.IsIPV4()` |
| `IsIPV6()` | Valid IPv6 address | `v.IsIPV6()` |
| `IsEvmAddress()` | Ethereum address | `v.IsEvmAddress()` |
| `IsBitcoinAddress()` | Bitcoin address | `v.IsBitcoinAddress()` |
| `IsCreditCard()` | Credit card number (Luhn) | `v.IsCreditCard()` |
| `IsDate()` | Date in YYYY-MM-DD format | `v.IsDate()` |
| `IsDataURI()` | Valid data URI | `v.IsDataURI()` |
| `IsHexColor()` | Hex color code | `v.IsHexColor()` |
| `IsRGB()` | RGB color format | `v.IsRGB()` |
| `IsHSL()` | HSL color format | `v.IsHSL()` |
| `IsValidPort()` | Port number (0-65535) | `v.IsValidPort()` |
| `IsValidPath()` | File system path | `v.IsValidPath()` |
| `IsULID()` | Valid ULID | `v.IsULID()` |

## 🔢 Available Integer Validators

| Validator | Description |
|-----------|-------------|
| `Min(n)` | Minimum value |
| `Max(n)` | Maximum value |
| `Gt(n)` | Greater than |
| `Gte(n)` | Greater than or equal |
| `Lt(n)` | Less than |
| `Lte(n)` | Less than or equal |
| `IsPositive()` | Must be > 0 |
| `IsNegative()` | Must be < 0 |
| `IsZero()` | Must equal 0 |

## 📊 Available Float Validators

| Validator | Description |
|-----------|-------------|
| `MinFloat(n)` | Minimum value |
| `MaxFloat(n)` | Maximum value |
| `GtFloat(n)` | Greater than |
| `GteFloat(n)` | Greater than or equal |
| `LtFloat(n)` | Less than |
| `LteFloat(n)` | Less than or equal |
| `IsPositiveFloat()` | Must be >= 0 |
| `IsNegativeFloat()` | Must be < 0 |

## 🎯 Real-World Example

```go
package main

import (
    "fmt"
    "github.com/mrbns/valgo/lib/v"
)

type User struct {
    Email    string
    Username string
    Age      int
    Balance  float64
}

func ValidateUser(user User) error {
    schema := v.NewPipesMap(map[string]v.Pipe{
        "email": v.NewStringPipe(
            user.Email,
            v.Empty(),
            v.IsEmail(),
            v.MaxLength(100),
        ),
        "username": v.NewStringPipe(
            user.Username,
            v.Empty(),
            v.MinLength(3),
            v.MaxLength(20),
            v.IsAlphaNumeric(),
        ),
        "age": v.NewNumberPipe(
            user.Age,
            v.Min(18, v.ErrMsg("Must be at least 18 years old")),
            v.Max(120),
        ),
        "balance": v.NewFloatPipe(
            user.Balance,
            v.MinFloat(0.0, v.ErrMsg("Balance cannot be negative")),
        ),
    })

    if err := schema.Validate(); err != nil {
        return fmt.Errorf("%s: %v", err.Key, err.Err)
    }

    return nil
}

func main() {
    user := User{
        Email:    "user@example.com",
        Username: "john_doe",
        Age:      25,
        Balance:  100.50,
    }

    if err := ValidateUser(user); err != nil {
        fmt.Println("Validation failed:", err)
    } else {
        fmt.Println("User is valid!")
    }
}
```

## 📝 License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## 🤝 Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## 📚 Documentation

For more detailed information about each validator, check the source code in:
- [`lib/v/string_actions.go`](lib/v/string_actions.go) - String validators
- [`lib/v/int_actions.go`](lib/v/int_actions.go) - Integer validators
- [`lib/v/float_actions.go`](lib/v/float_actions.go) - Float validators
- [`lib/is/string.go`](lib/is/string.go) - Low-level validation functions

## author
Mr Binary Sniper. @mrbns.