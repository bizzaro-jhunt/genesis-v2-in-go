# Run me to verify that all tests pass and all binaries are buildable before pushing!
# If you do not, then Travis will be sad.

BUILD_TYPE?=build

# Everything; this is the default behavior
#all: format tests shield plugins

genesis:
	@go fmt .
	@go build -o genesis .
	@echo "Your Genesis is ready in ./$@, $$(whoami)..."

.PHONY: genesis
