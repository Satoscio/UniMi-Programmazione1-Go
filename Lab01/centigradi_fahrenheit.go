/*
Problema: Scrivere un programma Go centigradi_fahrenheit.go che converta in gradi
Fahrenheit una temperatura espressa in gradi centigradi interi letta da standard input.
Nota. Un grado Fahrenheit è 5/9 di un grado centigrado e 0 gradi centigradi corrisponde a 32
gradi Fahrenheit. Usare costanti nella formula di conversione. Attenzione alla differenza in
Go tra 5/9 e 5./9.
Esempio di esecuzione:
temperatura in centigradi? 18
18 C = 64.4 F
*/

package main

import "fmt"

func main() {
    const F2C = 9./5.
    const ZEROF = 32.
    var tempCentigradi, tempFahrenheit float64
    fmt.Print("Temperatura in centigradi? ")
    fmt.Scan(&tempCentigradi)
    tempFahrenheit = float64(tempCentigradi) * F2C + ZEROF
    fmt.Println(tempCentigradi, "C =", tempFahrenheit, "F")
}
