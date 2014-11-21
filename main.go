package main

import (
	"fmt"
	"github.com/benbayard/crad.io/magic"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	crads, cmcs := crad.GetCrads()

	cc := &crad.CradController{
		Crads: crads,
		Cmcs:  cmcs,
	}

	router.GET("/cmcs/:cmc", cc.Cmc)

	// fmt.Printf("%#v", cmcs[1])

	log.Fatal(http.ListenAndServe(":8080", router))

}
