package main

import{
	"log"
	"net/http"
}

func main(){
	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Println("listening on port 3000.....")
	http.ListeningAndServe(":3000", nil)
}