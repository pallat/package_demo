package main

import (
	"fmt"

	"github.com/pallat/skooldio/packagedemo/book"
)

func main() {
	book.Name = "Go to Gopher"
	fmt.Println(book.Name)
}
