package main

import (
	"fmt"
	"github.com/mohammadDV/backend_gin/routes"
)

func main() {
	fmt.Println("test")

	r := routes.Urls()

	r.Run()

}
