FROM golang:latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
CMD [ "go build -o ./go-bank","./go-bank" ]
EXPOSE 8080