module github.com/micro/examples/blog/web

go 1.13

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/lithammer/shortuuid/v3 v3.0.4
	github.com/micro/examples/blog/posts v0.0.0-20200706113312-8a497d4e1aaa
	github.com/micro/go-micro/v2 v2.9.1-0.20200713161203-07fbb06ed8fd
)
