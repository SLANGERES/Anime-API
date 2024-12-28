package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/SLANGERES/Anime-API/PKG/utils"
	"github.com/SLANGERES/Anime-API/PKG/models"

)
var newAnime models.Anime

func GetAnime(w http.ResponseWriter,r *http.Request){

	newAnime:=models.GetAllAnime()

	res,_:=json.Marshal(newAnime)

	w.Header().Set("Content-Type","pkglication/json")

	w.WriteHeader(http.StatusOK)

	w.Write(res)
}
func AddNewAnime(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","pkglication/json")
	
	
	createanime:=models.Anime{}

	utils.ParseBody(r,createanime)

	b:=createanime.CreateAnime()

	res,_:=json.Marshal(b)

	w.WriteHeader(http.StatusOK)

	w.Write(res)
}

func GetAnimeById(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","pkglication/json")
	vars:=mux.Vars(r)
	animeID:=vars["animeId"]
	ID,err:=strconv.ParseInt(animeID,0,0)
	if err!=nil{
		fmt.Println("error while parsing")
	}
	animeDetail, _:=models.GetAnimeByID(ID)
	res, _:=json.Marshal(animeDetail)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}	

func UpdateAnime(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","pkglication/json")
	var UpdateAnime=&models.Anime{}
	utils.ParseBody(r,UpdateAnime)

	param:=mux.Vars(r)
	animeId:=param["animeId"]
	ID,err:=strconv.ParseInt(animeId,0,0)
	if err!=nil{
		fmt.Println("error while parsing")
	}
	animedetail,db:=models.GetAnimeByID(ID)
	if UpdateAnime.Name !=""{
		animedetail.Name=UpdateAnime.Name
	}
	if UpdateAnime.Writer!=""{
		animedetail.Writer=UpdateAnime.Name
	}
	if UpdateAnime.AnimationCompany!=""{
		animedetail.AnimationCompany=UpdateAnime.AnimationCompany
	}
	db.Save(&animedetail)

	res,_:=json.Marshal(animedetail)

	w.WriteHeader(http.StatusOK)

	w.Write(res)


}

func DeleteAnime(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","pkglication/json")
	param:=mux.Vars(r)
	animeID:=param["animeID"]
	ID,err:=strconv.ParseInt(animeID,0,0)
	if err!=nil{
		fmt.Println("Error while parsing")
	}
	animeDetail:=models.DeleteAnime(ID)
	res,_:=json.Marshal(animeDetail)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}