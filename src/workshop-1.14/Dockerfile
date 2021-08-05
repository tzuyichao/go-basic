# builder
FROM golang AS builder

COPY main.go .
RUN go env -w GO111MODULE=auto
RUN go build -o /hello-world

# app
FROM alpine
CMD ["/entry/hello-world"]

WORKDIR entry
COPY --from=builder /hello-world .
RUN chmod +x hello-world