package main

import "fmt"

func main() {
    var num string
    fmt.Scan(&num)
    dim := len(num)

    isValid := true
    for i := 0; i < dim; i++ {
        if (num[i]-'0')%2 != 0 {
            isValid = false
            break
        }
    }
    fmt.Println("risultato", isValid)
}
