package model

type ResponseJson struct {
	Status bool         `json:"status"`
	Data   []Attachment `json:"data"`
	Error  string       `json:"error"`
}

type Attachment struct {
	Path             string          `json:"path"`
	ServeLink        string          `json:"serve_link"`
	Width            int          `json:"width"`
	Height           int          `json:"height"`
	Downloadable     string          `json:"downloadable"`
	DownloadablePath string          `json:"downloadable_path"`
	Sizes            AttachmentSizes `json:"sizes"`
}
type AttachmentSizes struct {
	Small  string `json:"small"`
	Medium string `json:"medium"`
	Large  string `json:"large"`
}
