GO  ?= go
CD  ?= cd
NPM ?= npm

all:
	$(GO) run ./cmd/notepanel/main.go

install:
	$(CD) ./frontend; $(NPM) run build