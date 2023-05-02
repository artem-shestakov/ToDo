# ToDo app
This app is using [Telegram bot](https://github.com/artem-shestakov/ToDo-Telegram) as frontend.

## How to run
1. Set environment variables:

  | Env | Description |
  | --- | --- |
  | SERVER_ADDRESS | Address to listen
  | SERVER_PORT | Port to listen
  | DB_ADDRESS | Database address
  | DB_PORT | Database port
  | DB_USER | Database user
  | DB_PASS | Database password
  | DB_NAME | Database name
  | API_TOKEN |  API token to comunicate between telegram bot and this app
  or configuration file in `YAML` format. Example:
  ```yaml
  # app params
  server:
    address: 0.0.0.0
    port: 8000
  # database connections params
  database:
    address: 127.0.0.1
    port: 5432
    username: postgres
    password: postgres
    db_name: todo
  # secret api token to comunicate between app and telegram bot
  api_token: my_secret_t0ken
  ```
2. Run app
```shell
go run cmd/main.go -config config.yaml
```

## Params
* `-config` - Path to configuration file (default "config.yaml")
* `-h` - Show help text