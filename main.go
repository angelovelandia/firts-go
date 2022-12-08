package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("Hola mundo en GO")
	fmt.Println(uuid.New().String())
}
