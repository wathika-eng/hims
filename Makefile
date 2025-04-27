MAIN = main
SRC = main.go

.PHONY: run build air clean deps

deps:
	@echo "🔄 Tidying dependencies..."
	@go mod tidy

run: deps
	@echo "🚀 Running $(SRC)..."
	@go run $(SRC)

build: deps
	@echo "🔨 Building binary..."
	@go build -o bin/$(MAIN) $(SRC)

air: deps
	@echo "♻️  Running with Air..."
	@air

build-docker: 
	@echo "🐳 Building Docker image..."
	@COMPOSE_BAKE=true docker compose up --build

run-docker: 
	@echo "🐳 Building Docker image..."
	@COMPOSE_BAKE=true docker compose up 

clean:
	@echo "🧹 Cleaning up..."
	@rm -f $(MAIN)

lint:
	@echo "🔍 Running linters..."
	@golangci-lint fmt ./...
	@golangci-lint run ./...
	
test: deps
	@echo "🧪 Running tests..."
	@go test ./... -v
	@echo "✅ All tests passed!"