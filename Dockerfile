# build stage
FROM golang:1.12.5-alpine3.9 as build-stage
WORKDIR /app

COPY go-omdb/main.go ./ 

RUN go build main.go

# production stage
FROM alpine:3.9
COPY --from=build-stage /app/main /
EXPOSE 8080
#CMD ["/main"]
CMD while true;do echo "wait 2..";sleep 2; done
