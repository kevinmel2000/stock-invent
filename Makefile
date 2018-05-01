CURR_DIR = $(shell pwd)
MAIN_PATH = $(CURR_DIR)/main.go

all: update run

update:
	@dep ensure -v

run:
	@echo "RUNNING..."
	@go run $(MAIN_PATH)
