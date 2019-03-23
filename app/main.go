package main

import (
	"fmt"
	"net/http"
)



func main() {

	http.HandleFunc("/", index_handler)

	http.ListenAndServe(":8000", nil)
}


func index_handler(writer http.ResponseWriter,  request *http.Request ) {

	fmt.Fprintf(writer, "Go is awesome run from a container!")

}

