all:
	@go build -o countdown main.go 

clean:
	@rm -f ./countdown

container: all
	@docker build -t countdown-timer .

.PHONY: all container
