# targz

Simple package for creating `*tar.gz` files.

## Usage

```Go
package main

import (
    "log"

    "github.com/kelseyhightower/targz"
)

func main() {
    err := targz.Create("/Users/kelseyhightower/stuff", "stuff.tar.gz")
    if err != nil {
        log.Fatalln(err)
    }
}
```
