package main

import "github.com/BaoTo12/go-ecommerce/internal/routers"

func main() {
	r := routers.NewRouter()
	r.Run("127.0.0.1:8080")
}
