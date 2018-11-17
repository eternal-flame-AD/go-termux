// Package termux implements termux:API calls and wraps them into golang function calls.
/*
	Example:
	package main

	import tm "github.com/eternal-flame-AD/go-termux"

	func main() {
		if err := tm.ClipboardSet("ummmm"); err != nil {
			panic(err)
		}
	}


	We call termux:API methods directly so that this package would work without termux-api package installed.
*/
package termux
