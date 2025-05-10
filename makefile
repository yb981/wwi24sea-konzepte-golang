# Ausführbare Demos
CMDS := calculator datastructures-demo edsl-functions-demo edsl-vector-demo password concurrency

# Binaries in ./bin
BIN_DIR := bin

# build all
build:
	@mkdir -p $(BIN_DIR)
	@for cmd in $(CMDS); do \
		echo "Building $$cmd..."; \
		go build -o $(BIN_DIR)/$$cmd ./cmd/$$cmd/*.go || exit 1; \
	done

# run specific demo, e.g., make run CMD=calculator
run:
	@$(BIN_DIR)/$(CMD)

# test all packages (nicht cmd/)
test:
	go test ./...

# coverage
coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

coverage-report:
	go tool cover -func=coverage.out

clean:
	rm -rf $(BIN_DIR) coverage.out coverage.html