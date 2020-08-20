# freeport is a simple Go package for finding free TCP ports

This is a more up-to-date hard fork of https://github.com/phayes/freeport

It has https://github.com/phayes/freeport/pull/8 merged, the CLI removed (Go package only), and other minor improvements.

PRs welcome and appreciated.

## Usage

```go
package main

import "github.com/slimsag/freeport"

func main() {
	port, err := freeport.GetFreePort()
	if err != nil {
		log.Fatal(err)
	}
	// port should be ready to listen on
}

```
