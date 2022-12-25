FROM golang:1.19-alpine3.16

ADD . /go/src/app

WORKDIR /go/src/app/app

#ENV PORT 8080
#EXPOSE 8080

#RUN go build -mod=vendor -o main .
#CMD ["/go/src/app/main"]

#RUN go mod init
#RUN go get github.com/gin-gonic/gin
#RUN go get gorm.io/driver/mysql gorm.io/gorm


