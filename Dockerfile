FROM golang:latest
WORKDIR /go/src/app
COPY main.go /go/src/app
COPY vendor /go/src/app/vendor
COPY static /go/src/app/static
COPY public/views /go/src/app/public/views

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN apk --no-cache add curl
WORKDIR /root/
COPY --from=0 go/src/app/main .
COPY --from=0 go/src/app/static static
COPY --from=0 go/src/app/public/views public/views
CMD ["./main"]
LABEL version=demo-3