// The code is provided under the "MIT No Attribution" license.
//  See more here: https://github.com/aws/mit-0.
package main

import (
"fmt"
"net/http"

    "github.com/go-chi/chi"

)

func view(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain")
    w.Write([]byte("version"))
}

const PORT string = "9000"

func main() {
    r := chi.NewRouter()

    r.Route("/version", func(r chi.Router) {
      r.Get("/", view)
    })

    fmt.Printf("Server UP on port %s\n", PORT)
    http.ListenAndServe(fmt.Sprintf(":%s", PORT), r)
}

