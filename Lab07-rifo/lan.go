package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    var maxs string
    var max int

    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        if len(scanner.Text()) > max {
            max = len(scanner.Text())
            maxs = scanner.Text()
        }
    }
    fmt.Println(maxs)
}
