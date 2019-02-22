module github.com/agile6v/squeeze

require (
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.2.0
	github.com/gorilla/websocket v1.4.0
	github.com/spf13/cobra v0.0.3
	github.com/spf13/pflag v1.0.3
	golang.org/x/net v0.0.0-20190213061140-3a22650c66bd
	google.golang.org/grpc v1.18.0
)

replace (
	golang.org/x/net => github.com/golang/net v0.0.0-20190213061140-3a22650c66bd
	google.golang.org/grpc => github.com/grpc/grpc-go v1.18.0
)
