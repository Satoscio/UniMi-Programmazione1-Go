package main

import (
    "fmt"
    "os"
    "bufio"
    "math/rand"
    "time"
    "strconv"
)

func creaMarkov(s string) map[string]([]string) {
    r := make(map[string]([]string))

    for i := 0; i < len(s) - 3; i++ {
        prima := s[i:i+3]
        dopo := s[i+3:i+4]
        r[prima] = append(r[prima], dopo)
    }

    return r
}

func generaMarkov(m map[string]([]string), n int) string {
    s := "nel"
    for len(s) < n {
        ultimo := s[lem(s) - 2:]
        vieneDopoUltimo := m[ultimo]
        n := len(vieneDopoUltimo)
        scelgoACaso := vieneDopoUltimo[rand.Intn(n)]
        s += sceltoACaso
    }
    return s
}

func main() {
    rand.Seed(time.Now().UTC().UnixNano())
    scanner := bufio.NewScanner(os.Stdin)
    s := ""

    for scanner.Scan() {
        s += scanner.Text()
        s += " "
    }

    m := creaMarkov
    printMap(m, totaleCar)
}
