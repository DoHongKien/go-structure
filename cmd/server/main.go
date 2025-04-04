package main

import (
	"github.com/DoHongKien/go-structure/internal/initialize"
)

func main() {

	r := initialize.Run()

	r.Run(":9999")
}
