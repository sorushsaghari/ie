# Internet engineering project

## requirements:
go 1.8 +

postgresql

## running process:
1.install dependencies by
```
go mod vendor
```

2.setting configuration:

create a .env file from .env sample and fill it with your own app configuration

3.run application"

`for debug mode`
```
go run cmd/html/main.go
```
`or to build`
```
go build ./cmd/html
./html
```

`or easily run :` __run.sh__ 