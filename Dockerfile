FROM golang:1.14

WORKDIR /app

COPY . .

RUN go clean --modcache
RUN go mod download
RUN go get github.com/cosmtrek/air

CMD ["air"]