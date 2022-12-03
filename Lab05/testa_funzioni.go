package main

import (
    "fmt"
    "bufio"
    "os"
)

func hasUpper(s string) bool {
    for i := 0; i < len(s); i++ {
        if s[i] >= 'A' && s[i] <= 'Z' {
            return true
        }
    }
    return false
}

func firstUpper(s string) int {
    for i := 0; i < len(s); i++ {
        if s[i] >= 'A' && s[i] <= 'Z' {
            return i
        }
    }
    return -1
}

func lastUpper(s string) int {
    for i := len(s)-1; i >= 0; i-- {
        if s[i] >= 'A' && s[i] <= 'Z' {
            return i
        }
    }
    return -1
}

func countDigitsLettersOthers(s string) (int, int, int) {
    var dig, let, oth int

    for i := 0; i < len(s); i++ {
        if s[i] >= 'A' && s[i] <= 'Z' || s[i] >= 'a' && s[i] <= 'z' {
            let++
        } else if s[i] >= '0' && s[i] <= '9' {
            dig++
        } else {
            oth++
        }
    }
    return dig, let, oth
}
/*
// s l-1-i
func isPalindrome(s string) bool {
    for i := 0; i < len(s); i++ {
        l := len(s)
        if s[i] != s[l-1-i] {
            return false
        }
    }
    return true
}
*/
func main() {
    in := bufio.NewReader(os.Stdin)
    s, _ := in.ReadString('\n')
    fmt.Println("Ha maiuscole:", hasUpper(s))
    fmt.Println("Prima maiuscola in posizione:", firstUpper(s))
    fmt.Println("Ultima maiuscola in posizione:", lastUpper(s))
    dig, let, oth := countDigitsLettersOthers(s)
    fmt.Println("Cifre:", dig, "Lettere:", let, "Altri:", oth-1)
    //fmt.Println("È palindroma:", isPalindrome(s))
}
