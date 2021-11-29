package app

type NacosConfig struct {
	Server []struct {
		Scheme      string `yaml:"scheme"`
		IP          string `yaml:"ip"`
		Port        uint64 `yaml:"port"`
		ContextPath string `yaml:"contextPath"`
	} `yaml:"server"`
	Namespace           string `yaml:"namespace"`
	LogDir              string `yaml:"logDir"`
	CacheDir            string `yaml:"cacheDir"`
	ApiServiceName      string `yaml:"apiServiceName"`
	SearchServiceName   string `yaml:"searchServiceName"`
	EsSearchServiceName string `yaml:"esSearchServiceName"`
	RecoServiceName     string `yaml:"recoServiceName"`
	CmsServiceName      string `yaml:"cmsServiceName"`
	OmsServiceName      string `yaml:"omsServiceName"`
	SmsUserServiceName  string `yaml:"smsUserServiceName"`
	NftServiceName      string `yaml:"nftServiceName"`
}
