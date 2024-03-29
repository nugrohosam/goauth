module github.com/nugrohosam/gosampleapi

// +heroku goVersion go1.15
go 1.15

require (
	github.com/cnjack/throttle v0.0.0-20160727064406-525175b56e18
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fatih/structs v1.1.0
	github.com/garyburd/redigo v1.6.2 // indirect
	github.com/getsentry/sentry-go v0.7.0
	github.com/gin-gonic/contrib v0.0.0-20201101042839-6a891bf89f19
	github.com/gin-gonic/gin v1.6.3
	github.com/go-errors/errors v1.0.1
	github.com/go-playground/validator/v10 v10.4.0
	github.com/go-redis/cache/v8 v8.3.0
	github.com/go-redis/redis/v8 v8.4.8
	github.com/golang/protobuf v1.4.3
	github.com/gorilla/sessions v1.2.1
	github.com/jinzhu/copier v0.1.0
	github.com/lib/pq v1.8.0
	github.com/pkg/errors v0.9.1 // indirect
	github.com/romanyx/polluter v1.2.2 // indirect
	github.com/segmentio/kafka-go v0.4.8
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	golang.org/x/crypto v0.0.0-20200709230013-948cd5f35899
	golang.org/x/sys v0.0.0-20201029080932-201ba4db2418 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/genproto v0.0.0-20201030142918-24207fddd1c3 // indirect
	google.golang.org/grpc v1.33.1
	google.golang.org/protobuf v1.25.0
	gopkg.in/boj/redistore.v1 v1.0.0-20160128113310-fc113767cd6b
	gorm.io/driver/mysql v1.0.2
	gorm.io/driver/postgres v1.0.8
	gorm.io/gorm v1.20.12
)
