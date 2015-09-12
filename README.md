#Stringfy [![Build Status](https://travis-ci.org/hgsigner/stringfy.svg?branch=master)](https://travis-ci.org/hgsigner/stringfy) [![GoDoc](https://godoc.org/github.com/hgsigner/stringfy?status.svg)](https://godoc.org/github.com/hgsigner/stringfy)

Stringfy is a collection of string manipulation packages for GO.
- - -

##Install:

```bash
$ go get github.com/hgsigner/stringfy
```

##escaper

```go
package main

import (
  "fmt"
  "github.com/hgsigner/stringfy/escaper"
)

func main() {
  es := escaper.PerformOn("São Paulo")
  fmt.Prinln(es) // Sao Paulo
}
```

##underscore:

```go
package main

import( 
	"fmt"
	"github.com/hgsigner/stringfy/underscore"
)

func main() {
	un := underscore.PerformOn("São Paulo")
	fmt.Println(un) // sao_paulo
}
```

##parameterize:

```go
package main

import( 
  "fmt"
  "github.com/hgsigner/stringfy/parameterize"
)

func main() {
  pr := parameterize.PerformOn("São Paulo")
  fmt.Println(pr) // sao-paulo
}
```

##pluralize:

```go
package main

import( 
	"fmt"
	"github.com/hgsigner/stringfy/pluralize"
)

func main() {
  //Default plural
  p1 := pluralize.New()
  p1.Perform(2, "octopus")
  fmt.Println(p1) // 2 octopi
  
  //Custom plural
  p2 := pluralize.New()
  p2.Options(pluralize.AddPlural("boatys"))
  p2.Perform(2, "boat")
  fmt.Println(p2) // 2 boatys
}
```

##camelcase

```go
package main

import (
 "fmt"
 "github.com/hgsigner/stringfy/camelcase"
)

func main() {
 cc := camelcase.PerformOn("fizz_buzz_bazz")
 fmt.Prinln(cc) // FizzBuzzBazz
}
```

- - -
For more information, please refer to the [docs.](https://godoc.org/github.com/hgsigner/stringfy) **Work in progress. More packages to come.**
- - -
##Licensing
You can use or modify it to meet your needs.