# go-otp - simple generator of One-Time Passwords

[![GoDoc](https://godoc.org/github.com/vlaship/go-otp?status.svg)](https://godoc.org/github.com/vlaship/go-otp)
[![Go Report Card](https://goreportcard.com/badge/github.com/vlaship/go-otp)](https://goreportcard.com/report/github.com/vlaship/go-otp)
[![Build Status](https://github.com/vlaship/go-otp/workflows/build/badge.svg)](https://github.com/vlaship/go-otp/actions)
[![Coverage Status](https://coveralls.io/repos/github/vlaship/go-otp/badge.svg?branch=main)](https://coveralls.io/github/vlaship/go-otp?branch=main)

### Installing

To start using go-otp, install Go and run `go get`:

```bash
go get github.com/vlaship/go-otp
```
This will retrieve the library.

## Usage

Import the package into your project.

```go
import "github.com/vlaship/go-otp"
```

Use it like this:

```go
o := otp.Generate()
```

This will generate a new OTP.
