FROM golang:alpine AS build
WORKDIR /go/src/app
RUN apk add --update --no-cache gcc musl-dev vips-dev

COPY . .
COPY ./libs ./libs

WORKDIR /go/src/app/core
RUN go mod tidy
RUN go mod vendor
RUN go build -o /go/src/app/core main.go

# Final image
FROM alpine:latest
RUN apk add --update --no-cache vips

COPY --from=build /go/src/app/core/main /go/src/app/
ADD https://github.com/golang/go/raw/master/lib/time/zoneinfo.zip /zoneinfo.zip

EXPOSE 8081

ENTRYPOINT ["/go/src/app/main"]
