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
	go run ./tools/legacytest

generate-parser:
	go run ./tools/antlrgen -jar $(ANTLR_JAR) -grammar $(GRAMMAR) -out $(PARSER_OUT) -version $(ANTLR_VERSION)

modern-test:
	go test ./...

clean:
	rm -rf $(BUILD_DIR) tests/generated modern-tests/generated
