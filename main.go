package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/nlopes/slack"
	"github.com/nlopes/slack/slackevents"
	"github.com/shomali11/slacker"
	"log"
	"net/http"
	"os"
)

var token = os.Getenv("SLACK_TOKEN")
var port = os.Getenv("PORT")
var api = slack.New(token)

func main() {
	log.Printf("ENV: token=%#v, port=%v\n", token, port)

	bot := slacker.NewClient(token)
	AssignDefault(bot)
	AssignPing(bot)
	AssignEcho(bot)
	AssignRepeat(bot)
	AssignTest(bot)
	AssignTime(bot)
	AssignUpload(bot)
	AssignProcess(bot)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//go startEventsApi()
	//
	//api := slack.New(token)
	//rtm := api.NewRTM()
	//go rtm.ManageConnection()
	//
	//for msg := range rtm.IncomingEvents {
	//	fmt.Printf("Event Received: %v\n", msg)
	//}
}

func startEventsApi() {
	http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		body := buf.String()
		log.Printf("%v %v %v\n", r.Method, r.RequestURI, body)

		eventsAPIEvent, e := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionVerifyToken(&slackevents.TokenComparator{VerificationToken: token}))
		if e != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		if eventsAPIEvent.Type == slackevents.URLVerification {
			var r *slackevents.ChallengeResponse
			err := json.Unmarshal([]byte(body), &r)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			w.Header().Set("Content-Type", "text")
			w.Write([]byte(r.Challenge))
		}
		if eventsAPIEvent.Type == slackevents.CallbackEvent {
			innerEvent := eventsAPIEvent.InnerEvent
			switch ev := innerEvent.Data.(type) {
			case *slackevents.AppMentionEvent:
				log.Fatal(api.PostMessage(ev.Channel, slack.MsgOptionText("Yes, hello.", false)))
			}
		}

		res, err := json.Marshal("error")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(res)
	})

	fmt.Printf("[INFO] Server listening on port %v\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
