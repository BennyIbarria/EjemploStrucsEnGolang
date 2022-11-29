package mypackage

import "fmt"

//CaPublic es publica
type CarPublic struct {
	Brand string
	Year  int
}

func PrintMessaje() {
	fmt.Println("Hola desde mypackage")
}
