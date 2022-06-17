package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/minio/minio-go/v6"

	"myapp/model"
	"myapp/utils"
)

var useSSL = false

func UploadFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	file, handler, err := r.FormFile("File")
	if err != nil {
		fmt.Println(err)
		message := &model.ResponseJson{
			Status: false,
			Data:   []model.Attachment{},
			Error:  "Error Retrieving the File!",
		}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(message)
		return
	}

	extensions := utils.FillExtension()
	arr := strings.Split(handler.Filename, ".")
	extension := arr[len(arr)-1]

	key, ok := utils.Mapkey(extensions, extension)
	if !ok {
		message := &model.ResponseJson{
			Status: false,
			Data:   []model.Attachment{},
			Error:  "Extension is Not Supported!",
		}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(message)
		return
	}

	remotePath := key
	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header.Get("Content-Type"))

	//get the file extension
	payload := handler.Header["Content-Disposition"][0]
	payloadElement := strings.Split(payload, ";")
	fileName := payloadElement[2]
	temp := strings.Split(fileName, ".")
	extensionWithTick := temp[len(temp)-1]
	UploadedExtension := strings.Split(extensionWithTick, "\"")[0]

	//get the file name
	FullFileName := strings.Split(fileName, "\"")[1]
	FileName := strings.Replace(strings.Split(FullFileName, ".")[0], " ", "", -1)
	fmt.Println(FileName)

	fullPath := UploadLocal("storage/", FileName, UploadedExtension, file)
	width, height := 0, 0

	fmt.Println(fullPath)

	if key == "photo" {
		width, height = utils.FPB(utils.GetImageDimension(fullPath))
		ResizeImage(fullPath, UploadedExtension)
	}
	link, remoteFilePaath := UploadDO(fullPath, remotePath, handler.Header.Get("Content-Type"), r.FormValue("BucketName"))

	w.Header().Set("Content-Disposition", "attachment; filename="+handler.Filename)
	json.NewEncoder(w).Encode(model.ResponseJson{
		Status: true,
		Data: []model.Attachment{
			{
				ServeLink: link,
				Path:      remoteFilePaath,
				Width:     width,
				Height:    height,
			},
		},
		Error: "",
	})
}

func UploadDO(fullPath string, remotePath string, fileType string, bucketName string) (string, string) {
	var endPoint string = os.Getenv("MINIO_SERVER_ENDPOINT")
	var accessKeyID = os.Getenv("MINIO_ACCESS_KEY")
	var secretAccessKey = os.Getenv("MINIO_SECRET_KEY")
	fmt.Println(fullPath)
	file, _ := utils.PathToFile(fullPath, fileType)

	minioClient, err := minio.New(endPoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		fmt.Println(err)
	}

	if bucketName == "" {
		bucketName = "portfolio"
	}

	// bucketNames := os.Getenv("BUCKET_NAME")
	objectName := fullPath
	contentType := fmt.Sprintf("application/%s", file[0].ContentType)
	retainDate := time.Now().AddDate(0, 1, 0)
	remoteFilePath := remotePath + "/" + file[0].Filename
	fmt.Println(minioClient.ListBuckets())
	fmt.Println(bucketName, remoteFilePath, objectName, minioClient, file[0].Filename)
	_, err = minioClient.FPutObject(bucketName, remoteFilePath, objectName, minio.PutObjectOptions{
		ContentType: contentType,
		UserTags: map[string]string{
			"expired": retainDate.Format("2006-01-02"),
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	// e := os.Remove(fullPath)
	// if e != nil {
	// 	log.Fatal(e)

	// }
	return ServeImage(remoteFilePath, bucketName), remoteFilePath
}

// if remotePath != "photo" {
// 	var responses []model.Attachment
// 	if remotePath == "word" {
// 		responses = ServeImageDo("assets/mswordicon.png")
// 	}else if remotePath == "pdf" {
// 		responses = ServeImageDo("assets/pdficon.png")
// 	}else if remotePath == "excel" {
// 		responses = ServeImageDo("assets/msexcelicon.png")
// 	}else if remotePath == "powerpoint" {
// 		responses = ServeImageDo("assets/ppticon.png")
// 	}else if remotePath == "video"{
// 		responses = ServeImageDo("assets/mp4icon.png")
// 	}else if remotePath == "audio" {
// 		responses = ServeImageDo("assets/mp3icon.png")
// 	}else if remotePath == "compressed" {
// 		responses = ServeImageDo("assets/zip.png")
// 	}
// 	responses[0].Downloadable = "https://static."+os.Getenv("BUCKET_NAME")+"."+os.Getenv("DOMAIN_EXTENSION")+"/"+remoteFilePath
// 	responses[0].DownloadablePath = remoteFilePath

// 	return responses
// }
