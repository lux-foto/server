package main

import (
	"log"
	"net/http"
	"github.com/lux-foto/gphoto-camera"
	"github.com/lux-foto/movie-server"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/camera/{type}", gphoto.CameraHandler)
	r.HandleFunc("/video/shot/{shot}/{movie}/", movie_server.VideoShotServer)
	r.HandleFunc("/video/{name}/", movie_server.VideoMovieServer)
	log.Println("Listening on 3000")
	log.Fatal(http.ListenAndServe(":3000", r))

}

