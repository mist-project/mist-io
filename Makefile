v := false
html := false
go_test_flags := -tags no_test_coverage
go_test_coverage_flags := -func=coverage/coverage.out

# Add verbose flag if v=true
ifeq ($(v),true)
    go_test_flags += -v
endif

# Add 
ifeq ($(html),true)
    go_test_coverage_flags = -html=coverage/coverage.out
endif

run: compile-protos
	@go run src/main.go

build: compile-protos
	@go build -o bin/mist src/main.go

live-run: compile-protos
	@air --build.cmd "go build -o bin/mist src/main.go" --build.bin "./bin/mist"

compile-protos cp:
	@buf generate

# ----- TESTS -----

run-tests:
	@go test -cover -race ./... | grep -v 'testutil' | grep -v 'src/protos'


tbreak:
	go test ./... -run "$(t)"
	
tests t: test-auth test-ws test-message

test-auth:
	@echo -----------------------------------------
	@go test mist-io/src/auth -coverprofile=coverage/coverage.out  $(go_test_flags)
	@go tool cover $(go_test_coverage_flags)


test-ws:
	@echo -----------------------------------------
	@go test mist-io/src/api/ws -coverprofile=coverage/coverage.out  $(go_test_flags)
	@go tool cover $(go_test_coverage_flags)

test-message:
	@echo -----------------------------------------
	@go test mist-io/src/message -coverprofile=coverage/coverage.out  $(go_test_flags)
	@go tool cover $(go_test_coverage_flags)

test-internal:
	@echo -----------------------------------------
	@go test mist-io/src/internal/... -coverprofile=coverage/coverage.out  $(go_test_flags)
	@go tool cover $(go_test_coverage_flags)

# ----- FORMAT -----
lint:
	golangci-lint run --disable-all -E errcheck

lint-proto:
	@buf lint
