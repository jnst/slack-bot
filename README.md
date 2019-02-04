# slack-bot

🤖 Slack api bot using golang

## Golang

```bash
$ brew install go
$ go version
go version go1.11.5 darwin/amd64
```

## Modules (vgo)

```bash
$ env | grep GO111MODULE
GO111MODULE=on
```

## Run

```bash
$ SLACK_TOKEN=xxxxxx go run main.go
```

## Packages

* [nlopes/slack](https://github.com/nlopes/slack)

## Docker

```bash
$ docker build -t jnst/slack-bot .
$ docker run -it --rm jnst-slack-bot
```
