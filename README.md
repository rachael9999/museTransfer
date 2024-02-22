#cloud disk

backend: go-zero, xorm,
database: redis, MySQL

goctl api new core

goctl api go -api core.api -dir . -style go_zero

go run core.go -f etc/core-api.yaml
