package main

import "fmt"

func main() {

    //1.
    a := []int{1, 2, 3, 4}
    b := []int{5, 6, 7, 8, 9}
    fmt.Println("1. a, b:", a, b)

    //2.
    a = append(a, b...)
    fmt.Println("2. a:", a)

    // 3.
    b = make([]int, len(a))
    copy(b, a)
    fmt.Println("3. b:", b)

    i := 2
    j := 4
    //4.
    a = append(a[:i], a[i+1:]...)
    fmt.Println("4. a:", a)

    //5.
    a = append(a[:i], a[j:]...)
    fmt.Println("5. a:", a)

    //6.
    a = append(a, make([]int, j)...)
    fmt.Println("6. a:", a)

    x := 10
    //7.
    a = append(a[:i], append([]int{x}, a[i:]...)...)
    fmt.Println("7. a:", a)

    //8.
    a = append(a[:i], append(make([]int, j), a[i:]...)...)
    fmt.Println("8. a:", a)

    //9.
    a = append(a[:i], append(b, a[i:]...)...)
    fmt.Println("9. a:", a)

    //10.
    x, a = a[len(a)-1], a[:len(a)-1]
    fmt.Println("10. x, a:", x, a)

    //11.
    a = append(a, x)
    fmt.Println("11. a:", a)
}