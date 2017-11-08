package main

import (
	"fmt"
	"github.com/rylio/ytdl"
	"net/url"
	"time"
	// "os"
)

type YoutubeVideo struct {
	Id string
	Title string
	Duration time.Duration
	Extension string
	VideoLink string
	ThumbnailLink string
}

func (yv YoutubeVideo) String() string{
	return fmt.Sprintf("%s\n%s\n%d\n%s\n%s\n%s", yv.Id, yv.Title, yv.Duration, yv.Extension, yv.VideoLink, yv.ThumbnailLink)
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
				// thumbnail := url.QueryUnescape("")
				thumbnail := videoInfor.GetThumbnailURL(ytdl.ThumbnailQualityMaxRes).String()
				youtubeVideo = YoutubeVideo{Id: videoId, Title:(*videoInfor).Title, Extension: format.Extension, Duration: (*videoInfor).Duration,
					VideoLink: link, ThumbnailLink: thumbnail}
				fmt.Printf("%v\n", youtubeVideo)				
			// }
		}
	}
	return
}