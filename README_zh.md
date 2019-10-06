# 基于beego的手机号归属地查询接口实现

前一段时间学习beego，找了一个练手的项目，于是用beego实现了一下，经过测试非常稳定。

> 这可能是全网上能找到的最新最全最好用的中国境内手机号归属地信息库查询接口。

## 测试地址：
http://47.100.39.205:8018/v1.0/api/phoneLocation?phoneNum=18800000000

单 ip 每天访问次数限制为 10000 次。量不够的话可以联系我加就完事了。

返回数据为 json:
- 正确返回
```json
{
  "Data": {
    "PhoneNum": "18800000000",
    "Province": "北京",
    "City": "北京",
    "ZipCode": "100000",
    "AreaZone": "010",
    "CardType": "中国移动"
  },
  "Code": 0,
  "Message": "success"
}
```
- 参数错误
```json
{
  "Data": {
    "PhoneNum": "183519252811",
    "Province": "",
    "City": "",
    "ZipCode": "",
    "AreaZone": "",
    "CardType": ""
  },
  "Code": -2,
  "Message": "Server is busy."
}
```

- 查询次数用尽
```json
{
  "Data": {
    "PhoneNum": "18351925281",
    "Province": "",
    "City": "",
    "ZipCode": "",
    "AreaZone": "",
    "CardType": ""
  },
  "Code": -1,
  "Message": "The number of available times is exhausted."
}
```

## 快速安装

下载对应平台的打包程序phoneLocationOpenApi{{platform}}.tar.gz，解压后直接运行可执行程序即可。

## 测试
在阿里云轻量云上测试了一下（配置：1核 - 2GB内存 - 系统盘40g），请求1w次，并发数100，结果如下：

```
This is ApacheBench, Version 2.3 <$Revision: 1430300 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking *.*.*.* (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        AllRightsReserved
Server Hostname:        47.100.39.205
Server Port:            8018

Document Path:          /v1.0/api/phoneLocation?phoneNum=18351925281
Document Length:        214 bytes

Concurrency Level:      100
Time taken for tests:   17.028 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      3730000 bytes
HTML transferred:       2140000 bytes
Requests per second:    587.27 [#/sec] (mean)
Time per request:       170.280 [ms] (mean)
Time per request:       1.703 [ms] (mean, across all concurrent requests)
Transfer rate:          213.92 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   97 338.6      0    3008
Processing:     0   51 214.1      1    6061
Waiting:        0   51 214.1      1    6061
Total:          1  149 427.0      1    7587

Percentage of the requests served within a certain time (ms)
  50%      1
  66%     13
  75%     36
  80%    201
  90%    417
  95%   1004
  98%   1202
  99%   1775
 100%   7587 (longest request)

```

## 致谢

@xluohome https://github.com/xluohome
@astaxie https://github.com/astaxie