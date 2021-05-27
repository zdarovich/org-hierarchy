FROM golang:1.15-alpine as builder

RUN mkdir /build
COPY . /build/

WORKDIR /build

RUN go build -o org-hierarchy cmd/server/main.go

FROM alpine:latest

COPY --from=builder /build/org-hierarchy /app/
WORKDIR /app
RUN chmod +x ./org-hierarchy

CMD ["./org-hierarchy"]