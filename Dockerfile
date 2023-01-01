FROM golang:1.19 as backend-build
WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/static-debian11
COPY --from=backend-build /go/bin/app /
CMD ["/app"]
