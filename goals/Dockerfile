FROM golang:alpine
ENV CI=true

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o goals .
EXPOSE 3000

CMD ["./auth"]