package main

import (
    "fmt"
    "os"
    "strconv"
)

func printRoba(x, r int) {
    for i := 0; i < x; i++ {
        fmt.Printf("%c", r)
    }
}

/*
    Disegna una clessidra di asterischi con dim righe (sopra e sotto): nella
    parte superiore un certo numero di righe, le più in basso, (chiamato sopra)
    è piena e le restanti sono vuote; nella parte inferiore un certo numero di
    righe, le più in basso, (chiamato sotto) è piena e il restanti sono vuote.
*/

func disegnaClessidra(dim, sopra, sotto int) {

}

func main() {
    var op1, op2, ris float64
    var err error

    if len(os.Args) != 4 {
        fmt.Println("Numero di argomenti sbagliato!")
        os.Exit(1)
    }

    op1, err = strconv.ParseFloat(os.Args[1], 64)

    if err != nil {
        fmt.Printf("Valore non accettato %s\n", os.Args[1])
        os.Exit(1)
    }

    op2, err = strconv.ParseFloat(os.Args[3], 64)

    if err != nil {
        fmt.Printf("Valore non accettato %s\n", os.Args[3])
        os.Exit(1)
    }

    operazione := os.Args[2]

    switch operazione {
        case "+": ris = op1 + op2
        case "-": ris = op1 - op2
        case "x": ris = op1 * op2
        case "/": ris = op1 / op2
        default:
            fmt.Printf("Operazione %s non valida!", os.Args[2])
    }

    fmt.Printf("%.2f %s %.2f = %.2f\n", op1, os.Args[2], op2, ris)
}
