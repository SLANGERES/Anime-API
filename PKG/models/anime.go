package models

import (
	"github.com/SLANGERES/Anime-API/PKG/config"

	"gorm.io/gorm"

)

var db *gorm.DB

type Anime struct {
	gorm.Model
	Name             string `json:"name"`
	Writer           string `json:"writer"`
	AnimationCompany string `json:"animationCompany"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	
	db.AutoMigrate(&Anime{})
}

func (b *Anime) CreateAnime() *Anime {
	db.Create(&b)
	return b
}

func GetAllAnime() []Anime {
	var animes []Anime
	db.Find(&animes)
	return animes
}
func GetAnimeByID(Id int64) (*Anime, *gorm.DB){
	var getAnime Anime
	db:=db.Where("ID=?",Id).Find(&getAnime)
	return &getAnime, db
}
func DeleteAnime(Id int64) Anime{
	var anime Anime
	db.Where("ID=?",Id).Delete(anime)
	return anime
}