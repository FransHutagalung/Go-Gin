package main

import (
	"web/handler"  // cari dari go mod
	"log"
	"net/http"
	"time"
	"fmt"
)
func main() {
 
}
func Ticker(){
	timer := time.NewTicker(2 *time.Second)  // setiap 2 detik
	for{
		fmt.Println("wkwkwkwk")
		<- timer.C
	}
}
func koneksi(){
		// Buat Rooting
		mux := http.NewServeMux()
		about := func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("About Section"))
		} // cara ke dua
		mux.HandleFunc("/", handler.HomeHandle) // Pattern yang lain notomatis lari ke sini
		mux.HandleFunc("/about", about)
		mux.HandleFunc("/hello", handler.HelloHandler)
		mux.HandleFunc("/mario", handler.MarioHandler)
		mux.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) { // cara ke tiga
		w.Write([]byte("Kontak kami 087796162643"))
		}) // anonymous function
		mux.HandleFunc("/products", handler.ProductPage)
		mux.HandleFunc("/post-get" , handler.Postget)
		mux.HandleFunc("/form" , handler.Form)
		mux.HandleFunc("/process" , handler.Process)
	
		// BERIKUT ADALAH CARA MENGINISIASIKAN CSS KE DALAM TEMPLATE GOLANG 
		// http.Handle("/static/" , http.StripPrefix("/static/" , http.FileServer(http.Dir("assets"))))
	
		// pakai mux 
		fileSever := http.FileServer(http.Dir("assets"))
		mux.Handle("/static/"  , http.StripPrefix("/static" , fileSever))
	
	
		log.Println("Starting on port 8080")
		err := http.ListenAndServe(":8080", mux)
		log.Fatal(err)
}

