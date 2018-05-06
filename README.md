# Circle CI Slack Bot
consumes slack commands and processes them against the Circle CI API

## Build
```shell
docker build . -t sirkirby/circlecibot:latest
```

## Run
```shell
docker run -it --rm --name circlecibot -p 8085:80 -e SLACK_TOKEN=[your-token-here] sirkirby/circlecibot:latest
```