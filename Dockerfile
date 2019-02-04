FROM golang:1.11-alpine3.8 as builder

RUN apk add --no-cache ca-certificates git

WORKDIR /src

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./
RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /bot .

FROM alpine:3.8 AS final
COPY --from=builder /bot /usr/local/bin/bot

EXPOSE 3000

CMD ["bot"]
