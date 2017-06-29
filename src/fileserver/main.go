// fileserver project main.go
package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe(":6060", http.FileServer(http.Dir(".")))
}
