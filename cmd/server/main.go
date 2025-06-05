package main

import (
	"chapapp-backend-api/internal/router"
)

func main() {
	r := router.NewRouter()
	r.Run(":8080")
}
