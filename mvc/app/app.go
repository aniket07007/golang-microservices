package app

import(
	"net/http"
	"github.com/aniket07007/golang-microservices/mvc/controllers"
)

func StartApp() {
	http.HandleFunc("/user", controllers.GetUsers)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}