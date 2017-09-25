all: router db

.PHONY: router
router:
	go build -o router ready.go router.go

.PHONY: db
db:
	go build -o db ready.go db.go
