FROM golang:alpine

# make app directory
RUN mkdir /app
ADD . /app
WORKDIR /app

# build the app
RUN go build -o main .

# entry command
CMD ["/app/main"]