package main

import (
	"jena/internal/website"
)

func main() {
	host := "localhost"
	var port uint16 = 8080

	website.Start(host, port)
}
