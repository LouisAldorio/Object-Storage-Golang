package service

import (
	"encoding/json"
	"fmt"
	"myapp/model"
	"net/http"
	"os"
)

func ServeImage(path string) string {
	return fmt.Sprintf("%s/%s",os.Getenv("SERVE_ENDPOINT"), path)
}

func Serve(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var paths model.RequestJson
	var response []model.Attachment

	err := json.NewDecoder(r.Body).Decode(&paths)
	if err != nil {
		fmt.Println(err)
	}

	for _,v := range paths.Paths{
		response = append(response, model.Attachment{
			ServeLink: ServeImage(v),
		})
	}

	json.NewEncoder(w).Encode(model.ResponseJson{
		Status: true,
		Data: response,
		Error: "",
	})
}
