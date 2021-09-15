# gitlab-release-helper

上传 文件 到 s3

## Usage

为 server 设置环境变量，并启动

```yaml
Server__S3_AccessID: "AKidZZZZZZZZZZ"
Server__S3_AccessKey: "AKeyXXXXXXXXXX"
Server__S3_Bucket: "bucketname"
Server__S3_Endpoint: "s3.exmaple.com"
Server__S3_ObjectPrefixPath: ""
Server__S3_Region: ""
Server__S3_SSL: false
Server__Server_Addr: ""
Server__Server_Port: 8088
```

## 

## todo

+ [x] gitlab 的原因， `CI_JOB_TOKEN` 只能创建 `Release` 且可以包含 `Link`。 但是无法单独创建 `Link`。
+ [x] 由于 `go-gitlab` 库的原因，无法上传 `external=true` 的外部链接。
+ [x] 使用 `go-jarvis`，虽然方便的从环境变量获取参数， 但是 s3 的 **AK/SK** 都变成了明文不安全。
+ [x] 应该是 `C-S` 模式
    + c 向 s 发送 PUT 请求， s 返回 `PresignPutURL` 和 `Permanent Download URL`
    + c 向 s 发送 GET 请求， s 返回 `PresignGetURL`
