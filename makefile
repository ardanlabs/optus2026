# Check to see if we can use ash, in Alpine images, or default to BASH.
SHELL_PATH = /bin/ash
SHELL = $(if $(wildcard $(SHELL_PATH)),/bin/ash,/bin/bash)

# ==============================================================================
# Detect operating system and set the appropriate open command

UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Darwin)
	OPEN_CMD := open
else
	OPEN_CMD := xdg-open
endif

# ==============================================================================
# Modules support

tidy:
	go mod tidy
	go mod vendor

run-help:
	go run api/services/sales/main.go --help

run:
	go run api/services/sales/main.go | go run api/tooling/logfmt/main.go

statsviz:
	$(OPEN_CMD) http://localhost:3010/debug/statsviz

metrics-view:
	expvarmon -ports="localhost:3010" -vars="build,requests,goroutines,errors,panics,mem:memstats.HeapAlloc,mem:memstats.HeapSys,mem:memstats.Sys"

curl:
	curl -il -X GET http://localhost:3000/test
