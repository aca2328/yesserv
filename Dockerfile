FROM golang AS builder
COPY . /go/src/github.com/aca2328/yesserv
WORKDIR /go/src/github.com/aca2328/yesserv
RUN CGO_ENABLED=0 go build -ldflags="-w -s"

FROM scratch
COPY --from=builder /go/src/github.com/aca2328/yesserv/yesserv /yesserv
CMD ["/yesserv"]