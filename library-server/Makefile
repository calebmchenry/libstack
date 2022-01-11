all: restore-deps test build

help:  ## Prints the help/usage docs.
	@awk 'BEGIN {FS = ":.*##"; printf "Usage: make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

test:  ## Runs all tests.
	go1.18beta1 test -short -v ./...

vet:  ## Runs Golang's static code analysis
	go1.18beta1 vet ./...

clean:  ## Removes build/test outputs.
	rm -rf bin *.test

restore-deps:  ## Tidies up the go module.
	go1.18beta1 mod tidy	
