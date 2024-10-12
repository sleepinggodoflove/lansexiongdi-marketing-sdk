## 功能介绍

1. 接口 SDK。详见 [接口介绍](services)。
2. HTTP 客户端，支持请求签名和应答验签。如果 SDK 未支持你需要的接口，请用此客户端发起请求。
3. 回调通知处理库，支持回调通知的验签。详见 [回调通知验签](#回调通知的验签)。
4. 密钥生成下载、[敏感信息加解密](#敏感信息加解密) 等辅助能力。

## 快速开始

### 安装

#### 1、使用 Go Modules 管理你的项目

如果你的项目还不是使用 Go Modules 做依赖管理，在项目根目录下执行：

```shell
go mod init
```

#### 2、无需 clone 仓库中的代码，直接在项目目录中执行：
```shell
go get -u github.com/sleepinggodoflove/lansexiongdi-marketing-sdk
```
来添加依赖，完成 `go.mod` 修改与 SDK 下载。


## 示例

### 以 [获取key码下单](https://tvd8jq9lqkp.feishu.cn/wiki/PVq3wtanPicDu0kyfpLc0McMnAc?from=from_copylink) 为例

```go
package main

import (
	"context"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/api/v1/key"
	core2 "github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/core"
	"log"
)

func main() {
	core, err := core2.NewCore(&core2.Config{
		AppID:      "123",
		PrivateKey: "私钥",
		PublicKey:  "验签公钥",
		Key:        "业务参数加密key",
		BaseURL:    "请求地址",
	})
	if err != nil {
		log.Fatalf("new core err:%s", err)
	}
	a := &key.Key{core}
	r, err := a.Order(context.Background(), &key.OrderRequest{
		OutBizNo:   "outBizNo",
		ActivityNo: "activityNo",
		Number:     1,
	})
	if err != nil {
		log.Fatalf("key get err:%s", err)
	}
	log.Printf(r)
}
```