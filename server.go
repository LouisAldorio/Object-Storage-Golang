package main

import (
	"fmt"
	"log"
	"myapp/service"
	"net/http"
	"os"
	"os/user"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	dockerVolumeDir := usr.HomeDir + "/data-minio/" + os.Getenv("BUCKET_NAME")
	fmt.Println(dockerVolumeDir)

	router := mux.NewRouter()
	router.HandleFunc("/upload", service.UploadFile).Methods("POST")
	router.HandleFunc("/serve",service.Serve)
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(dockerVolumeDir))))

	log.Println("Listening on 8081")
	err = http.ListenAndServe(":8081", router)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
