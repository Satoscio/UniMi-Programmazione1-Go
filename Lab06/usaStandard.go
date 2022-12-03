/*
    1. se la stringa letta contiene S2 come sottostringa
    2. se la stringa letta contiene almeno una vocale
    3. quante volte S1 si trova nella stringa letta
    4. pos della prima e ultima vocale nella stringa letta
    5. concatenare S2 3 volte con se stessa e stampare
    6. concatenare S1 5 volte con se stessa e stampare

    concatenare eventuali cifre contenute nella stringa in una stringa s

    7. stampare l'int n corrispondente a s con fmt.Printf("%d\n", n)
    8. stampare il float64 f corrispondente a s con fmt.Printf("%f\n", f)

*/

package main

import (
    "fmt"
    "strings"
    "strconv"
)

func contaVocali(s, voc string) {
    nv := 0
    for _, c := range s {
        for _, v := range voc {
            if c == v {
                nv++
            }
        }
    }
    fmt.Println(s, "ha", nv, "vocali")
}

func posVocali(s, voc string) {
    prmVoc := strings.IndexAny(s, voc)
    ultVoc := strings.LastIndexAny(s, voc)
    fmt.Println("La prima vocale di", s, "è in pos.", prmVoc)
    fmt.Println("L'ultima vocale di", s, "è in pos.", ultVoc)
}

func estraiNumero(par, cifre string) int {
    var s string
    for _, ch := range par {
        for _, c := range cifre {
            if ch == c {
                s += string(ch)
            }
        }
    }
    n, _ := strconv.Atoi(s)
    return n
}

func main() {
    const VOCALI, CIFRE, S1, S2 = "aeiouAEIOU", "0123456789", "r", "rt"
    var par string
    fmt.Scan(&par)

    // 1
    if strings.Contains(par, S2) {
        fmt.Println(par, "contiene", S2)
    } else {
        fmt.Println(par, "non contiene", S2)
    }

    // 2
    contaVocali(par, VOCALI)

    // 3
    occ := strings.Count(par, S1)
    fmt.Println(par, "ha", occ, S1)

    // 4
    posVocali(par, VOCALI)

    // 5
    fmt.Println(strings.Repeat(S2, 3))

    //6
    fmt.Println(strings.Repeat(S1, 5))

    //7
    n := estraiNumero(par, CIFRE)
    fmt.Printf("%d\n", n)

    //8

}
