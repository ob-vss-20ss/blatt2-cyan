module github.com/ob-vss-20ss/blatt2-cyan

go 1.13

require (
	github.com/golang/protobuf v1.4.2
	github.com/google/uuid v1.1.1
	github.com/micro/go-micro/v2 v2.9.0
	github.com/micro/go-plugins/broker/nats/v2 v2.5.0
	github.com/micro/go-plugins/logger/zerolog/v2 v2.8.0
	github.com/micro/go-plugins/registry/etcdv3/v2 v2.8.0
	github.com/micro/go-plugins/store/redis/v2 v2.8.0
	github.com/micro/micro/v2 v2.9.0 // indirect
	github.com/rs/zerolog v1.19.0
	github.com/vesose/example-micro v0.0.0-20200609090234-d46ee1255141 // indirect
	golang.org/x/net v0.0.0-20200602114024-627f9648deb9 // indirect
	golang.org/x/sys v0.0.0-20200610111108-226ff32320da // indirect
	google.golang.org/genproto v0.0.0-20200612171551-7676ae05be11 // indirect
	google.golang.org/grpc v1.29.1 // indirect
	google.golang.org/protobuf v1.24.0
)

replace (
	github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)
