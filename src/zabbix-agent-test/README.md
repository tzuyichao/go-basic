# zabbix-agent-test

```
docker run --name zabbix-agent -e ZBX_SERVER_HOST="127.0.0.1,192.168.127.1" -p 10050:10050 -d zabbix/zabbix-agent:alpine-6.0-latest
```

```
go build -o zabbix-agent-test.exe .\main.go
```


```
.\zabbix-agent-test.exe localhost 10050 agent.ping
```