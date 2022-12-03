# Heading
## mi sono cagato nei pantaloni
### ciaoo

``` go
package main

import "fmt"

func main() {
    var x, y, ris int
    fmt.Scan(&x)
    fmt.Scan(&y)
    ris = f(x, y)
    fmt.Println(ris)
}

func f(a, b int) (c int) {
    var x, y int
    x = sqr(a)
    y = sqr(b)
    c = x + y + 1
    return
}

func sqr(x int) (a int) {
    a = x * x
    return
}

```
