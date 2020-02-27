# mischievous-slack-bot

## project init

### 1. NodeJS install (after using npm)

### 2. serverless Framework install (from npm)

```shell
$ npm install -g serverless

   ┌───────────────────────────────────────────────────┐
   │                                                   │
   │   Serverless Framework successfully installed!    │
   │                                                   │
   │   To start your first project run 'serverless'.   │
   │                                                   │
   └───────────────────────────────────────────────────┘

```

### 3. additional PATH

```shell
$ serverless -v
Framework Core: 1.64.0
Plugin: 3.4.1
SDK: 2.3.0
Components Core: 1.1.2
Components CLI: 1.4.0
```

### 4. init serverless app

from `${GOPATH}/src` directory

```shell
$ serverless create -t aws-go-mod -u https://github.com/kichion/ -p mischievous-slack-bot
Serverless: Generating boilerplate...
Serverless: Generating boilerplate in "/mischievous-slack-bot"
 _______                             __
|   _   .-----.----.--.--.-----.----|  .-----.-----.-----.
|   |___|  -__|   _|  |  |  -__|   _|  |  -__|__ --|__ --|
|____   |_____|__|  \___/|_____|__| |__|_____|_____|_____|
|   |   |             The Serverless Application Framework
|       |                           serverless.com, v1.64.0
 -------'

Serverless: Successfully generated boilerplate for template: "aws-go-mod"
```

### 5. deploy to AWS

rewrite `serverless.yml` (profile use ME)

```yml
provider:
  name: aws
  runtime: go1.x
#   add start
  region: ap-northeast-1
  profile: me
#   add end
```

command execute

```shell
$ make deploy
rm -rf ./bin ./vendor Gopkg.lock
chmod u+x gomod.sh
./gomod.sh
export GO111MODULE=on
env GOOS=linux go build -ldflags="-s -w" -o bin/hello hello/main.go
env GOOS=linux go build -ldflags="-s -w" -o bin/world world/main.go

serverless deploy --verbose
Serverless: Packaging service...
Serverless: Excluding development dependencies...
Serverless: Creating Stack...
Serverless: Checking Stack create progress...
CloudFormation - CREATE_IN_PROGRESS - AWS::CloudFormation::Stack - mischievous-slack-bot-dev
CloudFormation - CREATE_IN_PROGRESS - AWS::S3::Bucket - ServerlessDeploymentBucket
CloudFormation - CREATE_IN_PROGRESS - AWS::S3::Bucket - ServerlessDeploymentBucket
CloudFormation - CREATE_COMPLETE - AWS::S3::Bucket - ServerlessDeploymentBucket
CloudFormation - CREATE_IN_PROGRESS - AWS::S3::BucketPolicy - ServerlessDeploymentBucketPolicy
CloudFormation - CREATE_IN_PROGRESS - AWS::S3::BucketPolicy - ServerlessDeploymentBucketPolicy
CloudFormation - CREATE_COMPLETE - AWS::S3::BucketPolicy - ServerlessDeploymentBucketPolicy
CloudFormation - CREATE_COMPLETE - AWS::CloudFormation::Stack - mischievous-slack-bot-dev
Serverless: Stack create finished...
Serverless: Uploading CloudFormation file to S3...
Serverless: Uploading artifacts...
Serverless: Uploading service mischievous-slack-bot.zip file to S3 (6.12 MB)...
Serverless: Validating template...
Serverless: Updating Stack...
Serverless: Checking Stack update progress...
CloudFormation - UPDATE_IN_PROGRESS - AWS::CloudFormation::Stack - mischievous-slack-bot-dev
CloudFormation - CREATE_IN_PROGRESS - AWS::Logs::LogGroup - WorldLogGroup
CloudFormation - CREATE_IN_PROGRESS - AWS::ApiGateway::RestApi - ApiGatewayRestApi
CloudFormation - CREATE_IN_PROGRESS - AWS::Logs::LogGroup - WorldLogGroup
CloudFormation - CREATE_IN_PROGRESS - AWS::IAM::Role - IamRoleLambdaExecution
CloudFormation - CREATE_IN_PROGRESS - AWS::Logs::LogGroup - HelloLogGroup
CloudFormation - CREATE_COMPLETE - AWS::Logs::LogGroup - WorldLogGroup
CloudFormation - CREATE_IN_PROGRESS - AWS::ApiGateway::RestApi - ApiGatewayRestApi
CloudFormation - CREATE_IN_PROGRESS - AWS::Logs::LogGroup - HelloLogGroup
CloudFormation - CREATE_IN_PROGRESS - AWS::IAM::Role - IamRoleLambdaExecution
CloudFormation - CREATE_COMPLETE - AWS::Logs::LogGroup - HelloLogGroup
CloudFormation - CREATE_COMPLETE - AWS::ApiGateway::RestApi - ApiGatewayRestApi
CloudFormation - CREATE_IN_PROGRESS - AWS::ApiGateway::Resource - ApiGatewayResourceWorld
CloudFormation - CREATE_IN_PROGRESS - AWS::ApiGateway::Resource - ApiGatewayResourceHello
CloudFormation - CREATE_IN_PROGRESS - AWS::ApiGateway::Resource - ApiGatewayResourceWorld
CloudFormation - CREATE_IN_PROGRESS - AWS::ApiGateway::Resource - ApiGatewayResourceHello
CloudFormation - CREATE_COMPLETE - AWS::ApiGateway::Resource - ApiGatewayResourceWorld
CloudFormation - CREATE_COMPLETE - AWS::ApiGateway::Resource - ApiGatewayResourceHello
CloudFormation - CREATE_COMPLETE - AWS::IAM::Role - IamRoleLambdaExecution
CloudFormation - CREATE_IN_PROGRESS - AWS::Lambda::Function - HelloLambdaFunction
CloudFormation - CREATE_IN_PROGRESS - AWS::Lambda::Function - WorldLambdaFunction
CloudFormation - CREATE_IN_PROGRESS - AWS::Lambda::Function - HelloLambdaFunction
CloudFormation - CREATE_IN_PROGRESS - AWS::Lambda::Function - WorldLambdaFunction
CloudFormation - CREATE_COMPLETE - AWS::Lambda::Function - HelloLambdaFunction
CloudFormation - CREATE_COMPLETE - AWS::Lambda::Function - WorldLambdaFunction
CloudFormation - CREATE_IN_PROGRESS - AWS::Lambda::Version - HelloLambdaVersionXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
CloudFormation - CREATE_IN_PROGRESS - AWS::ApiGateway::Method - ApiGatewayMethodHelloGet
CloudFormation - CREATE_IN_PROGRESS - AWS::Lambda::Version - WorldLambdaVersionXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
CloudFormation - CREATE_IN_PROGRESS - AWS::Lambda::Permission - HelloLambdaPermissionApiGateway
CloudFormation - CREATE_IN_PROGRESS - AWS::ApiGateway::Method - ApiGatewayMethodWorldGet
CloudFormation - CREATE_IN_PROGRESS - AWS::ApiGateway::Method - ApiGatewayMethodHelloGet
CloudFormation - CREATE_IN_PROGRESS - AWS::Lambda::Permission - WorldLambdaPermissionApiGateway
CloudFormation - CREATE_IN_PROGRESS - AWS::ApiGateway::Method - ApiGatewayMethodWorldGet
CloudFormation - CREATE_IN_PROGRESS - AWS::Lambda::Permission - HelloLambdaPermissionApiGateway
CloudFormation - CREATE_IN_PROGRESS - AWS::Lambda::Version - HelloLambdaVersionXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
CloudFormation - CREATE_IN_PROGRESS - AWS::Lambda::Permission - WorldLambdaPermissionApiGateway
CloudFormation - CREATE_COMPLETE - AWS::ApiGateway::Method - ApiGatewayMethodHelloGet
CloudFormation - CREATE_IN_PROGRESS - AWS::Lambda::Version - WorldLambdaVersionXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
CloudFormation - CREATE_COMPLETE - AWS::ApiGateway::Method - ApiGatewayMethodWorldGet
CloudFormation - CREATE_COMPLETE - AWS::Lambda::Version - HelloLambdaVersionXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
CloudFormation - CREATE_COMPLETE - AWS::Lambda::Version - WorldLambdaVersionXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
CloudFormation - CREATE_IN_PROGRESS - AWS::ApiGateway::Deployment - ApiGatewayDeployment0000000000000
CloudFormation - CREATE_IN_PROGRESS - AWS::ApiGateway::Deployment - ApiGatewayDeployment0000000000000
CloudFormation - CREATE_COMPLETE - AWS::ApiGateway::Deployment - ApiGatewayDeployment0000000000000
CloudFormation - CREATE_COMPLETE - AWS::Lambda::Permission - HelloLambdaPermissionApiGateway
CloudFormation - CREATE_COMPLETE - AWS::Lambda::Permission - WorldLambdaPermissionApiGateway
CloudFormation - UPDATE_COMPLETE_CLEANUP_IN_PROGRESS - AWS::CloudFormation::Stack - mischievous-slack-bot-dev
CloudFormation - UPDATE_COMPLETE - AWS::CloudFormation::Stack - mischievous-slack-bot-dev
Serverless: Stack update finished...
Service Information
service: mischievous-slack-bot
stage: dev
region: ap-northeast-1
stack: mischievous-slack-bot-dev
resources: 17
api keys:
  None
endpoints:
  GET - https://foo.execute-api.ap-northeast-1.amazonaws.com/dev/hello
  GET - https://foo.execute-api.ap-northeast-1.amazonaws.com/dev/world
functions:
  hello: mischievous-slack-bot-dev-hello
  world: mischievous-slack-bot-dev-world
layers:
  None

Stack Outputs
HelloLambdaFunctionQualifiedArn: arn:aws:lambda:ap-northeast-1:************:function:mischievous-slack-bot-dev-hello:X
WorldLambdaFunctionQualifiedArn: arn:aws:lambda:ap-northeast-1:************:function:mischievous-slack-bot-dev-world:X
ServiceEndpoint: https://foo.execute-api.ap-northeast-1.amazonaws.com/dev
ServerlessDeploymentBucketName: mischievous-slack-bot-de-serverlessdeploymentbuck-xxxxxxxxxxxx

Serverless: Run the "serverless" command to setup monitoring, troubleshooting and testing.
```

### 6. create slack app

select `Bots` (Basic Information > Add features and functionality > Bots).
setup Events API (need Request URI from Event Subscriptions page).
install bot to Workspace.
getting app auth token(Install App > Bot User OAuth Access Token).
