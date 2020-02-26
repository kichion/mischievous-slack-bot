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

### 5. create slack app

select `Bots` (Basic Information > Add features and functionality > Bots).
setup Events API.
install bot to Workspace.
getting app auth token(Install App > Bot User OAuth Access Token).
