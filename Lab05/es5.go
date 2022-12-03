package main

import "fmt"

func main() {

   const MENU = "Menu del giorno:\n" +
      "a. pizza\n" +
      "b. penne al pomodoro\n" +
      "c. cotoletta e patatine\n" +
      "d. crostata e caffe`\n" +
      "f. fine"

   fmt.Println(MENU)

   var scelta, ordinazione string

myLoop:
   for {
      fmt.Println("ordinazione?")
      fmt.Scan(&scelta)
      switch scelta {
      case "a":
         ordinazione += " pizza"
      case "b":
         ordinazione += " penne al pomodoro"
      case "c":
         ordinazione += " cotoletta e patatine"
      case "d":
         ordinazione += " crostata e caffe`"
      case "f":
         fmt.Println("grazie, hai ordinato", ordinazione)
         break myLoop
      default:
         fmt.Println("ordinazione non valida")
      }
   }
}
