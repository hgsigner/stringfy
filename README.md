#Stringfy [![Build Status](https://travis-ci.org/hgsigner/stringfy.svg?branch=master)](https://travis-ci.org/hgsigner/stringfy) [![GoDoc](https://godoc.org/github.com/hgsigner/stringfy?status.svg)](https://godoc.org/github.com/hgsigner/stringfy)

Stringfy is a string manipulation package for GO.
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

	pm := stringfy.Parameterize("São Paulo")
 	fmt.Prinln(pm) // sao-paulo

 	un := stringfy.Underscore("São Paulo")
 	fmt.Prinln(un) // sao_paulo
 	
 	// Default plural
	pl1 := stringfy.NewPlural()
 	fmt.Println(pl1.Perform(2, "octopus")) // 2 octopi
 	
 	// Custom plural
 	pl2 := stringfy.NewPlural()
 	pl2.Options(stringfy.AddPlural("timesz"))
 	fmt.Println(pl2.Perform(2, "timey")) // 2 timesz

 	// Default truncate
 	tr1 := stringfy.NewTruncate()
 	tr1.Perform("Lorem ipsum dolor sit amet, consectetur adipiscing elit.") // Lorem ipsum dolor sit amet,...

 	// Custom truncate
 	tr2 := stringfy.NewTruncate()
 	tr2.Options(stringfy.AddLength(17))
 	tr2.Perform("Lorem ipsum dolor sit amet, consectetur adipiscing elit.") // Lorem ipsum do...

 	tr3 := stringfy.NewTruncate()
 	tr3.Options(stringfy.AddLength(13), stringfy.AddSeparator(" "))
 	tr3.Perform("Lorem ipsum dolor sit amet, consectetur adipiscing elit.") // Lorem ipsum...
}
```

##Options
**Truncate**

```go
stringfy.AddLength(int)       // Default: 30

stringfy.AddOmission(string)  // Default: "..."

stringfy.AddSeparator(string) // Default: ""
```

- - -
For more information, please refer to the [docs.](https://godoc.org/github.com/hgsigner/stringfy) **Work in progress. More packages to come.**
- - -
##Licensing
This package is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).
