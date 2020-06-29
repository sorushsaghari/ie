go mod vendor
cp .env.sample .env
go build ./cmd/html
./html