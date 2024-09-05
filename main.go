package main

import (
        "html/template"
        "net/http"
        "fmt"
)

type CountdownInfo struct {
        DestinationEvent string
}

func main() {

        // compile the template
        template := template.Must(template.ParseFiles(`./countdown-timer.html`))
        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                template.Execute(w, CountdownInfo{})
        })
	fmt.Println("Starting countdown web server.")
        http.ListenAndServe(":8080", nil)
}

