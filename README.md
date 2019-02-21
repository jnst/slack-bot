# slack-bot

🤖 Slack api bot by golang

## Golang

```bash
$ brew install go
$ go version
go version go1.11.5 darwin/amd64
```

### Modules (vgo)

```bash
$ env | grep GO111MODULE
GO111MODULE=on
```

>In my case, I want to manage all the source code in GOPATH, so this is necessary.


### Dependencies

* [shomali11/slacker](https://github.com/shomali11/slacker)
* [nlopes/slack](https://github.com/nlopes/slack)

## Run

```bash
$ SLACK_TOKEN=xxxxx PORT=3000 go run main.go
```

## Docker

```bash
$ docker build -t jnst/slack-bot .
$ docker run -it --rm -e SLACK_TOKEN=xxxxx -e PORT=3000 jnst/slack-bot
```

## Deploy on Heroku

```bash
$ heroku login
$ heroku create
$ heroku git:remote -a xxx-xxx-00001
$ heroku config:set SLACK_TOKEN=xxxxx
$ heroku config:set PORT=3000
$ heroku container:login
$ heroku container:push web
$ heroku container:release web
```
