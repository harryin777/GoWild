package app

type WXConfig struct {
	AllSaintsMusic WXAuth `yaml:"AllSaintsMusic"`
	OppoMusic      WXAuth `yaml:"OppoMusic"`
	HeyTapMusic    WXAuth `yaml:"HeyTapMusic"`
}

type WXAuth struct {
	AppID     string `yaml:"appID"`
	AppSecret string `yaml:"appSecret"`
}
