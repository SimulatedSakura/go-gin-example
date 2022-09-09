module github.com/SimulatedSakura/go-gin-example

go 1.19

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/astaxie/beego v1.12.3
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fvbock/endless v0.0.0-20170109170031-447134032cb6
	github.com/gin-gonic/gin v1.8.1
	github.com/go-ini/ini v1.67.0
	github.com/jinzhu/gorm v1.9.16
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.8.5
	github.com/unknwon/com v1.0.1
)

require (
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.20.0 // indirect
	github.com/go-openapi/spec v0.20.7 // indirect
	github.com/go-openapi/swag v0.22.3 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.11.0 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/goccy/go-json v0.9.11 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.5 // indirect
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	golang.org/x/crypto v0.0.0-20220829220503-c86fa9a7ed90 // indirect
	golang.org/x/net v0.0.0-20220907135653-1e95f45603a7 // indirect
	golang.org/x/sys v0.0.0-20220908164124-27713097b956 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.12 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/SimulatedSakura/go-gin-example/conf => ./go-application/go-gin-example/pkg/conf
	github.com/SimulatedSakura/go-gin-example/docs => ./go-application/go-gin-example/docs
	github.com/SimulatedSakura/go-gin-example/middleware => ./go-application/go-gin-example/middleware
	github.com/SimulatedSakura/go-gin-example/models => ./go-application/go-gin-example/models
	github.com/SimulatedSakura/go-gin-example/pkg/e => ./go-application/go-gin-example/pkg/e
	github.com/SimulatedSakura/go-gin-example/pkg/logging => ./go-application/go-gin-example/pkg/logging
	github.com/SimulatedSakura/go-gin-example/pkg/setting => ./go-application/go-gin-example/pkg/setting
	github.com/SimulatedSakura/go-gin-example/pkg/util => ./go-application/go-gin-example/pkg/util
	github.com/SimulatedSakura/go-gin-example/routers => ./go-application/go-gin-example/routers
)
