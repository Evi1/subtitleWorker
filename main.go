package main

import (
	"path/filepath"
)

var videoSuffix []string
var subSuffix []string

func init() {
	videoSuffix = make([]string, 0)
	subSuffix = make([]string, 0)
}

func main() {
	videoSuffix = append(videoSuffix, "mp4")
	videoSuffix = append(videoSuffix, "mkv")
	videoSuffix = append(videoSuffix, "flv")
	videoSuffix = append(videoSuffix, "avi")
	subSuffix = append(subSuffix, "ass")
	handleFolder(filepath.ToSlash(""))
}
