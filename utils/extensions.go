package utils

func Mapkey(m map[string][]string, value string) (key string, ok bool) {
	for k, v := range m {
		for _, val := range v {
			if val == value {
				key = k
				ok = true
				return
			}
		}
	}
	return
}

func FillExtension() map[string][]string {
	extensions := make(map[string][]string)

	//docs
	docs := []string{"docs", "doc", "docx"}
	extensions["word"] = docs

	//pdf
	pdf := []string{"pdf"}
	extensions["pdf"] = pdf

	//excel
	excel := []string{"xlsx", "xls", "csv", "xltx", "xlsb"}
	extensions["excel"] = excel

	//powerpoint
	powerpoint := []string{"ppt", "pptx"}
	extensions["powerpoint"] = powerpoint

	//image
	image := []string{"jpeg", "jpg", "png","JPG","PNG","JPEG","jPG"}
	extensions["photo"] = image

	//video
	video := []string{"mp4", "MP4"}
	extensions["video"] = video

	//audio
	audio := []string{"mp3", "MP3","mpeg","MPEG"}
	extensions["audio"] = audio

	//compressed
	compressed := []string{"zip","ZIP"}
	extensions["compressed"] = compressed

	//gif
	gif := []string{ "GIF", "gif"}
	extensions["gif"] = gif

	return extensions
}