CXX ?= g++
CXXFLAGS ?= -std=c++17 -Isrc -Ibuild
BISON ?= bison
FLEX ?= flex
ANTLR_VERSION := 4.13.1
ANTLR_JAR := tools/antlr/antlr-$(ANTLR_VERSION)-complete.jar
GRAMMAR := grammar/JavaSubset.g4
PARSER_OUT := internal/parser/gen

BUILD_DIR := build
TARGET := $(BUILD_DIR)/javago

.PHONY: all clean test generate-parser modern-test

all: $(TARGET)

$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

$(BUILD_DIR)/parser.cpp $(BUILD_DIR)/parser.hpp: src/parser.y | $(BUILD_DIR)
	$(BISON) -d -o $(BUILD_DIR)/parser.cpp src/parser.y

$(BUILD_DIR)/lexer.cpp: src/lexer.l $(BUILD_DIR)/parser.hpp | $(BUILD_DIR)
	$(FLEX) -o $(BUILD_DIR)/lexer.cpp src/lexer.l

$(TARGET): $(BUILD_DIR)/parser.cpp $(BUILD_DIR)/lexer.cpp src/ast.cpp src/diagnostics.cpp src/generator.cpp src/main.cpp
	$(CXX) $(CXXFLAGS) $(BUILD_DIR)/parser.cpp $(BUILD_DIR)/lexer.cpp src/ast.cpp src/diagnostics.cpp src/generator.cpp src/main.cpp -o $(TARGET)

test: all
	pwsh ./test.ps1

$(ANTLR_JAR):
	powershell -NoProfile -ExecutionPolicy Bypass -Command "New-Item -ItemType Directory -Force tools/antlr | Out-Null; Invoke-WebRequest -Uri https://www.antlr.org/download/antlr-$(ANTLR_VERSION)-complete.jar -OutFile $(ANTLR_JAR)"

generate-parser: $(ANTLR_JAR)
	powershell -NoProfile -ExecutionPolicy Bypass -File tools/antlr/generate-parser.ps1 -Jar $(ANTLR_JAR) -Grammar $(GRAMMAR) -OutputDir $(PARSER_OUT)

modern-test:
	go test ./...

clean:
	rm -rf $(BUILD_DIR) tests/generated modern-tests/generated
