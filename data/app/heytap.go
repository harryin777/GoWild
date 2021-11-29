package app

type HeyTapConfig struct {
	Heytap HeyTapAuth `yaml:"Heytap"`
	Oppo   HeyTapAuth `yaml:"Oppo"`
}

type HeyTapAuth struct {
	AppID           string `yaml:"appId"`
	AppKey          string `yaml:"appKey"`
	AppSecret       string `yaml:"appSecret"`
	AppServerSecret string `yaml:"appServerSecret"`
}
