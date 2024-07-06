package handler

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
	"web/entity"
)
func ProductPage(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNumber, error := strconv.Atoi(id)
	if error != nil || idNumber < 1 {
		http.NotFound(w, r)
		return
	}
	tmpl , err :=template.ParseFiles(path.Join("views" , "product.html") , path.Join("views" , "layout.html"))
	if err != nil{
			 log.Println("Error => " , err)
			 http.Error(w ,  "Terjadi Kesalahan" , http.StatusInternalServerError)
			 return
	}
	data := map[string]interface{}{
       "content" :idNumber , 
	}
	err = tmpl.Execute(w , data)
	if err != nil{
		log.Println("Error => " , err)
		http.Error(w,  "Terjadi Kesalahan" , http.StatusInternalServerError)
		return
}

}
func HomeHandle(k http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(k, r)
		return
	}

	// k.Write([]byte("Tampila	n Utama"))
	 tmpl , err :=template.ParseFiles(path.Join("views" , "index.html") , path.Join("views" , "layout.html"))
	 if err != nil{
	    	log.Println("Error => " , err)
		    http.Error(k ,  "Terjadi Kesalahan" , http.StatusInternalServerError)
				return
	 }
	//  data := map[string]interface{}{ // tipe value interface jika ingin mengakses keseluruhan tipe data 
	// 	"title":"Learning Golang WEB",
	// 	"content":"Programming",
	//  }
	//data["owner"]="Frans"


	// passing slice of struct
	// data := entity.Product{Id: 1 , Nama : "Cookies Pineaple " , Harga: 35000, Stock: 76}

	// cara ke dua slice 
	data2 := []entity.Product{
		{Id: 1 , Nama : "Cookies Pineaple " , Harga: 35000, Stock: 7},
		{Id: 2 , Nama : "Cookies Cheeze " , Harga: 90000, Stock: 2},
		{Id: 3 , Nama : "Cookies Peanut " , Harga: 40000, Stock: 12},
	}
	 // memasukkan data sebagai parameter kedua
   err = tmpl.Execute(k  , data2 ) // memasukkan data pada pada parameter kedua 
	 if err != nil{
		log.Println("Error => " , err)
		http.Error(k ,  "Terjadi Kesalahan" , http.StatusInternalServerError)
		return
}
}
func Form(w http.ResponseWriter , r *http.Request){
	  if r.Method == "GET" {
			tmpl , err := template.ParseFiles(path.Join("views" , "form.html") , path.Join("views" , "layout.html"))
			if err != nil{
           log.Println(err)
					 http.Error(w , "ADA KESALAHAN MOHON TENANG" , http.StatusInternalServerError)
					 return
			}
			err = tmpl.Execute(w , nil)
					if err != nil{
					log.Println(err)
					http.Error(w , "ADA KESALAHAN HARAP TENANG" , http.StatusInternalServerError)
					return
					}
			return
		}
		http.Error(w , "ADA KESALAHAN HARAP TENANG " , http.StatusInternalServerError)
}
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World saya sedang belajar Golang web "))
}
func MarioHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Halo saya Mario"))
}
func Postget(w http.ResponseWriter , r *http.Request){
    method := r.Method // GET ATAU POST

		switch method{
		case "GET":
			w.Write([]byte("INI ADALAH GET"))
			break
		case "POST":
			w.Write([]byte("INI ADALAH POST"))
    default:
			http.Error(w , "INI ADALAH EROR" , http.StatusBadRequest)		
		}
	   

}
func Process(w http.ResponseWriter , r *http.Request){
	if r.Method == "POST"{
		  err:=r.ParseForm()
			if err != nil{
				log.Println(err)
				http.Error(w , "ADA KESALAHAN HARAP TENANG" , http.StatusInternalServerError)
				return
			}
		 name := r.Form.Get("name")
		 pesan := r.Form.Get("pesan")

		 data := map[string]interface{}{
			"name" : name ,
			"message" :pesan ,
		 }
      tmpl , err := template.ParseFiles(path.Join("views" , "result.html") , path.Join("views" , "layout.html")) // join ke template
			if err != nil{
				log.Println(err)
				http.Error(w , "ADA KESALAHAN HARAP TENANG" , http.StatusInternalServerError)
				return
			}
			err = tmpl.Execute(w , data)
			if err != nil{
				log.Println(err)
				http.Error(w , "ADA KESALAHAN HARAP TENANG" , http.StatusInternalServerError)
				return
			}
		 
		 return // akan selesai disini
	}
	http.Error(w , "INI ADALAH EROR" , http.StatusBadRequest)
}
