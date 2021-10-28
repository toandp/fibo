# Fibo - Go Fiber API Boilerplate
> A starter project with Golang, Fiber and Gorm

Golang Fiber boilerplate with MySQL resource. Supports multiple configuration environments.

### Features
- Basic Auth with Login, Register
- Email confirmation on Registration
- Email forgot password
- REST API Authentication with JWT
- PostgresSQL or MySQL with GORM V2
- Logging via zap with file rotation
- Use of Redis for Cache and Session
- Hot Reload with Air
- Easy Config Settings based on .env
- Setup for Docker
- Easy and Almost Zero Downtime Production Deployment

### Boilerplate structure
```
├── api
|   └── v1
│       ├── login.go
│       ├── password.go
|       └── register.go
├── config
│   └── config.go
├── db
│   └── db.go
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── log
│   └── log.go
├── router
│   └── router.go
├── tmp
├── .air.toml
├── main.go
```

### Installation
- Clone the repo git clone https://github.com/toandp/fibo.git
- Make sure you have installed: Redis, MySQL or Postgres
- Copy .env.sample to .env

### Development
```bash
air -c .air.toml
```
Go to  [http://localhost:3000/](http://localhost:3000/)

### Production
```bash
docker build -t fibo .
docker run -d -p 3000:3000 fibo
```

Thanks to following libraries:

* [Fiber](https://github.com/gofiber/fiber/v2)
* [Gorm](https://github.com/go-gorm/gorm)
* [Zap](https://github.com/uber-go/zap)
* [Air](https://github.com/cosmtrek/air)