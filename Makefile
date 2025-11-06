# ANTLR version and jar file
ANTLR_VERSION := 4.13.2
ANTLR_JAR_NAME := antlr-$(ANTLR_VERSION)-complete.jar
ANTLR_URL := https://www.antlr.org/download/$(ANTLR_JAR_NAME)

# Directories
TOOLS_DIR := tools
GRAMMAR_DIR := grammar
PARSER_DIR := parser

# Full path to ANTLR jar
ANTLR_JAR := $(TOOLS_DIR)/$(ANTLR_JAR_NAME)

# Grammar file
GRAMMAR := $(GRAMMAR_DIR)/Flow.g4

# Default target
.PHONY: all
all: gen

# Download ANTLR if not present
$(ANTLR_JAR):
	@echo "Creating tools directory..."
	@mkdir -p $(TOOLS_DIR)
	@echo "Downloading ANTLR $(ANTLR_VERSION)..."
	@cd $(TOOLS_DIR) && curl -O $(ANTLR_URL)

# Generate parser from grammar
.PHONY: gen
gen: $(ANTLR_JAR)
	@echo "Generating Go parser from grammar..."
	@mkdir -p $(PARSER_DIR)
	@cd $(GRAMMAR_DIR) && java -jar ../$(ANTLR_JAR) -Dlanguage=Go -o ../$(PARSER_DIR) -package parser -visitor Flow.g4
	@echo "Parser generation complete!"

# Generate JavaScript parser
.PHONY: gen-js
gen-js: $(ANTLR_JAR)
	@echo "Generating JavaScript parser from grammar..."
	@mkdir -p web/parser
	@java -jar $(ANTLR_JAR) -Dlanguage=JavaScript -o web/parser -visitor $(GRAMMAR)
	@echo "Moving generated files to correct location..."
	@if [ -d "web/parser/grammar" ]; then \
		mv web/parser/grammar/* web/parser/; \
		rmdir web/parser/grammar; \
	fi
	@echo "JavaScript parser generation complete!"

# Clean generated files
.PHONY: clean
clean:
	@echo "Cleaning generated parser files..."
	@rm $(PARSER_DIR)/*
	@echo "Clean complete!"

# Clean everything including ANTLR jar
.PHONY: clean-all
clean-all: clean
	@echo "Removing ANTLR jar..."
	@rm -f $(ANTLR_JAR)
	@echo "Clean all complete!"

# Run tests
.PHONY: test
test: gen
	@echo "Running tests..."
	@go test ./...
	@echo "Tests complete!"

# Build WASM
.PHONY: wasm
wasm: gen
	@echo "Building WebAssembly module..."
	@mkdir -p web
	@GOOS=js GOARCH=wasm go build -o web/flow-transpiler.wasm wasm.go
	@echo "Downloading WASM support file..."
	@curl -s https://raw.githubusercontent.com/golang/go/release-branch.go1.21/misc/wasm/wasm_exec.js > web/wasm_exec.js
	@echo "WASM build complete!"

# Generate JavaScript files with WASM
.PHONY: wasm-js
wasm-js: wasm
	@echo "Generating JavaScript files..."
	@chmod +x scripts/generate-wasm-js.sh
	@./scripts/generate-wasm-js.sh
	@rm -f web/flow-transpiler.wasm web/wasm_exec.js
	@echo "Web files ready! Open web/flow-playground.html in your browser"

# Alias for wasm-js (main web target)
.PHONY: web
web: wasm-js
