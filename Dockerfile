FROM golang:1.19-alpine
WORKDIR /app


COPY go.mod /

COPY *.go ./
RUN  go build -o /go_api
EXPOSE 8909
CMD [ "/go_api" ]
