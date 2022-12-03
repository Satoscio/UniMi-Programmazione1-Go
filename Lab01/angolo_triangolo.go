/*
Problema: Scrivere un programma Go angolo_triangolo.go che, date in input le ampiezze in
gradi (float64) di due angoli di un triangolo, determini l’ampiezza del terzo angolo.
Nota. La somma degli angoli di un triangolo è sempre 180 gradi. È quindi il caso di utilizzare
una costante per questo valore.
Esempio di esecuzione:
angolo uno e due: 25.3 90.2
ampiezza terzo angolo: 64.49999999999999
*/

package main

import "fmt"

func main() {
    const sommaAngoli = 180.0
    var angolo1, angolo2, angolo3 float64
    fmt.Print("Angolo 1 e 2: ")
    fmt.Scan(&angolo1, &angolo2)
    angolo3 = sommaAngoli - (angolo1 + angolo2)
    fmt.Println("Ampiezza terzo angolo:", angolo3)
}
