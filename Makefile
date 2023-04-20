
all: compose-up test compose-down

.PHONY: compose-up
compose-up:
	export $$(cat .env | xargs) && \
	docker-compose up -d
	sleep 1

.PHONY: compose-down
compose-down:
	export $$(cat .env | xargs) && \
	docker-compose down

.PHONY: test
test:
	export $$(cat .env | xargs) && \
	go test -v -count=1 -cover ./...
