package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/SLANGERES/Anime-API/PKG/config"
	"github.com/SLANGERES/Anime-API/PKG/routes"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm/dialects/mysql"
)

func main(){
	r:=mux.NewRouter()
	routes.RegisterAnimeRoutes(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe("localhost:9044",r))
}