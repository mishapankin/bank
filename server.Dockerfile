FROM golang:1.19 as build_base

WORKDIR /tmp/app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .


FROM alpine:3.16
COPY --from=build_base /tmp/app/main /app/


ENV key=value

EXPOSE 8080
CMD ["/app/main"]