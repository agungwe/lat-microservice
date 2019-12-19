# transaction-microservice
Transaction Microservice

# How to run

- Set up .env
```
cp .env.example .env
```

- Consul
```
consul agent -dev
```

- server
```
go run main.go
```
