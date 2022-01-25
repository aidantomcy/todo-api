package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aidantomcy/todo-api/controller"
	"github.com/aidantomcy/todo-api/model"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	fmt.Println("Serving...")
	log.Fatal(http.ListenAndServe(":3000", mux))
}
