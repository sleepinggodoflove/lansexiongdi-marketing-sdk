## 功能介绍

1. 接口 SDK。详见 [接口介绍](services)。
2. HTTP 客户端，支持请求签名和应答验签。如果 SDK 未支持你需要的接口，请用此客户端发起请求。
3. 回调通知处理库，支持回调通知的验签。详见 [回调通知验签](#回调通知的验签)。
4. 密钥生成下载、[敏感信息加解密](#敏感信息加解密) 等辅助能力。


#### 名词解释
+ **商户 API 公钥**，是用来证实商户身份的
+ **商户 API 私钥**。是用来证实商户身份的
+ **商户 API 密钥**。是商户用来加密请求参数的密钥，为加强数据安全，使用的对称加密密钥。
> :warning: 不要把私钥文件暴露在公共场合，如上传到 Github，写在客户端代码等。

## 快速开始

### 安装

#### 1、使用 Go Modules 管理你的项目

如果你的项目还不是使用 Go Modules 做依赖管理，在项目根目录下执行：

```shell
go mod init
```

#### 2、在项目目录中执行：
```shell
go get -u github.com/sleepinggodoflove/lansexiongdi-marketing-sdk
```
来添加依赖，完成 `go.mod` 修改与 SDK 下载。


## 示例

#### [获取key码](https://tvd8jq9lqkp.feishu.cn/wiki/PVq3wtanPicDu0kyfpLc0McMnAc?from=from_copylink)

```go
package main

import (
	"context"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/api/v1/key"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/core"
	"log"
)

func main() {
	c, err := core.NewCore(&core.Config{
		AppID:      "123",
		PrivateKey: "私钥",
		PublicKey:  "验签公钥",
		Key:        "业务参数密钥key",
		SignType:   "签名类型",
		BaseURL:    "请求地址:https://api.lansexiongdi.com",
	})
	if err != nil {
		log.Fatalf("new core err:%v", err)
	}
	a := &key.Key{c}
	_,r, err := a.Order(context.Background(), &key.OrderRequest{
		OutBizNo:   "123456",
		ActivityNo: "123456",
		Number:     1,
	})
	if err != nil {
		log.Fatalf("key get err:%v", err)
	}
	log.Printf(r)
}
```

#### [查询key码](https://tvd8jq9lqkp.feishu.cn/wiki/GvRswEDyfiXGUUkkDCYc8xg4nVX?from=from_copylink)
```go
package main

import (
	"context"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/api/v1/key"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/core"
	"log"
)

func main() {
	c, err := core.NewCore(&core.Config{
		AppID:      "123",
		PrivateKey: "私钥",
		PublicKey:  "验签公钥",
		Key:        "业务参数密钥key",
		SignType:   "签名类型",
		BaseURL:    "请求地址",
	})
	if err != nil {
		log.Fatalf("new core err:%v", err)
	}
	a := &key.Key{c}
	_,r, err := a.Query(context.Background(), &key.QueryRequest{
		OutBizNo:   "123456",
		trade_no:   "123456",
	})
	if err != nil {
		log.Fatalf("key query err:%v", err)
	}
	log.Printf(r)
}
```

#### [作废key码](https://tvd8jq9lqkp.feishu.cn/wiki/R9NMw96eIiXLiRkOi7icANkynbb?from=from_copylink)
```go
package main

import (
	"context"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/api/v1/key"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/core"
	"log"
)

func main() {
	c, err := core.NewCore(&core.Config{
		AppID:      "123",
		PrivateKey: "私钥",
		PublicKey:  "验签公钥",
		Key:        "业务参数密钥key",
		SignType:   "签名类型",
		BaseURL:    "请求地址",
	})
	if err != nil {
		log.Fatalf("new core err:%v", err)
	}
	a := &key.Key{c}
	_,r, err := a.Discard(context.Background(), &key.DiscardRequest{
		OutBizNo:   "123456",
		trade_no:   "123456",
	})
	if err != nil {
		log.Fatalf("key query err:%v", err)
	}
	log.Printf(r)
}
```

#### [回调通知](https://alidocs.dingtalk.com/i/nodes/N7dx2rn0Jb6A1wvLixErNlLkJMGjLRb3?utm_scene=team_space)
```go
package main

import (
	"context"
	"net/http"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/api/v1/key"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/core"
	"log"
)

func main() {
	c, err := core.NewCore(&core.Config{
		AppID:      "123",
		PrivateKey: "私钥",
		PublicKey:  "验签公钥",
		Key:        "业务参数密钥key",
		SignType:   "签名类型",
		BaseURL:    "请求地址",
	})
	if err != nil {
		log.Fatalf("new core err:%v", err)
	}
	a := &key.Key{c}
	
	req := &http.Request{
		Header: nil, // 请求头
		Body:   nil, // 请求体
	}
	r, err := a.CallBack(context.Background(), req)
	if err != nil {
		log.Fatalf("key callBack err:%v", err)
	}
	if err != nil {
		log.Fatalf("key notify err:%v", err)
	}
	log.Printf(r)
}
```

#### [其它调用](https://alidocs.dingtalk.com/i/nodes/N7dx2rn0Jb6A1wvLixErNlLkJMGjLRb3?utm_scene=team_space)

```go
package main

import (
	"context"
	"encoding/json"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/api/v1/anyapi"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/core"
	"log"
)

func main() {
	c, err := core.NewCore(&core.Config{
		AppID:      "appid",
		PrivateKey: "私钥",
		PublicKey:  "验签公钥",
		Key:        "业务参数密钥key",
		SignType:   "签名类型",
		BaseURL:    "请求地址",
	})
	if err != nil {
		log.Fatalf("new core err:%v", err)
	}

	bizContent := struct {
		Source       string `json:"source"`         // 来源
		AppId        string `json:"app_id"`         // 应用Id
		MchPublicKey string `json:"mch_public_key"` // 客户公钥
		NotifyUrl    string `json:"notify_url"`     // 事件通知地址,可为空
	}{
		Source:       "来源",
		AppId:        "123",
		MchPublicKey: "123",
		NotifyUrl:    "https://xx.com/xx",
	}

	a := &anyapi.AnyApi{c}

	method := "/openapi/v1/xxx"

	_, r, err := a.AnyApi(context.Background(), method, bizContent)
	if err != nil {
		log.Fatalf("call err:%v", err)
	}

	if !r.IsSuccess() {
		log.Fatalf("err:%s", r.Message)
	}

	var bizDataContent = struct {
		Ciphertext string `json:"ciphertext,omitempty"`
	}{}

	_ = json.Unmarshal(r.Data, &bizDataContent)

	bizJsonContent, _ := c.CryptographySuite.Cipher.Decode(bizDataContent.Ciphertext)

	log.Printf("bizJsonContent=%s", bizJsonContent)
}

```