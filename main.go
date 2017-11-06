package main

import (
	"fmt"
	"net/http"
)


func main(){
	// fmt.Println("Hello Go, Long time no see!")
	// if len(os.Args) <= 1 {
	// 	fmt.Println("No Youtube ID provided!")
	// 	return
	// }

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request was:")
		r.ParseForm()
		fmt.Println((*r).Form)
		videoId := ((*r).Form["id"])[0]
		link := Download(videoId)
		fmt.Fprintf(w, link.String())
	})

	http.ListenAndServe(":80", nil)

	
}