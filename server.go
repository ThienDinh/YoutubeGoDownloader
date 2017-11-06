package main

import (
	"fmt"
	"github.com/rylio/ytdl"
	"net/url"
	// "os"
)

func Download(videoId string)  (link string){	
	videoInfor, err :=ytdl.GetVideoInfoFromID(videoId)
	if(err != nil) {
		fmt.Println("An error occurred!")
	} else {
		fmt.Println((*videoInfor).Title)
		fmt.Println((*videoInfor).Description)
		formatList := (*videoInfor).Formats.Filter(ytdl.FormatVideoEncodingKey,
			[](interface{}){""}).Filter(ytdl.FormatAudioEncodingKey, [](interface{}){"aac", "mp3"})
		for _, format := range formatList {
			// fileName := fmt.Sprintf("downloaded video/%s_%d.%s", "video", format.Itag, format.Extension)
			// os.Remove(fileName)
			// file, err := os.Create(fileName)
			if err != nil {
				fmt.Println("Error!")
			} else {
				// videoInfor.Download(format, file)
				downloadUrl, _ := videoInfor.GetDownloadURL(format)
				link, _ = url.QueryUnescape(downloadUrl.String())
				fmt.Printf("%v\n", link)				
			}
		}
	}
	return
}