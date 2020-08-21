FROM golang:latest 

WORKDIR /server

COPY config /server/config
COPY util /server/util 
COPY apps /server/apps
COPY dbi /server/dbi
COPY router /server/router
COPY go.mod /server
COPY go.sum /server
COPY main.go /server

RUN go mod download
RUN go build -o main /server/main.go

EXPOSE 8080

CMD [ "./main" ]
