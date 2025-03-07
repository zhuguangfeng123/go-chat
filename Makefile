.PHONY: docker
docker:
	@GOOS=linux go build -o gochat .
	@docker rmi zgf/gochat:v0.0.1
	@docker build -t zgf/gochat:v.0.0.1 .

mock:
	@mockgen -source .\internal\service\user\user.go -package svcmocks -destination .\internal\service\mocks\user.mock.gen.go
	@mockgen -source .\internal\service\code\code.go -package svcmocks -destination .\internal\service\mocks\code.mock.gen.go
	@mockgen -source .\internal\repository\code.go -package repomocks -destination .\internal\repository\mocks\code.mock.gen.go
	@mockgen -source .\internal\repository\user.go -package repomocks -destination .\internal\repository\mocks\user.mock.gen.go
	@mockgen -source .\internal\repository\dao\user.go -package daomocks -destination .\internal\repository\dao\mocks\user.mock.gen.go
	@mockgen -source .\internal\repository\cache\user.go -package cachemocks -destination .\internal\repository\cache\mocks\user.mock.gen.go


	@mockgen -package redismocks -destination .\internal\repository\cache\redismocks\cmdable.mock.gen.go github.com/redis/go-redis/v9 Cmdable
	@go mod tidy