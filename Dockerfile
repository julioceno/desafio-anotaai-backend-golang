FROM golang:1.22.5-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -o /bin/desafio-anotaai-backend-golang ./cmd/desafio-anotaai-backend-golang/desafio-anotaai-backend-golang.go

FROM scratch

WORKDIR /app

COPY --from=builder /bin/desafio-anotaai-backend-golang .

EXPOSE 8080

ENTRYPOINT [ "./desafio-anotaai-backend-golang" ]