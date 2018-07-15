# AWS AppSync Resolvers w/ CloudFormation, Lambda & SAM

[![MIT License](https://badgen.now.sh/badge/License/MIT/blue)](https://github.com/sbstjn/appsync-resolvers-example/blob/master/LICENSE.md)
[![Read Tutorial](https://badgen.now.sh/badge/Read/Tutorial/orange)](https://sbstjn.com/serverless-graphql-with-appsync-and-lambda.html)
[![Code Library](https://badgen.now.sh/badge/Code/Library/cyan)](https://github.com/sbstjn/appsync-resolvers)

Fully working GraphQL API example project using [appsync-resolvers] for AWS AppSync and *ready to be deployed* with CloudFormation using the [Serverless Application Model]. Includes AWS Lambda functions for custom `Query` and `Field` resolvers written in Go. You only need the `aws` CLI application and no other third-party frameworks!&nbsp; 🎉

See [Serverless GraphQL with AWS AppSync and Lambda](https://sbstjn.com/serverless-graphql-with-appsync-and-lambda.html) on [sbstjn.com](https://sbstjn.com) for a detailed guide how to set up and configure this project. Or just run `make configure build package deploy` and you are ready to go …

## Schema

```graphql
type Person {
    id: Int!
    name: String!
    age: Int!

    friends: [Person!]!
}

type Query {
    people: [Person!]!
    person(id: Int): Person!
}

schema {
    query: Query
}
```

## Configuration

The `Makefile` contains all tasks to set up the CloudFormation stack.

```bash
# Install Go dependencies
$ > make install

# Create S3 Bucket to store deploy artifacts
$ > make configure

# Build go binary for AWS Lambda
$ > make build

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

#### Request list of all people

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

#### Request specific person

```bash
$ > curl \
    -XPOST https://3mhugdjvrzeclk5ssrc7qzjpxn.appsync-api.eu-west-1.amazonaws.com/graphql \
    -H "Content-Type:application/graphql" \
    -H "x-api-key:da2-jlewwo38ojcrfasc3dpaxqgxcc" \
    -d '{ "query": "query { person(id: 2) { name friends { name } } }" }' | jq
```

```json
{
  "data": {
    "person": {
      "name": "Paul Gascoigne",
      "friends": [
        {
          "name": "Frank Ocean"
        }
      ]
    }
  }
}
```

## License

Feel free to use the code, it's released using the [MIT license](LICENSE.md).

## Contribution

You are welcome to contribute to this project! 😘 

To make sure you have a pleasant experience, please read the [code of conduct](CODE_OF_CONDUCT.md). It outlines core values and beliefs and will make working together a happier experience.

[appsync-resolvers]: https://github.com/sbstjn/appsync-resolvers
[Serverless Application Model]: https://github.com/awslabs/serverless-application-model
