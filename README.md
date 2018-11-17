# go-termux

[![API Documentation](https://img.shields.io/badge/api-GoDoc-blue.svg?style=flat-square)](https://godoc.org/github.com/eternal-flame-AD/go-termux)

Golang wrapper for termux:API. This package calls termux:API methods directly so that this package would work without termux-api package installed.

## Examples

### Acquires the current battery percentage

```golang
package main

import (
        "fmt"

        tm "github.com/eternal-flame-AD/go-termux"
)

func main() {
        if stat, err := tm.BatteryStatus(); err != nil {
                panic(err)
        } else {
                fmt.Printf("The current battery percentage is %d%%.\n", stat.Percentage)
        }
}
```

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

### Displays a short toast
```golang
package main

import (
        tm "github.com/eternal-flame-AD/go-termux"
)

func main() {
        if err := tm.Toast("Hello World!", tm.ToastOption{
                FontColor: "#FF0000",
                Position:  tm.Top,
                Short:     true,
                BGColor:   "#00FF00",
        }); err != nil {
                panic(err)
        }
}
```