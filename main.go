package main

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	dbc "github.com/agent-e11/pagination_go/dbcontrol"
	"github.com/julienschmidt/httprouter"
)

func main() {
    router := httprouter.New()

    router.GET("/list/:test", List)
    router.GET("/home/:page", Index)

    log.Print("Running http server...")
    http.ListenAndServe(":8000", router)
}

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    tmpl := template.Must(template.ParseFiles("index.tmpl.html"))

    pageNum, err := strconv.Atoi(ps.ByName("page"))
    if err != nil {
        w.WriteHeader(404)
    }
    
    data := map[string]any{
        "Items": dbc.GetPage(dbc.DB, pageNum, 5),
        "Page": pageNum,
        "NextPage": pageNum + 1,
        "PrevPage": pageNum - 1,
    }

    tmpl.Execute(w, data)
}

func List(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    tmpl := template.Must(template.ParseFiles("list.tmpl.html"))

    pageNum, err := strconv.Atoi(ps.ByName("page"))
    if err != nil {
        w.WriteHeader(404)
    }

    tmpl.Execute(w, dbc.GetPage(dbc.DB, pageNum, 5))
}
