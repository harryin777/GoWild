package loader

import (
	"GoWild/data/app"
	"fmt"
	"github.com/msbranco/goconfig"
	"os"
	"path/filepath"
	"sync"
)

const (
	RunModeEnv  = "RUN_MODE"
	ConfPathEnv = "SERVER_CONF_PATH"
)

const (
	DebugMode   = "debug"
	ReleaseMode = "release"
	TestMode    = "test"
)

var (
	configPath       string
	runMode          string
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

	var workPath string
	workPath = os.Getenv(ConfPathEnv)
	if len(workPath) == 0 {
		workPath, err = os.Getwd()
		if err != nil {
			panic(err)
		}
	}

	configPath = filepath.Join(workPath, "config", runMode)
	fmt.Println(fmt.Sprintf("current RUN_MODE : %v", runMode))
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

func IsDebugMode() bool {
	return runMode != ReleaseMode
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
