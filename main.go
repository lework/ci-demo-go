package main

import (
    "log"
    "net/http"
    "os"
    "strings"
)

var version = "master"
var build = ""

func getNameLen(name string) int {
    return len(name)
}

func showVersion(w http.ResponseWriter, r *http.Request) {
    message := "Version: " + version + " Build: " + build
    log.Println(message)
    w.Write([]byte(message))
}

func sayHello(w http.ResponseWriter, r *http.Request) {
    message := r.URL.Path
    message = strings.TrimPrefix(message, "/")
    message = "Hello, you got the message: " + message
    log.Println(message)
    w.Write([]byte(message))
}

func main() {
    // use PORT environment variable, or default to 80
    port := "80"
    if fromEnv := os.Getenv("PORT"); fromEnv != "" {
        port = fromEnv
    }
    http.HandleFunc("/version", showVersion)
    http.HandleFunc("/", sayHello)
    log.Println("Listen server on " + port + " port")
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatal(err)
    }
}
