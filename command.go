package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/nlopes/slack"
	"github.com/shomali11/slacker"
	"log"
	"time"
)

// AssignDefault assign default command
func AssignDefault(bot *slacker.Slacker) {
	bot.Init(func() {
		log.Println("Connected!")
	})

	bot.Err(func(err string) {
		log.Println(err)
	})

	bot.DefaultCommand(func(request slacker.Request, response slacker.ResponseWriter) {
		var attachments []slack.Attachment
		attachments = append(attachments, slack.Attachment{
			Title:      "そんなコマンドないです",
			ImageURL:   "https://www.irasutoya.com/2018/09/blog-post_143.html",
		})
		response.Reply("Say what?", slacker.WithAttachments(attachments))
	})

	bot.DefaultEvent(func(event interface{}) {
		//fmt.Println(event)
	})
}

// AssignPing assign ping command
func AssignPing(bot *slacker.Slacker) {
	definition := &slacker.CommandDefinition{
		Description: "Ping!",
		Example:     "ping",
		Handler: func(request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("pong")
		},
	}
	bot.Command("ping", definition)
}

// AssignEcho assign echo command
func AssignEcho(bot *slacker.Slacker) {
	definition := &slacker.CommandDefinition{
		Description: "Echo a word!",
		Example:     "echo hello",
		Handler: func(request slacker.Request, response slacker.ResponseWriter) {
			word := request.Param("word")

			var attachments []slack.Attachment
			attachments = append(attachments, slack.Attachment{
				Color:      "#2eb886",
				AuthorName: "いらすとや",
				AuthorLink: "https://www.irasutoya.com",
				AuthorIcon: "https://1.bp.blogspot.com/-s7wD--x4LBo/WUJZO318J0I/AAAAAAABE1k/cLyYpUhHxzou8EfHWbcd02LpnTfHU006gCLcBGAs/s1600/logo_sml.png",
				Title:      "強そうなゴリラのイラスト",
				Text:       "ゴリラは、霊長目ヒト科ゴリラ属（Gorilla）に分類される構成種の総称。",
				ImageURL:   "https://2.bp.blogspot.com/-ruMSXp-w-qk/XDXbUFVC3FI/AAAAAAABQ-8/QRyKKr--u9E1-Rvy2SQqt0QPWeq1ME6wgCLcBGAs/s180-c/animal_gorilla.png",
			})

			response.Reply(word, slacker.WithAttachments(attachments))
		},
	}
	bot.Command("echo <word>", definition)
}

// AssignRepeat assign repeat command
func AssignRepeat(bot *slacker.Slacker) {
	definition := &slacker.CommandDefinition{
		Description: "Repeat a word a number of times!",
		Example:     "repeat hello 10",
		Handler: func(request slacker.Request, response slacker.ResponseWriter) {
			word := request.StringParam("word", "Hello!")
			number := request.IntegerParam("number", 1)
			for i := 0; i < number; i++ {
				response.Reply(word)
			}
		},
	}
	bot.Command("repeat <word> <number>", definition)
}

// AssignTest assign test command
func AssignTest(bot *slacker.Slacker) {
	definition := &slacker.CommandDefinition{
		Description: "Tests errors",
		Handler: func(request slacker.Request, response slacker.ResponseWriter) {
			response.ReportError(errors.New("oops"))
		},
	}
	bot.Command("test", definition)
}

// AssignTime assign time command
func AssignTime(bot *slacker.Slacker) {
	definition := &slacker.CommandDefinition{
		Description: "Server time!",
		Handler: func(request slacker.Request, response slacker.ResponseWriter) {
			response.Typing()
			time.Sleep(time.Second)
			response.Reply(time.Now().Format(time.RFC1123))
		},
	}
	bot.Command("time", definition)
}

// AssignUpload assign upload command
func AssignUpload(bot *slacker.Slacker) {
	definition := &slacker.CommandDefinition{
		Description: "Upload a word!",
		Handler: func(request slacker.Request, response slacker.ResponseWriter) {
			word := request.Param("word")
			filename := request.Param("filename")
			channel := request.Event().Channel

			rtm := response.RTM()
			rtm.SendMessage(rtm.NewOutgoingMessage("Uploading file ...", channel))

			client := response.Client()
			_, err := client.UploadFile(slack.FileUploadParameters{Content: word, Filename: filename, Channels: []string{channel}})

			if err != nil {
				response.ReportError(fmt.Errorf("failed to upload: %v", err))
			}
		},
	}
	bot.Command("upload <word> <filename>", definition)
}

// AssignProcess assign process command
func AssignProcess(bot *slacker.Slacker) {
	definition := &slacker.CommandDefinition{
		Description: "Process!",
		Handler: func(request slacker.Request, response slacker.ResponseWriter) {
			timedContext, cancel := context.WithTimeout(request.Context(), time.Second)
			defer cancel()

			select {
			case <-timedContext.Done():
				response.ReportError(errors.New("timed out"))
			case <-time.After(time.Minute):
				response.Reply("Processing done!")
			}
		},
	}
	bot.Command("process", definition)
}
