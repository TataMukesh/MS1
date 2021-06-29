FROM golang:alpine

RUN mkdir /app

WORKDIR /app

ADD go.mod .
ADD go.sum .

RUN go mod download
ADD . .

RUN go get github.com/githubnemo/CompileDaemon

EXPOSE 8081

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main
#FROM golang:1.13

#WORKDIR /Experiments/MS1

#COPY . .

#RUN go get -d -v ./...
#RUN go install -v ./...
#RUN go get -u github.com/gin-gonic/gin
#RUN go get -u github.com/jinzhu/gorm
#RUN go get -u github.com/go-sql-driver/mysql

#RUN go get github.com/githubnemo/CompileDaemon

#ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main


#EXPOSE 8081