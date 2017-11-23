package main 

import ("fmt"
		"net/http"
		"html/template")

type Data struct{
	Name string
	Price float32
	MarkCap int
}

func indexHandler(w http.ResponseWriter, r *http.Request){

	fmt.Fprintf(w,"<h1>New tag</h1>")
}
func indexInfo(w http.ResponseWriter, r *http.Request){
	data := Data{Name: "Expanse",Price: 1.83, MarkCap: 17000000}
	t, _ := template.ParseFiles("index.html")
	t.Execute(w,data)
}


func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/info", indexInfo)
	http.ListenAndServe(":8001", nil)
}