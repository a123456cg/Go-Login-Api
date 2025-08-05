package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/a123456cg/Go-Login-Api/handlers" // 替換成你的 go module 名稱
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "✅ Go Login API 啟動成功！")
	}).Methods("GET")

	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST") // 🔥 新增這行

	fmt.Println("🚀 Server is running at http://localhost:8087")
	log.Fatal(http.ListenAndServe(":8087", r))
}
