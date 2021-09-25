package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Solution(str string) (int, int, int, int) {
	s := strings.Split(str, "\n")

	var fileLastName []string

	for _, data := range s {
		Data := strings.Split(data, ".")
		fileLastName = append(fileLastName, Data[len(Data)-1])
	}
	var otherSize int
	var totalMusicSize int
	var totalImageSize int
	var totalMoviesSize int
	for _, data := range fileLastName {
		Data := strings.Split(data, " ")
		DataSize := strings.Split(Data[1], "b")
		Size, err := strconv.Atoi(DataSize[0])
		if err != nil {
			log.Fatal(err)
		}
		if (Data[0] == "mp3") || (Data[0] == "flac") || (Data[0] == "aac") {
			totalMusicSize += Size
		} else {
			if (Data[0] == "jpg") || (Data[0] == "gif") || (Data[0] == "bmp") {
				totalImageSize += Size
			} else {
				if (Data[0] == "mp4") || (Data[0] == "mkv") {
					totalMoviesSize += Size
				} else {
					otherSize += Size
				}
			}
		}

	}
	return totalImageSize, totalMusicSize, totalMoviesSize, otherSize
}
func main() {
	var str = "music.song.mp3 11b\nmusic.flac 1000\nnote3.txt 5b\nvideo.mp4 200b\ngame.exe 100b\nmov!l.mkv 10000b\nimage.jpg 10b"

	imageSize, musicSize, moviesSize, otherSize := Solution(str)

	fmt.Printf("Music: %db\n", musicSize)
	fmt.Printf("Images: %db\n", imageSize)
	fmt.Printf("Movice: %db\n", moviesSize)
	fmt.Printf("Others: %db\n", otherSize)
}
