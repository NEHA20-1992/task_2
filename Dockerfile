FROM golang:1.19-alpine
WORKDIR /app


COPY go.mod ./
COPY go.sum ./

COPY *.go ./
RUN  go build -o /go_api1
EXPOSE 8909
CMD [ "app/go_api1" ]
