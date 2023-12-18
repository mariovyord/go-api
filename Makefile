.PHONY: dev

dev:
	docker-compose up -d
	air

.PHONY: stop

stop:
	docker-compose down