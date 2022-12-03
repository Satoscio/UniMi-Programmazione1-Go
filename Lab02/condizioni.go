package main

import "fmt"

func main() {
    const VAL_LIMITE = 10.0
    var num int
    var numF float64

    fmt.Print("Uguale a 10: ")
    fmt.Scan(&num)
    fmt.Println(num == 10)

    fmt.Print("Diverso da 10: ")
    fmt.Scan(&num)
    fmt.Println(num != 10)

    fmt.Print("Diverso da 10 e 20: ")
    fmt.Scan(&num)
    fmt.Println(num != 10 && num != 20)

    fmt.Print("Diverso da 10 o 20: ")
    fmt.Scan(&num)
    fmt.Println(num != 10 || num != 20)

    fmt.Print("Maggiore o uguale a 10: ")
    fmt.Scan(&num)
    fmt.Println(num >= 10)

    fmt.Print("Compreso tra [10, 20]: ")
    fmt.Scan(&num)
    fmt.Println(num >= 10 && num <= 20)

    fmt.Print("Compreso tra (10, 20): ")
    fmt.Scan(&num)
    fmt.Println(num > 10 && num < 20)

    fmt.Print("Compreso tra [10, 20): ")
    fmt.Scan(&num)
    fmt.Println(num >= 10 && num < 20)

    fmt.Print("Esterno a [10, 20]: ")
    fmt.Scan(&num)
    fmt.Println(num < 10 || num > 20)

    fmt.Print("Compreso in [10, 20] o [30, 40]: ")
    fmt.Scan(&num)
    fmt.Println((num >= 10 && num <= 20) || (num >= 30 && num <= 40))

    fmt.Print("Multiplo di 4 ma non di 100: ")
    fmt.Scan(&num)
    fmt.Println(num % 4 == 0 && num % 100 != 0)

    fmt.Print("Dispari e compreso in [0, 100]: ")
    fmt.Scan(&num)
    fmt.Println(num % 2 == 1 && num >= 0 && num <= 100)

    fmt.Print("Float vicino a 10.0: ")
    tolleranza := 0.000001
    fmt.Scan(&numF)
    fmt.Println(numF < 10.0 + tolleranza && numF > 10.0 - tolleranza)
}
