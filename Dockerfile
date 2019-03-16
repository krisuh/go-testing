FROM golang:latest as compiler
WORKDIR /go/src/github.com/krisuh/go-testing
ADD . /go/src/github.com/krisuh/go-testing/
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure -v
RUN go GOARCH=arm GOOS=linux build -o program

FROM arm32v6/alpine
WORKDIR /app
COPY --from=compiler /go/src/github.com/krisuh/go-testing/program /app
CMD ./program