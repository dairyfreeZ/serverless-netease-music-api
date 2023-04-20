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

## 参数说明
### signin
```
{
  "username": "foo@bar.com", // [required] 登陆邮箱
  "password": "fc5e038d38a57032085441e7fe7010b0", // [required] 登陆密码，md5 encoded
  "state": {
    "location": "s3://bkt", // [optional] 存取state(目前即cookies)的s3 bucket的URI，如未提供或invalid或expired，将根据提供的账号密码重新login
    "region": "us-west-2" // [optional] s3 bucket的region
  },
  "ip": "0.0.0.1" // [optional] 用于header的ip addr
}
```
### visit
```
{
  "username": "foo@bar.com", // [required] 登陆邮箱
  "password": "fc5e038d38a57032085441e7fe7010b0", // [required] 登陆密码，md5 encoded
  "state": {
    "location": "s3://bkt", // [optional] 存取state(cookies)的s3位置，如未提供或invalid或expired，将根据提供的账号密码重新login
    "region": "us-west-2" // [optional] s3 bucket的region
  },
  "ip": "0.0.0.1", // [optional] 用于header的ip addr
  "path": "#" // [optional] 想要访问的地址，比如"foo"则返回"music.163.com/foo"的内容，不提供即访问主页
}
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

In AWS Console, create a Test event with the following payload to trigger the Lambda

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

In FC Console, create a Test event with the following payload to trigger the FC

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

## 历史
APR-22 2023
- 支持aws s3存取state,目前只包括cookies
- 支持visit api: 在login的状态下访问网页，范围取得的内容
- 支持自定义header ip值
