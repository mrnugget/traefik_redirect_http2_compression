FROM golang:alpine3.6 AS binary
ADD ./ /app
WORKDIR /app
RUN go build -o http main.go

FROM alpine:3.6
WORKDIR /app
EXPOSE 3000
COPY --from=binary /app/http /app
CMD ["/app/http"]
