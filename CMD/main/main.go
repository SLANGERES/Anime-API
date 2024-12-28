package main

import (

	"log"
	"net/http"

	"github.com/SLANGERES/Anime-API/PKG/routes"
	"github.com/gorilla/mux"
	
)

func main(){
	r:=mux.NewRouter()
	routes.RegisterAnimeRoutes(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe("localhost:9044",r))
}