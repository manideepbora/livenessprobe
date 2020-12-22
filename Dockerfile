## Base image for the go app
FROM golang:1.14.10-alpine3.11

## create a app directory
RUN mkdir /app

## copy everything in root directory 
## into /app directory

ADD . /app

## specify that we wish to execute any further command from /app directory

WORKDIR /app

#COPY cmd/go.mod . 
#COPY cmd/go.sum .
RUN go mod download

COPY . .
## go build to compile to binary


## Build the binary
## RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main cmd
## RUN go build -o ./out/go-sample-app .
## FROM scratch 
## RUN ls
## COPY ./out/go-sample-app/. ./app
## ENTRYPOINT ["/app"]
RUN go mod tidy
RUN go build -o main . 

FROM alpine:latest
ADD . /app

WORKDIR /app

COPY --from=0 /app/main .
## start the command

CMD ["/app/main"]