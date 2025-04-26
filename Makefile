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
	@go build -o $(MAIN) $(SRC)

air: deps
	@echo "♻️  Running with Air..."
	@air

clean:
	@echo "🧹 Cleaning up..."
	@rm -f $(MAIN)

test:
	@echo "🧪 Running tests..."
	@go test ./... -v
	@echo "✅ All tests passed!"