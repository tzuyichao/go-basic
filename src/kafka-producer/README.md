# Memo

因為confluent kafka go client是`librdkafka`的wrapper，而這個native code library目前不支援Windows，所以用docker來驗證一下。

```
docker build . --tag kafka-producer
```

```
docker run --rm --network host kafka-producer
```