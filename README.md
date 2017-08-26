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
curl -XGET -d '{"color":"green","message":"Test from curl"}' -H 'Content-Type: application/json' http://172.19.0.4:80/api/v1/bitcloset/
```

```bash
curl -v -XGET http://172.19.0.4:80/api/v1/bitcloset/?force=true -d '{"users":[{"md5":"016aae09d8a73b52fd356c1ebe9c33b6"},{"sha256":"183b6d1af7a59cc4b90983a5d3e0a7f164ce778e1002ba19f29fe0fe59f5ff98"}]}'  -H "Content-Type:application/json" -H "Authorization:Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX3VpZCI6InRlc3QiLCJpYXQiOjE0NTc2NDE0NDAsImV4cCI6MTQ1NzcyNzg0MH0.OcyguLVwPbZV2rMpZQwuolyiRpAtOltKBAQpxXTNPD0"
```
