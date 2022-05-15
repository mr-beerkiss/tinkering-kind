FROM golang:1.18 AS build
WORKDIR /go/src/app/
ADD . /go/src/app/

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOARCH=arm64 GOOS=linux go build -o /go/bin/server

FROM gcr.io/distroless/static
COPY --from=build /go/bin/server /
CMD ["/server"]