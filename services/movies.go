package services

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/simple_CRUD_API_With_Golang/Type"
	"math/rand"
	"net/http"
	"strconv"
)

var movies []Type.Movie

func init() {
	movie1 := Type.Movie{
		ID:    "1000001",
		Isbn:  "124782748",
		Title: "小王子",
		Director: &Type.Director{
			FirstName: "德·圣·埃克苏佩里",
			LastName:  "安托万",
		},
	}
	movie2 := Type.Movie{
		ID:    "1000002",
		Isbn:  "248294824",
		Title: "活着",
		Director: &Type.Director{
			FirstName: "余",
			LastName:  "华",
		},
	}
	movies = append(movies, movie1, movie2)
}
func GetMoives(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func GetMoiveById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	for _, m := range movies {
		if m.ID == id {
			json.NewEncoder(w).Encode(m)
			return
		}
	}
	json.NewEncoder(w).Encode("不存在")
}

func UpdataMoiveByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	var movie Type.Movie
	json.NewDecoder(r.Body).Decode(&movie)
	if !movie.IsValid() {
		json.NewEncoder(w).Encode("更新失败：输入的字段不完整")
		return
	}
	for index, m := range movies {
		if m.ID == id {
			movies[index] = movie
			json.NewEncoder(w).Encode("更新成功")
			return
		}
	}
	json.NewEncoder(w).Encode("更新失败 id不存在")
}

func DeleteMovieByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	for index, m := range movies {
		if m.ID == id {
			movies = append(movies[:index], movies[index+1:]...)
			json.NewEncoder(w).Encode("删除成功")
			return
		}
	}
	json.NewEncoder(w).Encode("删除失败：id不存在")
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Type.Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		json.NewEncoder(w).Encode("添加失败: 输入格式错误")
		return
	}
	if !movie.IsValid() {
		json.NewEncoder(w).Encode("添加失败: 输入字段不全")
		return
	}
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode("添加成功")
}
