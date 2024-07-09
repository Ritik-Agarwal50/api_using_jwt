FROM golang:1.19
RUN mkdir /app
ADD . /app
WORKDIR /app
COPY go.mod go.sum ./
COPY . .
RUN go mod download
RUN go build -o main ./cmd/main.go || { echo 'Go build failed' ; exit 1; }
EXPOSE 6543
CMD ["/app/main"]

