server:
	go run ./cmd/http/.

exp:
	go run ./cmd/http/exp.go

db:
	docker compose up

.PHONY: server exp