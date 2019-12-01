FROM golang:1.13 AS build
RUN apt-get update && apt-get install unzip
WORKDIR /go/src/github.com/marjamis/UtilityBelt/
COPY go.* ./
RUN go mod download
COPY *.go ./
COPY kubernetes/ ./kubernetes/
COPY redis/ ./redis/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ub .

FROM alpine:3.10.3
LABEL maintainer=marjamis
RUN apk --no-cache add ca-certificates
# USER nobody
WORKDIR /app/
COPY --from=build /go/src/github.com/marjamis/UtilityBelt/ub .
COPY ./static/ /app/static/
ENTRYPOINT ["./ub"]
