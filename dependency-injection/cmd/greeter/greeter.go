package main

import (
	"fmt"
	"github.com/jspaaks/learn-go-with-tests/dependency-injection/pkg/greeter"
	"log"
	"net/http"
	"os"
)

func main(){
	greeter.Greet(os.Stdout, "Elodie")
	fmt.Println("Starting the server... Ctrl-c to exit")
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	greeter.Greet(w, "world")
}
