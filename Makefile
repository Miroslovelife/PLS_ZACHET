.PHONY: all test clean

# Variables
TEST_DIRS := ./internal/models ./internal/repositories ./internal/services

# Default target
all: test

# Run tests
test:
	@echo "Running tests..."
	@for dir in $(TEST_DIRS); do \
		echo "Testing $$dir..."; \
		go test $$dir || exit 1; \
	done

# Clean generated files (if any)
clean:
	@echo "Cleaning up..."
	@rm -rf ./bin