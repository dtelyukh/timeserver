Time server RFC 868 implementation with command line client.

Run server:
```
cd cmd/server
go build
sudo ./server -p 37
```

Test it with client:
```
cd cmd/client
go build
./client localhost 37
```
