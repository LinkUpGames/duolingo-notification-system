FROM golang:tip-alpine3.22

WORKDIR /app

# Copy Soruce files
COPY . .

RUN chmod +w go.mod
RUN go mod tidy
RUN go build server.go

CMD [ "./server" ]
