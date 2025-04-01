package main

import (
	"github.com/DoHongKien/go-structure/internal/router"
)

func main() {

	// router.NewRouter()

	r := router.NewRouter()

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
