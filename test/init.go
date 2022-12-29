package test

import (
	"fmt"
	"hxbk_resp/utils"
	"time"
)

func init() {
	t1 := time.Now()
	dataStr := t1.Format("2006-01-02")
	utils.Init(dataStr, 1)
	fmt.Println("------test init------")
}
