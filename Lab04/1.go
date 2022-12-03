package main
import  "fmt"

func main() {
    var num string
    fmt.Scan(&num)
    dim := len(num)

    m := num[0]
    for i := 1; i < dim; i++ {
        if num[i] < m {
            m = num[i]
        }
    }
    fmt.Println(string(m))
}
