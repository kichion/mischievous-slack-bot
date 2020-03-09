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

### 5. events Subscription endpoint create

### 6. deploy to AWS

rewrite `serverless.yml` (profile use ME)

```yml
provider:
  name: aws
  runtime: go1.x
  #   add start
  region: ap-northeast-1
  profile: me
  #   add end
....
functions:
  events:
    handler: bin/events
    events:
      #   add start
      - http:
          path: events
          method: get
      - http:
          path: events
          method: post
      #   add end
```

command execute

```shell
$ make deploy
rm -rf ./bin ./vendor Gopkg.lock
chmod u+x gomod.sh
./gomod.sh
export GO111MODULE=on
....

Serverless: Run the "serverless" command to setup monitoring, troubleshooting and testing.
```

### 8. create slack app

select `Bots` (Basic Information > Add features and functionality > Bots).
setup Events API (need Request URI from Event Subscriptions page).
install bot to Workspace.
getting app auth token(Install App > Bot User OAuth Access Token).

### 9. Let's enjoy!!

## Arrangement

arrangement `mischievous-slack-bot.json` in `S3`

`mischievous-slack-bot.json` is secret of google project service_account
