FROM golang:1.21-alpine

WORKDIR /app
#COPY . /app
COPY go.mod /app
COPY startup.sh /app
COPY air.toml /app

RUN go mod download
RUN apk add --no-cache bash
#RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install github.com/go-delve/delve/cmd/dlv@v1.21.2
# airのインストール（ホットリロード）
# RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/cosmtrek/air@v1.49.0

ENV PATH="/go/bin:${PATH}"

RUN chmod +x ./startup.sh

EXPOSE 8080

CMD ["/bin/bash", "./startup.sh"]
