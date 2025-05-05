package main

import (
    "fmt"
    "github.com/jspaaks/learn-go-with-tests/select/pkg/racer"
    "time"
)

func main() {
    url0 := "https://google.com"
    url1 := "https://fairsoftwarechecklist.net/v0.2"
    fastest, err := racer.Racer(url0, url1, 3 * time.Second)
    if err != nil {
        fmt.Println(err)
    }
    if url0 == fastest {
        fmt.Printf("%q responded faster than %q\n", url0, url1)
    } else {
        fmt.Printf("%q responded faster than %q\n", url1, url0)
    }
}
