package main


import (
	_ "embed"
	"fmt"
)
var (
	//go:embed common/empty.creds
	DefaultJwt []byte
)


func main() {
	fmt.Println("jwt is:", string(DefaultJwt))
}
