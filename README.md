# 无服务器网易云APIs

## 一些参考
[Binaryify/NeteaseCloudMusicApi](https://github.com/Binaryify/NeteaseCloudMusicApi)

[darknessomi/musicbox](https://github.com/darknessomi/musicbox)

## 测试过的环境
```
go 1.19

# for deployment only
Terraform v1.2.3
AWS CLI 2.7.3
```

## 使用说明
### 获取代码
```
git clone https://github.com/dairyfreeZ/serverless-netease-music-api
cd serverless-netease-music-api
```

### 部署到AWS Lambda
Build Code
```
cd configuration/lambda/entry
./release.sh
```

[Configure](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html) your AWS CLI if not already
```
aws configure
```

Deploy the binary to Lambda
```
cd ../terraform
terraform init
terraform apply -var="region=us-west-2" -auto-approve
```

In AWS Console, crate a Test event with the following payload to trigger the Lambda

```
{
  "username": "foo@bar.com",
  "password": "fc5e038d38a57032085441e7fe7010b0"
}
```
`username`: 登陆邮箱

`password`: md5-encoded密码，macOS可通过如下command获得
```
echo -n "your_human_readable_password" | md5
```

### 部署到阿里云函数
Build Code
```
cd configuration/fc/entry
./release.sh
```

安装并配置阿里云[CLI](https://help.aliyun.com/product/29991.html)

Deploy the binary to FC
```
cd ../terraform
terraform init

export ALICLOUD_ACCESS_KEY="anaccesskey"
export ALICLOUD_SECRET_KEY="asecretkey"
export ALICLOUD_REGION="cn-beijing"
terraform apply -auto-approve
```

In FC Console, crate a Test event with the following payload to trigger the Lambda

```
{
  "username": "foo@bar.com",
  "password": "fc5e038d38a57032085441e7fe7010b0"
}
```
`username`: 登陆邮箱

`password`: md5-encoded密码，macOS可通过如下command获得
```
echo -n "your_human_readable_password" | md5
```

## Integration Instructions  
- 用其它AWS services定时trigger lambda
- 可与API Gateway结合，用HTTP request trigger Lambda

## 支持APIs
- 签到

## 想做的 (!= todo)
-[] Save/Read Cookie and other artifacts to/from a cloud storage

-[] Read login info from a more secured place than the plaintext payload
