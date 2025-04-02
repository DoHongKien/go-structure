package main

import (
	"github.com/DoHongKien/go-structure/internal/router"
)

func main() {

	r := router.NewRouter()

	r.Run(":9999")
}
