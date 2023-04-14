module github.com/dairyfreeZ/serverless-netease-music-api/configuration

go 1.19

require (
	github.com/aliyun/fc-runtime-go-sdk v0.2.7
	github.com/aws/aws-lambda-go v1.39.1
	github.com/dairyfreeZ/serverless-netease-music-api/sdk v0.0.0-00010101000000-000000000000
)

require (
	github.com/sirupsen/logrus v1.9.0 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
)

replace github.com/dairyfreeZ/serverless-netease-music-api/sdk => ../sdk
