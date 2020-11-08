FROM golang
RUN go get github.com/gin-gonic/gin
WORKDIR /go/src/gin-api
