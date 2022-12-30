module github.com/cjreeder/microservice_test

go 1.19

replace (
	github.com/cjreeder/log_service v0.0.0-20221228221831-d5d707c52ec8 => /home/creeder/go/src/github.com/cjreeder/log_service
	github.com/cjreeder/logging-library v0.0.0-20221229220916-c9d186c68110 => /home/creeder/go/src/github.com/cjreeder/logging-library
)

require (
	github.com/cjreeder/log_service v0.0.0-20221228221831-d5d707c52ec8
	github.com/gin-gonic/gin v1.8.2
	github.com/spf13/pflag v1.0.5
)

require (
	github.com/cjreeder/logging-library v0.0.0-20221229220916-c9d186c68110 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.11.1 // indirect
	github.com/goccy/go-json v0.9.11 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/labstack/gommon v0.4.0 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-colorable v0.1.11 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.6 // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
	golang.org/x/crypto v0.0.0-20211215153901-e495a2d5b3d3 // indirect
	golang.org/x/net v0.4.0 // indirect
	golang.org/x/sys v0.3.0 // indirect
	golang.org/x/text v0.5.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
