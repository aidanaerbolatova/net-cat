server: 
	go run ./cmd/server/.
client:
	go run ./cmd/client/.
build:
	docker build --rm -t tcp-chat .
	docker image prune --filter label=stage=builder -f

run:
	docker run --rm --name tcp-chat -p 8989:8989 tcp-chat

.PHONY: server client build run