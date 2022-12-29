package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var mu sync.Mutex
var UserCount int64
var OrgCount int64

func GeneraterIntIdStr() string {
	//time.Sleep(time.Second)
	//r := rand.New(rand.NewSource(time.Now().Unix()))
	rand.Seed(time.Now().Unix())
	return strconv.Itoa(rand.Int())

}

//随机生成账户名称
func GetRandUserAccount(length int) string {
	mu.Lock()
	defer mu.Unlock()
	letters := "1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GetRandUserName(level int) string {
	atomic.AddInt64(&UserCount, 1)
	return fmt.Sprintf("user-%v-%v", level, UserCount)
}

func GetRandOrgName(level int) string {
	atomic.AddInt64(&OrgCount, 1)
	return fmt.Sprintf("org-%v-%v", level, OrgCount)
}
