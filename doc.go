//Simple string manipulation for Go.
//
// Example of usage:
//
//  package main
//
//  import (
//    "fmt"
//    "github.com/hgsigner/stringfy"
//  )
//
//  func main() {
//    es := stringfy.Escape("São Paulo")
//    fmt.Prinln(es) // Sao Paulo
//
//    cc := stringfy.CamelCase("fizz_buzz_bazz")
//    fmt.Prinln(cc) // FizzBuzzBazz
//
//    un := stringfy.Parameterize("São Paulo")
//    fmt.Prinln(un) // sao-paulo
//
//    un := underscore.PerformOn("São Paulo")
//    fmt.Prinln(un) // sao_paulo
//  }
package stringfy
