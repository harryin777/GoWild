package ip

import (
	"GoWild/helper/logger"
	"GoWild/utils"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"github.com/thinkeridea/go-extend/exnet"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

var (
	once sync.Once
	loc  Location
)

type section struct {
	min uint
	max uint
}

var (
	apnicURL = "http://ftp.apnic.net/apnic/stats/apnic/delegated-apnic-latest"
)

type Location interface {
	ReadLocal()
	Dispatcher()
	GetLocation(ip string) bool
}

type location struct {
	s  []section
	mu sync.Mutex
}

func LocationInstances() Location {
	once.Do(func() {
		loc = &location{}
	})
	return loc
}

func (l *location) update() {
	defer func() {
		if r := recover(); r != nil {
			msg := fmt.Sprintf("panic: %v", r)
			_ = msg
		}
	}()
	resp, err := http.Get(apnicURL)
	if err != nil {
		log.Default().Println(err)
		return
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	lines := strings.Split(string(data), "\n")
	l.mu.Lock()
	defer l.mu.Unlock()
	for _, v := range lines {
		val := strings.Split(v, "|")
		if len(val) == 7 && val[1] == "CN" && val[2] == "ipv4" {
			min, _ := exnet.IPString2Long(val[3])
			max, _ := exnet.IPString2Long(val[3])
			l.s = append(l.s, section{min: min, max: max + cast.ToUint(val[4])})
		}
	}
}

func (l *location) Dispatcher() {
	timeTickerChan := time.Tick(24 * time.Hour)
	for {
		l.update()
		<-timeTickerChan
	}
}

func (l *location) GetLocation(ip string) bool {
	val, _ := exnet.IPString2Long(ip)
	for _, v := range l.s {
		if val >= v.min && val <= v.max {
			return true
		}
	}
	return false
}

func (l *location) ReadLocal() {
	data, err := ioutil.ReadFile("./delegated-apnic-latest")
	if err != nil {
		defer func() {
			logger.ErrLogger().WithFields(logrus.Fields{
				"service":   utils.GetServiceName(),
				"func":      utils.GetFuncName(),
				"backtrace": utils.GetStack(),
			}).Error(errors.New("delegated-apnic-latest file not found"))
		}()
		return
	}
	lines := strings.Split(string(data), "\n")
	for _, v := range lines {
		val := strings.Split(v, "|")
		if len(val) == 7 && val[1] == "CN" && val[2] == "ipv4" {
			min, _ := exnet.IPString2Long(val[3])
			max, _ := exnet.IPString2Long(val[3])
			l.s = append(l.s, section{min: min, max: max + cast.ToUint(val[4])})
		}
	}
}
