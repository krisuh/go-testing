FROM golang:latest as compiler
WORKDIR /app
COPY main.go /app
RUN go get github.com/stianeikeland/go-rpio
RUN GOARM=7 GOARCH=arm GOOS=linux go build -o blinker

FROM arm32v6/alpine
WORKDIR /app
COPY --from=compiler /app/blinker /app
CMD ./blinker