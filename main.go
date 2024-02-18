package main

import (
	"fmt"

	dbc "github.com/agent-e11/pagination_go/dbcontrol"
)

func main() {
    var quantity int = 5
    for i := 0; i < 10; i++ {
        fmt.Printf("Page %d: %v\n", i, dbc.GetPage(dbc.DB, i, quantity))
    }
}
