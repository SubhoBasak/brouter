package main

import (
	"net/http"
	"test/routers"

	"github.com/SubhoBasak/brouter"
)

func main() {
	router := brouter.Router{}

	router.Router("/blog", routers.BlogRouter())
	router.Router("/post", routers.PostRouter())

	http.ListenAndServe("127.0.0.1:5000", &router)
}
