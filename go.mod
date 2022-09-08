module github.com/SimulatedSakura/go-gin-example

go 1.19

require (
	github.com/astaxie/beego v1.12.3
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.8.1
	github.com/go-ini/ini v1.67.0
	github.com/jinzhu/gorm v1.9.16
	github.com/unknwon/com v1.0.1
)

require (
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.11.0 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/goccy/go-json v0.9.11 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.5 // indirect
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	golang.org/x/crypto v0.0.0-20220829220503-c86fa9a7ed90 // indirect
	golang.org/x/net v0.0.0-20220906165146-f3363e06e74c // indirect
	golang.org/x/sys v0.0.0-20220906165534-d0df966e6959 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace (
	github.com/SimulatedSakura/go-gin-example/conf => ./go-application/go-gin-example/pkg/conf
	github.com/SimulatedSakura/go-gin-example/middleware => ./go-application/go-gin-example/middleware
	github.com/SimulatedSakura/go-gin-example/models => ./go-application/go-gin-example/models
	github.com/SimulatedSakura/go-gin-example/pkg/e => ./go-application/go-gin-example/pkg/e
	github.com/SimulatedSakura/go-gin-example/pkg/setting => ./go-application/go-gin-example/pkg/setting
	github.com/SimulatedSakura/go-gin-example/pkg/util => ./go-application/go-gin-example/pkg/util
	github.com/SimulatedSakura/go-gin-example/routers => ./go-application/go-gin-example/routers
)
