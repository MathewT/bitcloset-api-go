# bitcloset-api-go

```bash
go build .
```

```bash
mv bitcloset-api-go bin
```
```bash
docker rm $(docker ps -aq); docker rmi bitcloset-api-go
```

```bash
docker build . -t bitcloset-api-go
```

```bash
docker run -it bitcloset-api-go
```

```bash 
curl -XGET -d '{"color":"green","message":"Test from Mathew"}' -H 'Content-Type: application/json' http://172.17.0.2:8000/api/v1/bitcloset/
```
