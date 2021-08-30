package service

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"

	"github.com/nfnt/resize"
)

func UploadLocal(path string, FileName string, extension string, file multipart.File) string {

	tempFile, err := ioutil.TempFile("./"+path, FileName+"-*."+extension)
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	tempFile.Write(fileBytes)

	return tempFile.Name()
}

func ResizeImage(path string,extension string){
	resizeFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	var img image.Image

	if extension == "png" || extension == "PNG" {
		img ,err = png.Decode(resizeFile)
	}else {
		img, err = jpeg.Decode(resizeFile)		
	}

	if err != nil {
		log.Fatal(err)
	}

	resizeFile.Close()

	m := resize.Resize(1250, 0, img, resize.Lanczos3)

	out, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	jpeg.Encode(out, m, nil)
}	
