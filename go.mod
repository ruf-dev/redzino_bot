module github.com/ruf-dev/redzino_bot

go 1.24.2

require (
	github.com/Red-Sock/go_tg v0.0.24
	github.com/go-telegram-bot-api/telegram-bot-api/v5 v5.5.1
	github.com/lib/pq v1.10.9
	github.com/pressly/goose/v3 v3.24.3
	github.com/sirupsen/logrus v1.9.3
	go.redsock.ru/rerrors v0.0.3
	go.redsock.ru/toolbox v0.0.11
	go.vervstack.ru/matreshka v1.0.78
	golang.org/x/sync v0.14.0
)

require (
	github.com/mfridman/interpolate v0.0.2 // indirect
	github.com/sethvargo/go-retry v0.3.0 // indirect
	go.redsock.ru/evon v0.0.25 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250519155744-55703ea1f237 // indirect
	google.golang.org/grpc v1.72.2 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/Red-Sock/go_tg => /Users/alexbukov/redsock/go_tg
)