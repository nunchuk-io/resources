# fully-qualified path to this Makefile
MKFILE_PATH := $(realpath $(lastword $(MAKEFILE_LIST)))

# fully-qualified path to the current directory
CURRENT_DIR := $(patsubst %/,%,$(dir $(MKFILE_PATH)))

default: help

.PHONY: help
## This help screen
help:
	@printf "Available targets:\n\n"
	@awk '/^[a-zA-Z\-\_0-9%:\\]+/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
          if (helpMessage) { \
            helpCommand = $$1; \
            helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
      gsub("\\\\", "", helpCommand); \
      gsub(":+$$", "", helpCommand); \
            printf "  \x1b[32;01m%-35s\x1b[0m %s\n", helpCommand, helpMessage; \
          } \
        } \
        { lastLine = $$0 }' $(MAKEFILE_LIST) | sort -u
	@printf "\n"


## Build for linux
linux:
	@cd $(CURRENT_DIR)/contact-form && GOOS=linux GOARCH=amd64 go build -o ${CURRENT_DIR}/bin/contact-form

## Deploy to AWS lambda
deploy: linux
	@cd $(CURRENT_DIR)/terraform && terraform init && terraform apply -auto-approve

## Clean
clean:
	@rm -rf $(CURRENT_DIR)/bin/contact-form
