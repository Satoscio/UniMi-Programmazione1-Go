/*
    per compilare
    go mod init server
    go mod tidy

    solo una volta
*/

package main

import (
    "github.com/holizz/terrapin"
    "fmt"
    "math"
    "net/http"
    "image"
    "image/png"
)

func pippoFunc(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte(`

        <!DOCTYPE html>
        <html>
            <h1 style=\" font-family: 'Monocraft'; \">Choba</h1>
        </html>

    `))
}

func mainPage(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte(`

        <!DOCTYPE html>
        <html>

            <head>
                <title>Buon Natale</title>
            </head>
            <body>
                <h1>Buon Natale!</h1>
                <p>
                    <img src="/mainImage.png">
                </p>
            </body>

        </html>

    `))
}

func mainImage(w http.ResponseWriter, r *http.Request) {
    campo := image.NewRGBA(image.Rect(0, 0, 500, 500))
    t := terrapin.NewTerrapin(campo, terrapin.Position{250.0, 450.0})
    kochCurve(t, 50, 3)
    png.Encode(w, campo)
}

func kochCurve(t *terrapin.Terrapin, lung float64, liv int) {
    if liv == 0 {
        t.Forward(100)
        return
    }
    
    t.Forward(lung)
    t.Left(math.Pi / 3.0)
    t.Forward(lung)
    t.Right(math.Pi - math.Pi / 3.0)
    t.Forward(lung)
    t.Left(math.Pi / 3.0)
    t.Forward(lung)
}

func main() {
    fmt.Println("Listening on http://localhost:3000/")
    http.HandleFunc("/pippo", pippoFunc)
    http.HandleFunc("/main", mainPage)
    http.HandleFunc("/mainImage", mainImage)
    http.ListenAndServe(":3000", nil)
}
