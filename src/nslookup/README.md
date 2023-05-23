## Execution


```powershell=
$env:GODEBUG="netdns=9"; go run .\nslookup.go google.com
```

## Source

[How can I do DNS lookup in Go?](https://jameshfisher.com/2017/08/03/golang-dns-lookup/)