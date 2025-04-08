default:
	@echo "Building executable"
	@go build -o ./build/survival_game
	@echo "Starting executable"
	@./build/survival_game
