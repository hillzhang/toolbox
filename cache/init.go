package cache

import (
	"time"
)

func InitCache(host,pwd string, idle, max,db int, toc, tor, tow, defaultExpiration time.Duration) {
	Instance = NewRedisCache(host,pwd, idle, max,db, toc, tor, tow, defaultExpiration)
}
