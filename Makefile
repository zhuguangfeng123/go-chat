.PHONY: docker
docker:
	@GOOS=linux go build -o gochat .
	@docker rmi zgf/gochat:v0.0.1
	@docker build -t zgf/gochat:v.0.0.1 .