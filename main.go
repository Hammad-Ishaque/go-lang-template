package main

import (
	"net/http"

	"./controller"
	"./model"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	http.ListenAndServe(":3000", mux)
}

// func main() {
//     fmt.Println("Hello, World!")
// }
