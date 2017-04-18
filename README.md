# Stringfy [![Build Status](https://travis-ci.org/hgsigner/stringfy.svg?branch=master)](https://travis-ci.org/hgsigner/stringfy) [![GoDoc](https://godoc.org/github.com/hgsigner/stringfy?status.svg)](https://godoc.org/github.com/hgsigner/stringfy)

Stringfy is a string manipulation package for GO.
- - -

## Install:

```bash
$ go get github.com/hgsigner/stringfy
```

## Usage:

```go
package main

import (
  	"fmt"
	"github.com/hgsigner/stringfy"
)

func main() {
	es := stringfy.Escape("São Paulo")
 	fmt.Prinln(es) // Sao Paulo
}
```
## Functions:

### CamelCase:

```go
package main

import (
  	"fmt"
	"github.com/hgsigner/stringfy"
)

func main() {
	cc := stringfy.CamelCase("fizz_buzz_bazz")  
	fmt.Prinln(cc) // FizzBuzzBazz
}
```
- - -
### Escape:

```go
package main

import (
  	"fmt"
	"github.com/hgsigner/stringfy"
)

func main() {
	es := stringfy.Escape("São Paulo")
 	fmt.Prinln(es) // Sao Paulo
}
```
- - -
### Parameterize:

```go
package main

import (
  	"fmt"
	"github.com/hgsigner/stringfy"
)

func main() {
	pm := stringfy.Parameterize("São Paulo")
 	fmt.Prinln(pm) // sao-paulo
}
```
- - -
### Underscore:

```go
package main

import (
  	"fmt"
	"github.com/hgsigner/stringfy"
)

func main() {
	un := stringfy.Underscore("São Paulo")
 	fmt.Prinln(un) // sao_paulo
}
```
- - -
### Plural:

```go
package main

import (
  	"fmt"
	"github.com/hgsigner/stringfy"
)

func main() {
	// Default plural
	pl1 := stringfy.NewPlural()
 	fmt.Println(pl1.Perform(2, "octopus")) // 2 octopi
 	
 	// Custom plural
 	pl2 := stringfy.NewPlural()
 	pl2.Options(stringfy.AddPlural("timesz"))
 	fmt.Println(pl2.Perform(2, "timey")) // 2 timesz
}
```

### Options:

```go
stringfy.AddPlural(string)
```
- - -
### Truncate:

```go
package main

import (
  	"fmt"
	"github.com/hgsigner/stringfy"
)

func main() {
	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit."
 	
 	// Default truncate
 	tr1 := stringfy.NewTruncate()
 	tr1.Perform(text) // Lorem ipsum dolor sit amet,...

 	// Custom truncate
 	tr2 := stringfy.NewTruncate()
 	tr2.Options(stringfy.AddLength(17))
 	tr2.Perform(text) // Lorem ipsum do...

 	tr3 := stringfy.NewTruncate()
 	tr3.Options(stringfy.AddLength(13), stringfy.AddSeparator(" "))
 	tr3.Perform(text) // Lorem ipsum...
}
```

### Options:

```go
stringfy.AddLength(int)       // Default: 30

stringfy.AddOmission(string)  // Default: "..."

stringfy.AddSeparator(string) // Default: ""
```

- - -
### Excerpt:

```go
package main

import (
  	"fmt"
	"github.com/hgsigner/stringfy"
)

func main() {
	text := "Lorem ipsum dolor sit amet consectetur adipiscing elit."
 	
 	ex1 := stringfy.NewExcerpt()
 	ex1.Options(stringfy.AddRadius(4))
 	ex1.Perform(text, "sit") // ...lor sit ame...

 	ex2 := stringfy.NewExcerpt()
 	ex2.Options(stringfy.AddRadius(4), stringfy.AddOmission("(...)"))
 	ex2.Perform(text, "sit") // (...)lor sit ame(...)

 	ex3 := stringfy.NewExcerpt()
 	ex3.Options(stringfy.AddRadius(1), stringfy.AddSeparator(" "))
 	ex3.Perform(text, "sit") // ...dolor sit amet...
}
```

### Options:

```go
stringfy.AddRadius(int)       // Default: 100

stringfy.AddOmission(string)  // Default: "..."

stringfy.AddSeparator(string) // Default: ""
```

### WordWrap:

```go
package main

import (
  	"fmt"
	"github.com/hgsigner/stringfy"
)

func main() {
	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit sil."

	wwr1 := stringfy.NewWordWrap()
 	wwr1.Perform(text) // "Lorem ipsum dolor sit amet, consectetur adipiscing elit sil."

 	wwr2 := stringfy.NewWordWrap()
 	wwr2.Options(stringfy.AddLineWidth(10))
 	wwr2.Perform(text) // "Lorem\nipsum\ndolor sit\namet,\nconsectetur\nadipiscing\nelit sil."

 	wwr3 := stringfy.NewWordWrap()
 	wwr3.Options(stringfy.AddLineWidth(35))
 	wwr3.Perform(text) // "Lorem ipsum dolor sit amet,\nconsectetur adipiscing elit sil."
}
```

### Options:

```go
stringfy.AddLineWidth(int)       // Default: 80
```

- - -
For more information, please refer to the [docs.](https://godoc.org/github.com/hgsigner/stringfy) **Work in progress. More packages to come.**
- - -
## Licensing
This package is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).
