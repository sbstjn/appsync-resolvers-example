PROJECT_NAME ?= appsync-router-example
ENV ?= stable

AWS_BUCKET_NAME ?= $(PROJECT_NAME)-artifacts-$(ENV)
AWS_STACK_NAME ?= $(PROJECT_NAME)-stack-$(ENV)
AWS_REGION ?= eu-west-1
GOOS ?= darwin

FILE_TEMPLATE = template.yml
FILE_PACKAGE = package.yml

SCHEMA := $(shell cat schema.graphql)
EXPIRATION = $(shell echo $$(( $(shell date +%s) + 604800 ))) # 7 days from now (timestamp)

clean:
	rm -rf dist

install:
	@ dep ensure

test:
	@ go test ./... -v

build:
	@ go build -o dist/handler_$(GOOS) ./src

build-lambda: 
	@ GOOS=linux make build

configure:
	@ aws s3api create-bucket \
		--bucket $(AWS_BUCKET_NAME) \
		--region $(AWS_REGION) \
		--create-bucket-configuration LocationConstraint=$(AWS_REGION)

package:
	@ aws cloudformation package \
		--template-file $(FILE_TEMPLATE) \
		--s3-bucket $(AWS_BUCKET_NAME) \
		--region $(AWS_REGION) \
		--output-template-file $(FILE_PACKAGE)

deploy:
	@ aws cloudformation deploy \
		--template-file $(FILE_PACKAGE) \
		--region $(AWS_REGION) \
		--capabilities CAPABILITY_IAM \
		--stack-name $(AWS_STACK_NAME) \
		--force-upload \
		--parameter-overrides \
			ParamProjectName=$(PROJECT_NAME) \
			ParamSchema="$(SCHEMA)" \
			ParamKeyExpiration=$(EXPIRATION) \
			ParamENV=$(ENV)

describe:
	@ aws cloudformation describe-stacks \
		--region $(AWS_REGION) \
		--stack-name $(AWS_STACK_NAME)

outputs:
	@ make describe \
		| jq -r '.Stacks[0].Outputs'

.PHONY: clean install build build-lambda configure package deploy describe output
