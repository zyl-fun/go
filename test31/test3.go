package test31

import (
	"html/template"
	"net/http"
)

func Test3tmpl(w http.ResponseWriter, r *http.Request) {
	t1, err := template.ParseFiles("test31/index.html", "test31/footer.html", "test31/header.html")
	if err != nil {
		panic(err)
	}
	t1.Execute(w, nil)
}

func Test3() {
	server := http.Server{
		Addr: "127.0.0.1:9999",
	}
	http.HandleFunc("/tmpl", Test3tmpl)
	server.ListenAndServe()
}
