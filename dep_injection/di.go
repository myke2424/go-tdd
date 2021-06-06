package main

// Dependency Injection!
// If you can't test a function easily, it's usually because of dependencies hard-wired into a function or global state
// We can inject dependencies via an interface
// Using an interface makes our code easier to test


// If you ever feel like a method/function has too many responseibilities, e.g. generating data and writing to a db
// Or handling HTTP requests and doing domain level logic.... DI is probably the tool we need here.
// It allows our code to be re-used in different contexts.
// We can inject any depdency we want as long as it conforms to the interface

import (
	"fmt"
	"io"
  //  "log"
    "net/http"
)

// We're decoupling where the data goes from how to generate it in our greet function with the I.O Writer interface
func Greet(writer io.Writer, name string) {
	// Fprintf is like printf but instead takes a writer to send the string to
	// If we used regular printf, we couldnt test this because it writes to stdout
	// Here we are writing to our buffer which is testable
	fmt.Fprintf(writer, "Hello, %s", name)
}

// We use the ResponseWriter to write our response.
// http.ResponseWriter implements io.Writer so that is why we can re-use our Greet function here
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
    Greet(w, "world")
}

func main() {
 //    log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler)))
 db := MySQL{Port: 3000, Data: map[string]string{}}
     db.Persist("name", "mike")
     db.Persist("age", "25")
     fmt.Printf("data: '%v'", db.Data)
}


type Database interface {
    Persist(data ...interface{}) error
}


type MySQL struct {
     Port int
     Data map[string]string
}

func (m MySQL) Persist(key, value string) error {
   m.Data[key] = value
   return nil
}
