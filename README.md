# beego based mobile phone number attribution query interface implementation

- [中文文档](README_zh.md)  

I learned beego some time ago and found a project that I practiced. So I realized it with beego and it was very stable after testing.

> This may be the latest and most useful online mobile phone number attribution information query interface in China.

## Test API
http://47.100.39.205:8018/v1.0/api/phoneLocation?phoneNum=18800000000

Single ip is limited to 10,000 visits per day. If the amount is not enough, you can contact me and you will be finished.

The returned data is json:

- success

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

- Server is busy
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

- exhausted
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

## Start up

Download the packager phoneLocationOpenApi{{platform}}.tar.gz of the corresponding platform, and run the executable program directly after decompressing.

## Benchmark
Tested on Alibaba Cloud lightweight cloud (configuration: 1 core - 2GB memory - system disk 40g), request 1w times, concurrent number 100, the results are as follows:
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

## Thanks

@xluohome https://github.com/xluohome
@astaxie https://github.com/astaxie
