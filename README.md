# AppSync Router Example

> Example project using [appsync-router] to create an AWS AppSync GraphQL API using an AWS Lambda function as a custom data source and resolver written in Go.

## Configuration

The `Makefile` contains all tasks to set up a CloudFormation stack using the [Serverless Application Model] by Amazon. You only need the `aws` CLI application.

```bash
# Create S3 Bucket to store deploy artifacts
$ > make configure

# Build go binary for AWS Lambda
$ > make build-lambda

# Create deployable artifact
$ > make package

# Deploy CloudFormation stack
$ > make deploy
```

## Usage

```bash
# Show CloudFormation stack output
$ > make outputs

[
  {
    "OutputKey": "APIKey",
    "OutputValue": "da2-jlewwo38ojcrfasc3dpaxqgxcc",
    "Description": "API Key"
  },
  {
    "OutputKey": "GraphQL",
    "OutputValue": "https://3mhugdjvrzeclk5ssrc7qzjpxn.appsync-api.eu-west-1.amazonaws.com/graphql",
    "Description": "GraphQL URL"
  }
]
```

### Send GraphQL Requests

To request a list of all people, use the `people` query:

```bash
$ > curl \
    -XPOST https://3mhugdjvrzeclk5ssrc7qzjpxn.appsync-api.eu-west-1.amazonaws.com/graphql \
    -H "Content-Type:application/graphql" \
    -H "x-api-key:da2-jlewwo38ojcrfasc3dpaxqgxcc" \
    -d '{ "query": "query { people { name } }" }' | jq
```

```json
{
  "data": {
    "people": [
      {
        "name": "Frank Ocean"
      },
      {
        "name": "Paul Gascoigne"
      },
      {
        "name": "Uwe Seeler"
      }
    ]
  }
}
```

To request a specific person, use the `person` query:

```bash
$ > curl \
    -XPOST https://3mhugdjvrzeclk5ssrc7qzjpxn.appsync-api.eu-west-1.amazonaws.com/graphql \
    -H "Content-Type:application/graphql" \
    -H "x-api-key:da2-jlewwo38ojcrfasc3dpaxqgxcc" \
    -d '{ "query": "query { person(id: 2) { name } }" }' | jq
```

```json
{
  "data": {
    "person": {
      "name": "Paul Gascoigne"
    }
  }
}
```

## License

Feel free to use the code, it's released using the [MIT license](LICENSE.md).

## Contribution

You are welcome to contribute to this project! 😘 

To make sure you have a pleasant experience, please read the [code of conduct](CODE_OF_CONDUCT.md). It outlines core values and beliefs and will make working together a happier experience.

[appsync-router]: https://github.com/sbstjn/appsync-router
[Serverless Application Model]: https://github.com/awslabs/serverless-application-model