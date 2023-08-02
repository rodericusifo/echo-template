# COLOR VARIABLE
GREEN=\033[0;32m
RED=\033[0;31m
BLUE=\033[0;34m
LIGHT_BLUE=\033[1;34m
ORANGE=\033[0;33m
NOCOLOR=\033[0m

# STYLE VARIABLE
BLUE_TRIPLE_EQUALS=$(LIGHT_BLUE)===$(NOCOLOR)

# STYLE FUNCTION
define log_action
$(BLUE_TRIPLE_EQUALS) $(ORANGE)${1}$(NOCOLOR) $(BLUE_TRIPLE_EQUALS)
endef

gen-wire:
	@echo "$(call log_action,Generate Wire)"
	wire wire/core/service/auth/wire.go
	wire wire/core/service/employee/wire.go
	wire wire/core/resource/user/wire.go
gen-mock:
	@echo "$(call log_action,Generate Mock)"
	mockery --all

test-cover: gen-mock
	@echo "$(call log_action,Test Coverage)"
	go test `go list ./... | grep -v mocks` -cover -coverprofile=coverage.out -covermode=count

start-dev: gen-wire
	@echo "$(call log_action,Start Program)"
	go run cmd/main.go -env dev

build: gen-wire
	@echo "$(call log_action,Build Program)"
	go build -o dist/main cmd/main.go