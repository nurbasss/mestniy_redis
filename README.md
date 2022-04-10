# INF 368 midterm project

ID-190103335 Nurbakyt Akhmetshaiykov

Simple Redis clone with golang
Furthermore, it save data into local json, so even after server terminating, you will be able to retrieve data from repo

## How to start

Clone the repo

```bash
git clone https://github.com/nurbasss/mestniy_redis.git
```

```bash
go run .\cmd\api\main.go 
```

## Endpoints
Can not implement route like in task so please use:

- localhost:4200/put?key=keyname&value=valuename
- localhost:4200/get?key=keyname
