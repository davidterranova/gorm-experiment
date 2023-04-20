
.PHONY: compose-up
compose-up:
	export $$(cat .env | xargs) && \
	docker-compose up -d

.PHONY: compose-down
compose-down:
	export $$(cat .env | xargs) && \
	docker-compose down
