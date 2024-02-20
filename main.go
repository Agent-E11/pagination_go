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

    router.GET("/list-page/:num", ListPage)
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

func ListPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    tmpl := template.Must(template.ParseFiles("list.tmpl.html"))

    pageNum, err := strconv.Atoi(ps.ByName("num"))
    if err != nil {
        log.Print("Invalid number")
        w.WriteHeader(404)
    }
    log.Printf("Generating page with number: %d", pageNum)

    tmpl.Execute(w, struct{
        Items []string
        Page int
        PrevPage int
        NextPage int
    }{
        Items: dbc.GetPage(dbc.DB, pageNum, 5),
        Page: pageNum,
        PrevPage: pageNum - 1,
        NextPage: pageNum + 1,
    })
}
