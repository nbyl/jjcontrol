FROM node:18 as frontend-build
WORKDIR /usr/src/app
COPY frontend/package.json frontend/yarn.lock ./
RUN yarn
COPY frontend/ ./
RUN yarn build

FROM golang:1.19 as backend-build
WORKDIR /go/src/app
COPY . .
COPY --from=frontend-build /usr/src/app/dist web/dist
RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/static-debian11
COPY --from=backend-build /go/bin/app /
CMD ["/app"]
