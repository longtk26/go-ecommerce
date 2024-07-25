package main

import "github.com/longtk26/go-ecommerce/internal/routers"

func main() {
	routers := routers.NewRouter()	

	routers.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
