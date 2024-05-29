# Makefile

# Define the output directory and binary name
OUT_DIR := out
OUT_BINARY := $(OUT_DIR)/out

# Phony targets are not real files
.PHONY: build run start

# Build target: Compiles the Go project and places the output in the specified directory
build:
	@echo "Building the project..."
	@echo "Output directory: $(OUT_DIR)"
	@echo "Output binary: $(OUT_BINARY)"
	@mkdir -p $(OUT_DIR)
	@go build -o $(OUT_BINARY)
	@echo "************\tBuild Ready\t************"

# Run target: Executes the compiled binary
run:
	@echo "Running the binary: $(OUT_BINARY)"
	@echo "************\tStarting\t************"
	@./$(OUT_BINARY)

# Start target: Runs the build target and then the run target
start: build run