# FileName 
BINARY = myapp

# Source Directory

SRC_DIR = .

# Build the project
# 1. Run tests
# 2. Generate code coverage report
test:
	@echo "Running tests..."
	go test ./...
	@echo "Tests completed"
	go test ./... -coverprofile=coverage.out
	@echo "Coverage report generated"
	@echo "Opening coverage report in HTML..."
	go tool cover -html=coverage.out
# Comando para correr staticcheck
check:
	@echo "Running staticcheck..."
	staticcheck ./...