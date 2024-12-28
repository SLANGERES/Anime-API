package routes


//in go lang it work on absolute path in it 
import (


	"github.com/SLANGERES/Anime-API/PKG/controllers"
	"github.com/gorilla/mux"
)

//Main Router Function
var RegisterAnimeRoutes=func(router *mux.Router){

	//THERE IS ROUTER FUNCTION THAT HANDLE VIA FUNCTION THAT ARE IN CONTROLLER DIR
	router.HandleFunc("/anime/",controllers.AddNewAnime).Methods("POST")
	
	router.HandleFunc("/anime/",controllers.GetAnime).Methods("GET")
	
	router.HandleFunc("/anime/{animeId}",controllers.GetAnimeById).Methods("GET")

	router.HandleFunc("/anime/{animeId}",controllers.UpdateAnime).Methods("PUT")

	router.HandleFunc("/anime/{animeId}",controllers.DeleteAnime).Methods("DELETE")

}
