FROM golang:alpine as builder

RUN apk add --no-cache ca-certificates git

WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./

RUN go install -v -ldflags="-s -w" .


FROM alpine

RUN apk --no-cache add ca-certificates

COPY --from=builder /go/bin/redir /bin/redir

EXPOSE 5000
CMD ["/bin/redir"]
