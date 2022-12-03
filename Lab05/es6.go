package main

import (
   "fmt"
   "unicode"
)

func main() {

   var word string

   fmt.Print("una parola: ")
   fmt.Scan(&word)

   var cV, cAV, cC int
   for _, ch := range word {
      if unicode.IsLetter(ch) {
         switch ch {
         case 'a', 'e', 'i', 'o', 'u':
            cV++
         case 'à', 'è', 'é', 'ì', 'ò', 'ù':
            cAV++
         default:
            cC++
         }
      }
   }
   fmt.Println("1. Vocali:", cV)
   fmt.Println("2. Vocali accentate:", cAV)
   fmt.Println("3. Consonanti:", cC)
}
