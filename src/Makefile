# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARY_NAME=main
ZIPCMD=zip
ZIP_NAME=../bin/function.zip

# Build binary that will work in Lambda
build:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) -v
	$(ZIPCMD) $(ZIP_NAME) $(BINARY_NAME)

# Delete binary and zip file.
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(ZIP_NAME)

