
all: compose-up sleep test compose-down

.PHONY: compose-up
compose-up:
	export $$(cat .env | xargs) && \
	docker-compose up -d

# to make sure that the database is ready before running the tests
.PHONY: sleep
sleep:
	sleep 1

.PHONY: compose-down
compose-down:
	export $$(cat .env | xargs) && \
	docker-compose down

.PHONY: test
test:
	export $$(cat .env | xargs) && \
	go test -v -count=1 -cover ./...
