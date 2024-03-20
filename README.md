# Passwordify
Generate a complex password from a cli.

---

## How to setup

1. Clone the repo
2. `go build pswify.go`
3. Add pswify binary to your path
4. Have fun!


## Usage
Basic password with 12 characters and no numbers.
```bash
pswify gen
```
Password with 16 characters and 4 numbers.
```
pswify -char 16 -num 4 gen
```