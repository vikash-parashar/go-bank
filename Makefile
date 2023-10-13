# remove the previous built binary to make sure that latest changes should be reflected in build binary
build: 
	@rm -rf go-bank 
	@go build -o go-bank

#  run the newley builted binary file
run:
	@./go-bank

# to run all tests
test:
	@go test -v ./...
# adding `@` will hide these tasks from priting