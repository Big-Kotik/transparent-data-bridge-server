FROM golang:1.20

# --mount works only externaly, need to mount from Makefile
ENV basicDir="/files/"

WORKDIR /data_bridge
COPY . .
RUN go mod download
RUN go mod tidy
RUN go build -o bridge cmd/bridge/main.go

CMD ./bridge
