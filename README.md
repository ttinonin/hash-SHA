# SHA CLI - Hash generator (SHA256, SHA384, SHA512)

This is a comand line tool made in Go to generate hases SHA-256, SHA-384 and SHA-512 from a given word

## Usage

### Build 

On your terminal run:

```bash
go buid ./sha256.go
```

### Executing

To get a hash of a given string run:

```bash
./sha256 "This is a test"
```

Without flags, the program will return SHA-256 by default:

```bash
SHA256: c7be1ed902fb8dd4d48997c6452f5d7e509fbcdbe2808b16bcf4edce4c07d14e
```
