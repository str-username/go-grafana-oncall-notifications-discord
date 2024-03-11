### Info
A simple script that sends notifications about the current engineer on duty from grafana oncall to discord using webhook


### Run
```shell
echo '{"grafana_user_name": "<@discord_user_gid>"}' > example.json
docker build -f build/Dockerfile .
docker run -it \
--env SCHEDULE_URL="http(s)://grafana-api-url" \
--env USERS_URL="http(s)://discord_webhook_url" \
--env TOKEN="grafana_api_token" \
--env USERS_FILE="etc/users.json"
```
