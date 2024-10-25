GIT_MSG ?= "Atualização do código"

commit:
	git add .
	git commit -m "$(GIT_MSG)"
	git push

server:
	go run ./cmd/http/.

exp:
	go run ./cmd/http/exp.go

db:
	docker compose up

.PHONY: server exp