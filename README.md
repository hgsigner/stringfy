#Stringfy

Stringfy is a collection of string manipulation packages for GO.
- - -

##Install:

```bash
$ go get github.com/hgsigner/stringfy
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
	fmt.Prinln(un) // sao_paulo
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
  //Defual plural
  p1 := pluralize.New()
  p1.Perform(2, "octopus")
  fmt.Prinln(p1) // 2 octopi
	
	//Custom plural
  p2 := pluralize.New()
  p2.Options(pluralize.AddPlural("boatys"))
  p2.Perform(2, "boat")
  fmt.Prinln(p2) // 2 boatys
}
```

- - -
**Work in progress. More packages to come.**