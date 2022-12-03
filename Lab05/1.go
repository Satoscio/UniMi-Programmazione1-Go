package main

import "fmt"

func averageCh(num string) (average byte) {
   dim := len(num)
   var sum int
   for i := 0; i < dim; i++ {
      n := int(num[i] - '0')
      sum += n
   }
   average = byte(sum/dim) + '0'
   return
}

func countGreater(num string, ch byte) int {
   dim := len(num)
   count := 0
   for i := 0; i < dim; i++ {
      if num[i] > ch {
         count++
      }
   }
   return count
}

func main() {
   var numero string
   fmt.Scan(&numero)

   media := averageCh(numero)
   fmt.Println("media", string(media))
   fmt.Println("conteggio", countGreater(numero, media))
}
