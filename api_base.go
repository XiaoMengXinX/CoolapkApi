package coolapk

import token "github.com/XiaoMengXinX/FuckCoolapkTokenV2"
import "context"

const defaultAPIEndpoint = "http://api.coolapk.com/v6"
const defaultUserAgent = `Dalvik/2.1.0 (Linux; U; Android 11) +CoolMarket/12.1-2203161-universal`

type APIResp interface {
	Deserialize(resp string)
}

type APIClient interface {
	Request(c *Coolapk, result APIResp, method, path string, ctx context.Context, params map[string]interface{}) error
}

type Coolapk struct {
	APIEndpoint string
	DeviceID    string
	UserAgent   string
	Cookie      string
	Client      APIClient
}

func (c *Coolapk) init() {
	c.APIEndpoint = defaultAPIEndpoint
	c.UserAgent = defaultUserAgent
	c.DeviceID, _ = token.GetToken()
	c.Client = &CoolapkClient{}
}

func New() Coolapk {
	var c Coolapk
	c.init()
	return c
}
