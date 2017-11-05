package main

import (
	"fmt"
	"github.com/rylio/ytdl"
	"os"
)

func main(){
	fmt.Println("Hello Go, Long time no see!")
	if len(os.Args) <= 1 {
		fmt.Println("No Youtube ID provided!")
		return
	}
	videoInfor, err :=ytdl.GetVideoInfoFromID(os.Args[1])
	if(err != nil) {
		fmt.Println("An error occurred!")
	} else {
		fmt.Println((*videoInfor).Title)
		fmt.Println((*videoInfor).Description)
		formatList := (*videoInfor).Formats.Filter(ytdl.FormatVideoEncodingKey,
			[](interface{}){""}).Filter(ytdl.FormatAudioEncodingKey, [](interface{}){"aac", "mp3"})
		for _, format := range formatList {
			fileName := fmt.Sprintf("downloaded video/%s_%d.%s", "video", format.Itag, format.Extension)
			os.Remove(fileName)
			file, err := os.Create(fileName)
			if err != nil {
				fmt.Println("Error!")
			} else {
				videoInfor.Download(format, file)
				fmt.Printf("%v\n", format)				
			}
		}
	}
}