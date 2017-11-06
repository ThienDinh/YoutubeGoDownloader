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
		r.ParseForm()
		fmt.Println("Request was:")
		videoId := (*r).Form["id"][0]
		link := Download(videoId)
		fmt.Println(videoId)
		fmt.Fprintf(w, link.String())
	})

	http.ListenAndServe(":80", nil)

	
}