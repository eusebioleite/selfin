package main

import (
	"fmt"

	"github.com/eusebioleite/selfin/security"
)

func main() {
	hash, _ := security.HashPassword("admin")
	fmt.Println(hash)

}
