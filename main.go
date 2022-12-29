package main

import (
	"github.com/gin-gonic/gin"
	"hxbk_resp/logic"
	"hxbk_resp/utils"
	"time"
)

func init() {
	t1 := time.Now()
	dataStr := t1.Format("2006-01-02")
	utils.Init(dataStr, 1)
}

func main() {

	r := gin.Default()
	r.GET("/xml", func(c *gin.Context) {
		//计数器清零
		utils.UserCount = 0
		utils.OrgCount = 0

		resp := logic.GetFullData3()
		//resp := logic.GetFullData2()
		//c.BindXML(resp)
		//c.ShouldBindXML(resp)
		//c.Header("header", xml.Header)
		//header := `<?xml version="1.0" encoding="GBK"?>` + "\n"
		//c.Header("header", header)
		c.Header("Content-Type", "text/xml;charset=GBK")
		c.Writer.Write(resp)

	})

	r.GET("/xml-fast", func(c *gin.Context) {
		c.XML(200, logic.GetFastData())
	})

	r.Run(":8080")
}
