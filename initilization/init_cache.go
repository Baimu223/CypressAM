package initilization

import (
	"github.com/patrickmn/go-cache"
	"time"
	"xkginweb/global"
)

func InitCache() {
	c := cache.New(5*time.Minute, 24*60*time.Minute)
	global.Cache = c
}
