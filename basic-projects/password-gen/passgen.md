# Password Generator - Learning Notes

## Project Overview

### Generate Password
```bash
go run . --length=20 --special=true --small=true --large=true --numbers=true
```

**Output:**
```
Here you go : aB3!xY9@mN7#pQ2$rT5%
```

---

## crypto/rand vs math/rand

### math/rand
```go
import "math/rand"

rand.Seed(time.Now().UnixNano())  // Predictable seed
n := rand.Intn(10)                // Predictable output
```
- **Predictable** - Can be guessed
- **Not cryptographically secure**
- Use for: Games, simulations

### crypto/rand (Secure)
```go
import "crypto/rand"
import "math/big"

randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
```
- **Unpredictable** - Cannot be guessed
- **Cryptographically secure**
- Use for: Passwords, tokens, keys

---

## How crypto/rand Works

### rand.Reader
- OS-level random source
- `/dev/urandom` (Unix)
- True randomness from system entropy

### rand.Int()
```go
rand.Int(rand.Reader, max *big.Int) (*big.Int, error)
```
- Returns random number: `0 <= n < max`
- Uses `big.Int` for large numbers
- Returns error if random source fails

---

## Key Patterns

### Build Character Set
```go
charset := ""
if includeSpecial { charset += "!@#$%^&*()_+" }
if includeNumbers { charset += "0123456789" }
if includeLargeCase { charset += "ABCDEFGHIJKLMNOPQRSTUVWXYZ" }
if includeSmallCase { charset += "abcdefghijklmnopqrstuvwxyz" }
```

## CLI Flags

```go
flag.Int("length", 16, "Password Length")
flag.Bool("special", true, "Include Special Characters")
flag.Parse()
```

**Usage:**
```bash
go run . --length=20 --special=false
```

## Key Learnings

1. **crypto/rand** for security, **math/rand** for games
2. **Slices grow with append**: `password = append(password, char)`
3. **Blank identifier** `_` for unused loop variables
4. **Flag package** for CLI arguments
5. **Error handling** at function level, not just UI
6. **strings.ContainsAny** checks any character from set

---

## Project Structure

```
password-gen/
├── main.go           → CLI entry point
├── generator.go      → Core logic
└── generator_test.go → Unit tests
```

---

## Security Notes

-  Uses `crypto/rand` (secure)
-  Validates input length
-  Checks for empty charset
-  Returns errors properly
