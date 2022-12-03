package main

import (
    "fmt"
    "os"
    "strconv"
    "time"
)

func attendi(n_seconds float64){
   time.Sleep(time.Duration(n_seconds) * time.Second)
}

func printRoba(x int, r rune) {
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
    for r := 0; r < dim; r++ {
        printRoba(r, ' ')
        fmt.Print("*")
        if r >= dim - sopra {
            printRoba(2 * (dim - 1) + 2 - 2 * r, '*')
        } else {
            printRoba(2 * (dim - 1) + 2 - 2 * r, ' ')
        }
        fmt.Println("*")
    }

    for r := dim - 1; r >= 0; r-- {
        printRoba(r, ' ')
        fmt.Print("*")
        if r > dim - sotto {
            printRoba(2 * (dim - 1) + 2 - 2 * r, '*')
        } else {
            printRoba(2 * (dim - 1) + 2 - 2 * r, ' ')
        }

        fmt.Println("*")
    }
}

func main() {
    var sec int
    var err error

    if len(os.Args) != 2 {
        fmt.Println("Numero di argomenti sbagliato!")
        os.Exit(1)
    }

    sec, err = strconv.Atoi(os.Args[1])

    if err != nil {
        fmt.Printf("Valore non accettato %s\n", os.Args[1])
        os.Exit(1)
    }

    for s := 0; s <= sec; s++ {
        disegnaClessidra(sec, sec - s, s)
        fmt.Println()
    }


}
