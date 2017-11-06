package main

import (
	"fmt"
	"github.com/rylio/ytdl"
	"net/url"
	// "os"
)

type YoutubeVideo struct {
	Id string
	Title string
	Extension string
	Link string
}

func (yv YoutubeVideo) String() string{
	return fmt.Sprintf("%s\n%s\n%s\n%s", yv.Id, yv.Title, yv.Extension, yv.Link)
}

func Download(videoId string)  (youtubeVideo YoutubeVideo){	
	videoInfor, err :=ytdl.GetVideoInfoFromID(videoId)
	if(err != nil) {
		fmt.Println("An error occurred!")
	} else {
		formatList := (*videoInfor).Formats.Filter(ytdl.FormatVideoEncodingKey,
			[](interface{}){""}).Filter(ytdl.FormatAudioEncodingKey, [](interface{}){"aac", "mp3"})
		for _, format := range formatList {
			// fileName := fmt.Sprintf("downloaded video/%s_%d.%s", (*videoInfor).ID, format.Itag, format.Extension)
			// os.Remove(fileName)
			// file, err := os.Create(fileName)
			// if err != nil {
			// 	fmt.Println("Error!")
			// } else {
				// videoInfor.Download(format, file)
				downloadUrl, _ := videoInfor.GetDownloadURL(format)
				link, _ := url.QueryUnescape(downloadUrl.String())
				youtubeVideo = YoutubeVideo{Id: videoId, Title:(*videoInfor).Title, Extension: format.Extension, Link: link}
				fmt.Printf("%v\n", youtubeVideo)				
			// }
		}
	}
	return
}