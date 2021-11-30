module client_server

go 1.13

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/gin-gonic/gin v1.7.7
	github.com/golang/protobuf v1.5.2
	github.com/jinzhu/now v1.1.3 // indirect
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/consul/v2 v2.9.1
	github.com/spf13/cast v1.3.0
	github.com/spf13/viper v1.6.3
	golang.org/x/sys v0.0.0-20211124211545-fe61309f8881 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/protobuf v1.27.1
	gopkg.in/ini.v1 v1.65.0 // indirect
	gorm.io/driver/mysql v1.2.0
	gorm.io/gorm v1.22.3
)
