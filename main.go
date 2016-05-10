package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Persona struct {
	Email    string
	Password string
}

func test(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("cual es el req. body ", req.Body)
	//req.Body = `{"Name":"Adam","Age":36,"Job":"CEO"}`

	decoder := json.NewDecoder(req.Body)
	var t Persona
	err := decoder.Decode(&t)
	if err != nil {
		fmt.Println("err", err)
	}
	log.Println(t.Email)
	log.Println(t.Password)
}

func main() {
	http.HandleFunc("/test", test)
	log.Fatal(http.ListenAndServe(":8082", nil))
}
