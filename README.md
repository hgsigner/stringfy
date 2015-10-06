#Stringfy [![Build Status](https://travis-ci.org/hgsigner/stringfy.svg?branch=master)](https://travis-ci.org/hgsigner/stringfy) [![GoDoc](https://godoc.org/github.com/hgsigner/stringfy?status.svg)](https://godoc.org/github.com/hgsigner/stringfy)

Stringfy is a collection of string manipulation packages for GO.
- - -

##Install:

```bash
$ go get github.com/hgsigner/stringfy
```

##Usage:

```go
package main

import (
  	"fmt"
	"github.com/hgsigner/stringfy"
)

func main() {
	es := stringfy.Escape("São Paulo")
  	fmt.Prinln(es) // Sao Paulo

  	cc := stringfy.CamelCase("fizz_buzz_bazz")
  	fmt.Prinln(cc) // FizzBuzzBazz

	un := stringfy.Parameterize("São Paulo")
 	fmt.Prinln(un) // sao-paulo

 	un := stringfy.Underscore("São Paulo")
 	fmt.Prinln(un) // sao_paulo
 	
 	// Default plural
	pl1 := stringfy.NewPlural()
 	pl1.Perform(2, "octopus")
 	fmt.Println(pl1) // 2 octopi
 	
 	// Custom plural
 	pl2 := stringfy.NewPlural()
 	pl2.Options(stringfy.AddPlural("timesz"))
 	pl2.Perform(2, "timey")
 	fmt.Println(pl2) // 2 timesz
}
```

- - -
For more information, please refer to the [docs.](https://godoc.org/github.com/hgsigner/stringfy) **Work in progress. More packages to come.**
- - -
##Licensing
You can use or modify it to meet your needs.