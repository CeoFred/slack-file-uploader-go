package main

import (
	"fmt"
	"os"
	"github.com/slack-go/slack"
)
func main(){
	os.Setenv("SLACK_BOT_TOKEN","")
	os.Setenv("CHANNEL_ID", "C010NHVPB43")
	 api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	 channelAdr := []string{os.Getenv("CHANNEL_ID")}
	 fileArr := []string{"hello.js"}

  for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelAdr,
			File: fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Name: %s\n, URL: %s \n", string(file.Name), file.URL)
	}
}