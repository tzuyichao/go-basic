# zabbix-agent-test

```
docker run --rm --name zabbix-agent -e ZBX_SERVER_HOST="192.168.1.1" -p 10050:10050 -d zabbix/zabbix-agent:alpine-6.0-latest
```

```
go build -o zabbix-agent-test.exe .\main.go
```




```
.\zabbix-agent-test.exe localhost 10050 agent.ping
```

## Trouble shoot

```
docker logs zabbix-agent
```

```
Starting Zabbix Agent [a2783e7b7337]. Zabbix 6.0.31 (revision b6d9375).
Press Ctrl+C to exit.

     7:20240706:025658.013 Starting Zabbix Agent [a2783e7b7337]. Zabbix 6.0.31 (revision b6d9375).
     7:20240706:025658.013 **** Enabled features ****
     7:20240706:025658.013 IPv6 support:          YES
     7:20240706:025658.014 TLS support:           YES
     7:20240706:025658.014 **************************
     7:20240706:025658.014 using configuration file: /etc/zabbix/zabbix_agentd.conf
     7:20240706:025658.014 agent #0 started [main process]
    79:20240706:025658.014 agent #1 started [collector]
    80:20240706:025658.014 agent #2 started [listener #1]
    81:20240706:025658.014 agent #3 started [listener #2]
    82:20240706:025658.014 agent #4 started [listener #3]
    83:20240706:025658.014 agent #5 started [active checks #1]
    83:20240706:025658.016 Unable to connect to [192.168.1.1]:10051 [cannot connect to [[192.168.1.1]:10051]: [111] Connection refused]
    83:20240706:025658.016 Active check configuration update started to fail
    82:20240706:025704.682 failed to accept an incoming connection: connection from "192.168.127.1" rejected, allowed hosts: "192.168.1.1"
```