FROM golang:1.20.0-alpine

# credit goes to https://github.com/confluentinc/confluent-kafka-go/issues/156#issuecomment-376311118
RUN apk update && apk upgrade

RUN apk add --update --no-cache alpine-sdk bash ca-certificates \
      libressl \
      tar \
      git openssh openssl yajl-dev zlib-dev cyrus-sasl-dev openssl-dev build-base coreutils
WORKDIR /root
RUN git clone https://github.com/edenhill/librdkafka.git
WORKDIR /root/librdkafka
RUN /root/librdkafka/configure
RUN make
RUN make install
#For golang applications
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

WORKDIR /app

COPY . .

RUN go get -d -v github.com/confluentinc/confluent-kafka-go/v2/kafka

# credit goes to https://github.com/confluentinc/confluent-kafka-go/issues/454#issuecomment-613387798
# -tags musl is required
# but according https://github.com/confluentinc/confluent-kafka-go#installing-librdkafka
# > After installing librdkafka you will need to build your Go application with -tags dynamic.
# also works for me without -tags musl
RUN GOOS=linux go build -tags dynamic -a -installsuffix cgo -o main .

CMD ["/app/main"]