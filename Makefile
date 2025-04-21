default:
	@echo "Building executable"
	@go build -o ./build/survival_game
	@echo "Starting executable"
	@./build/survival_game

dev:
	@echo "Starting Dev Mode"
	@air run main.go
