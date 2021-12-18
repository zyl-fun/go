package test31

import (
	"html/template"
	"net/http"
)

func tmpl(w http.ResponseWriter, r *http.Request) {
	tl, err := template.ParseFiles("D:\\my_go_code\\helloworld\\test31\\test.html")
	if err != nil {
		panic(err)
	}
	s := []string{"多课网", "GO", "张云雷"}

	tl.Execute(w, s)
}

func Test2() {
	server := http.Server{
		Addr: "127.0.0.1:9999",
	}
	http.HandleFunc("/tmpl", tmpl)
	server.ListenAndServe()
}