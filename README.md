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

- - -
**Work in progress. More packages to come.**