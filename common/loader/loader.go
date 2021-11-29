package loader

import (
	"GoWild/data/app"
	"GoWild/utils"
	"encoding/json"
	"fmt"
	"github.com/msbranco/goconfig"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

const (
	RunModeEnv    = "RUN_MODE"
	RunEnv        = "RUN_ENV"
	ConfPathEnv   = "SERVER_CONF_PATH"
	ServerChannel = "SERVER_CHANNEL"
)

const (
	DebugMode   = "debug"
	ReleaseMode = "release"
	TestMode    = "test"
)

const (
	AWSProductEnv = "aws_product_env"
	AWSDevelopEnv = "aws_develop_env"
)

var (
	configPath       string
	runMode          string
	channel          string
	configFile       *goconfig.ConfigFile
	appConfigMap     map[string]*goconfig.ConfigFile
	nacosConfig      *app.NacosConfig
	wxConfig         *app.WXConfig
	wxConfigOnce     sync.Once
	heyTapConfig     *app.HeyTapConfig
	heyTapConfigOnce sync.Once

	actor    *app.IDs
	songlist *app.IDs

	curEnv string
)

func init() {
	var err error
	appConfigMap = make(map[string]*goconfig.ConfigFile)

	hasRunMode := false
	if runMode, hasRunMode = os.LookupEnv(RunModeEnv); !hasRunMode {
		fmt.Println(fmt.Sprintf("Can not find environment RUN_MODE. default: %v", TestMode))
		runMode = TestMode
	}

	channel = utils.ConvertChannelId(os.Getenv(ServerChannel))

	var workPath string
	workPath = os.Getenv(ConfPathEnv)
	if len(workPath) == 0 {
		workPath, err = os.Getwd()
		if err != nil {
			panic(err)
		}
	}

	configPath = filepath.Join(workPath, "conf", runMode)

	if curEnv, hasRunMode = os.LookupEnv(RunEnv); !hasRunMode {
		fmt.Println(fmt.Sprintf("Can not find environment RunEnv. default: %v", TestMode))
	}
}

func LoadJsonConfig(config interface{}, filename string) {
	var err error
	var decoder *json.Decoder

	file := OpenFile(filename)
	defer func() {
		_ = file.Close()
	}()

	decoder = json.NewDecoder(file)
	if err = decoder.Decode(config); err != nil {
		msg := fmt.Sprintf("Decode json fail for config file at %s. Error: %v", filename, err)
		panic(msg)
	}
}

func LoadConfigFile(fileName string) *goconfig.ConfigFile {
	if configFile, exist := appConfigMap[fileName]; exist {
		return configFile
	} else {
		var err error
		fullPath := filepath.Join(configPath, fileName)
		configFile, err = goconfig.ReadConfigFile(fullPath)
		appConfigMap[fileName] = configFile
		if err != nil {
			panic(err)
		}
		return configFile
	}
}

func LoadAppConfig() *goconfig.ConfigFile {
	if configFile == nil {
		var err error
		fullPath := filepath.Join(configPath, "app.cfg")
		configFile, err = goconfig.ReadConfigFile(fullPath)
		if err != nil {
			panic(err)
		}
	}
	return configFile
}

func LoadNacosConfig() *app.NacosConfig {
	if nacosConfig == nil {
		var err error
		file := OpenFile("nacos.yaml")
		defer func() {
			_ = file.Close()
		}()
		data, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}
		if err := yaml.Unmarshal(data, &nacosConfig); err != nil {
			panic(err)
		}
	}
	return nacosConfig
}

func LoadWXConfig() *app.WXConfig {
	wxConfigOnce.Do(func() {
		var err error
		file := OpenFile("wx.yaml")
		defer func() {
			_ = file.Close()
		}()
		data, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}
		if err := yaml.Unmarshal(data, &wxConfig); err != nil {
			panic(err)
		}
	})
	return wxConfig
}

func LoadHeyTapConfig() *app.HeyTapConfig {
	heyTapConfigOnce.Do(func() {
		var err error
		file := OpenFile("heytap.yaml")
		defer func() {
			_ = file.Close()
		}()
		data, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}
		if err := yaml.Unmarshal(data, &heyTapConfig); err != nil {
			panic(err)
		}
	})
	return heyTapConfig
}

func LoadRabbitMQConfig(config interface{}, filename string) {
	var err error
	file := OpenFile(filename)
	defer func() {
		_ = file.Close()
	}()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(data, config); err != nil {
		panic(err)
	}
}

func LoadKafkaConfig(config interface{}, filename string) {
	var err error
	var decoder *json.Decoder

	file := OpenFile(filename)
	defer func() {
		_ = file.Close()
	}()

	decoder = json.NewDecoder(file)
	if err = decoder.Decode(config); err != nil {
		msg := fmt.Sprintf("Decode json fail for config file at %s. Error: %v", filename, err)
		panic(msg)
	}
}

func LoadRocketConfig(config interface{}, filename string) {
	var err error
	var decoder *json.Decoder

	file := OpenFile(filename)
	defer func() {
		_ = file.Close()
	}()

	decoder = json.NewDecoder(file)
	if err = decoder.Decode(config); err != nil {
		msg := fmt.Sprintf("Decode json fail for config file at %s. Error: %v", filename, err)
		panic(msg)
	}
}

func IsDebugMode() bool {
	return runMode != ReleaseMode
}

func IsAWSProduct() bool {
	return curEnv == AWSProductEnv
}

func Channel() string {
	return channel
}

func LoadJsonFile(filename string) (cfg string) {
	file := OpenFile(filename)
	defer func() {
		_ = file.Close()
	}()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		msg := fmt.Sprintf("Read config to string error. file at %s. Error: %v", filename, err)
		panic(msg)
	}

	cfg = string(content)

	if IsDebugMode() {
		println(fmt.Sprintf("Load config at %s. Config content: %s", filename, cfg))
	}
	return cfg
}

func GetFullPath(filename string) string {
	return filepath.Join(configPath, filename)
}

func OpenFile(filename string) *os.File {
	fullPath := filepath.Join(configPath, filename)

	var file *os.File
	var err error

	if file, err = os.Open(fullPath); err != nil {
		msg := fmt.Sprintf("Can not load config at %s. Error: %v", fullPath, err)
		panic(msg)
	}

	return file
}

func GetActorYaml() *app.IDs {
	if actor == nil {
		var err error
		file := OpenFile("actor.yaml")
		defer func() {
			_ = file.Close()
		}()
		data, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}
		if err := yaml.Unmarshal(data, &actor); err != nil {
			panic(err)
		}
	}
	return actor
}

func GetSonglistIDYaml() *app.IDs {
	if songlist == nil {
		var err error
		file := OpenFile("songlist.yaml")
		defer func() {
			_ = file.Close()
		}()
		data, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}
		if err := yaml.Unmarshal(data, &songlist); err != nil {
			panic(err)
		}
	}
	return songlist
}
