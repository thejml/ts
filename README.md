# ts
Add timestamps to any piped input

## Example
```
 $ ping www.google.com|ts
Jan  7 14:43:29.826(     40ms): PING www.google.com (172.217.12.228): 56 data bytes
Jan  7 14:43:29.826(       0s): 64 bytes from 172.217.12.228: icmp_seq=0 ttl=55 time=11.973 ms
Jan  7 14:43:30.829(   1.002s): 64 bytes from 172.217.12.228: icmp_seq=1 ttl=55 time=13.237 ms
Jan  7 14:43:31.832(   1.003s): 64 bytes from 172.217.12.228: icmp_seq=2 ttl=55 time=15.650 ms
Jan  7 14:43:32.830(    998ms): 64 bytes from 172.217.12.228: icmp_seq=3 ttl=55 time=13.950 ms
```

## Building
With a properly setup go lang GOPATH:

```
go build ts.go
```

If using golang configured for module support, setup your GOPATH and then:
```
go mod init
go build ts.go
```


