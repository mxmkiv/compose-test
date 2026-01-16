FROM golang:1.25-alpine

WORKDIR /docker-app
COPY go.mod go.sum ./
RUN go mod download

# COPY . .
# RUN go build -o main ./main.go
# CMD ["./main"]

CMD [ "go", "run", "." ]