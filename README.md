# go-termux

[![API Documentation](https://img.shields.io/badge/api-GoDoc-blue.svg?style=flat-square)](https://godoc.org/github.com/xconstruct/go-termux)

Golang wrapper for termux:API. This package calls termux:API methods directly so that this package would work without termux-api package installed.

## Examples

### Sets current clipboard content
```golang
package main

import tm "github.com/eternal-flame-AD/go-termux"

func main() {
    if err := tm.ClipboardSet("ummmm"); err != nil {
        panic(err)
    }
}
```