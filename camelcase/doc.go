// It formats strings to CamalCase, escaping special characters from the text.
// E.g.
// São Paulo => SaoPaulo
// fizz-buzz-bazz => FizzBuzzBazz
// fizz_buzz_bazz => FizzBuzzBazz
//
// Example of usage:
//
//  package main
//
//  import (
//    "fmt"
//    "github.com/hgsigner/stringfy/camelcase"
//  )
//
//  func main() {
//    cc := camelcase.PerformOn("fizz_buzz_bazz")
//    fmt.Prinln(cc) // FizzBuzzBazz
//  }
package camelcase
