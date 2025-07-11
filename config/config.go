package config

var (
	//  初始化请求授权服务器的地址、账号、密码
	Url      = "https://login.microsoftonline.com/1b408d81-d41c-4d6c-9345-ce89c6c9ac39/oauth2/v2.0/authorize?client_id=a1b9057a-9fde-4119-b25c-ffd2e59e4f9a&response_type=code&redirect_uri=https%3A%2F%2Flogin.microsoftonline.com%2Fcommon%2Foauth2%2Fnativeclient&response_mode=query&scope=offline_access%20Files.ReadWrite&state=12345"
	Username = "jwtan@alauda.io"
	Password = "@200201WeI0711"
	// token请求获取的URL
	TargetURL = "https://login.microsoftonline.com/1b408d81-d41c-4d6c-9345-ce89c6c9ac39/oauth2/v2.0/token"
	// Basic auth的账号密码
	AuthUsername string = "nike1234"
	AuthPassword string = "farewell@33550336"
)

// 定义反序列化接收json
type Microbody struct {
	Token_type     string `json:"token_type"`
	Scope          string `json:"scope"`
	Expires_in     int    `json:"expires_in"`
	Ext_expires_in int    `json:"ext_expires_in"`
	Access_token   string `json:"access_token"`
	Refresh_token  string `json:"refresh_token"`
}
