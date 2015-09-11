// It escapes special characters from a text.
// E.g. São Paulo => Sao Paulo
//
// Example of usage:
//
//  package main
//
//  import (
//    "fmt"
//    "github.com/hgsigner/stringfy/escaper"
//  )
//
//  func main() {
//    es := escaper.PerformOn("São Paulo")
//    fmt.Prinln(es) // Sao Paulo
//  }
package escaper
