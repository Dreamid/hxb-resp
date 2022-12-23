package main

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"hxbk_resp/models"
)

func main() {
	r := gin.Default()
	r.GET("/xml", func(c *gin.Context) {
		//time.Sleep(time.Second * 500)
		//c.BindHeader(xml.Header)
		resp := GetDullData()
		//c.BindXML(resp)
		//c.ShouldBindXML(resp)
		//c.Header("header", xml.Header)
		//header := `<?xml version="1.0" encoding="GBK"?>` + "\n"
		//c.Header("header", header)
		c.Header("Content-Type", "text/xml;charset=UTF-8")
		//c.Header("-Type", "text/xml;charset=GBK")
		//c.Header("Content-Type", "text/xml;charset=GBK")
		c.Writer.Write(resp)

	})

	r.GET("/xml-fast", func(c *gin.Context) {
		c.XML(200, GetFastData())
	})

	r.Run(":8080")
}

func GetDullData() []byte {
	data := models.HxDataStruct{
		Result:          true,
		MaxSerialNumber: 1248448,
		NowSerialNumber: 1248448,
		Org: models.OrgType{
			OrgId:   "ee47838c23090d990123090e0f8a0002",
			OrgName: "华夏银行",
			Org: &models.OrgType{
				OrgId:   "0001.02.000000000001",
				OrgName: "总行",
				Staff: []*models.StaffType{
					&models.StaffType{
						StaffId:   "e72033ae795acd2201795af280ac0007",
						StaffName: "张三",
					},
					&models.StaffType{
						StaffId:   "e72033ae795acd2201795af2dcbc001a",
						StaffName: "李四",
					},
					&models.StaffType{
						StaffId:   "e72033ae795acd2201795af356a002d",
						StaffName: "王五",
					},
					&models.StaffType{
						StaffId:   "e72033ae795acd2201795af3738d0040",
						StaffName: "成燕红",
					},
					&models.StaffType{
						StaffId:   "e72033ae795acd2201795af392db0053",
						StaffName: "邱悦",
					},
					&models.StaffType{
						StaffId:   "e72033ae795acd2201795af39e690066",
						StaffName: "邵卓",
					},
					&models.StaffType{
						StaffId:   "e72033ae795acd2201795af3a5380079",
						StaffName: "李新",
					},
					&models.StaffType{
						StaffId:   "e72033ae795acd2201795af3b1370098",
						StaffName: "王海洁",
					},
					&models.StaffType{
						StaffId:   "e72033ae795acd2201795af3c23d00c2",
						StaffName: "王一平",
					},
				},
				Org: &models.OrgType{
					OrgId:   "e720ed6d82854ab70182f7c6de710d0a",
					OrgName: "测试机构",
					Org: &models.OrgType{
						OrgId:   "e720ed6d82854ab70182f7c7a2b10d20",
						OrgName: "测试机构2",
						Staff: []*models.StaffType{
							&models.StaffType{
								StaffId:   "e720ed6d82854ab70182f7cd38730d5d",
								StaffName: "测试人员2",
							},
						},
					},
				},
			},
		},
	}
	//b, err := xml.Marshal(data)
	//if err != nil {
	//	log.Println(err)
	//	return nil
	//}
	b, _ := xml.MarshalIndent(data, "", "	")
	//xml.Marshal(data)
	//xmlHeader := `<?xml version="1.0" encoding="UTF-8"?>` + "\n"
	resp := append([]byte(xml.Header), b...)
	return resp
}

func GetFastData() *models.HxDataStruct {

	data := models.HxDataStruct{
		Result:          true,
		MaxSerialNumber: 1248460,
		NowSerialNumber: 1248460,
		DataIncr: models.DataIncrType{
			ChangeData: []*models.ChangeDataType{
				&models.ChangeDataType{
					StaffChange: &models.StaffChangeType{
						Operation: "增加",
						Staff: models.StaffType{
							StaffId:   "e720ed6d84660a850184e6c260f20e75",
							StaffName: "测试22",
							Account:   "lyzdcessss",
						},
					},
				},
				&models.ChangeDataType{
					StaffChange: &models.StaffChangeType{
						Operation: "修改",
						Staff: models.StaffType{
							StaffId:   "e72033ae795acd2201795af280ac0007",
							StaffName: "胡文亮",
							Account:   "kmhw1",
						},
					},
				},
				&models.ChangeDataType{
					StaffChange: &models.StaffChangeType{
						Operation: "删除",
						Staff: models.StaffType{
							StaffId:   "e72033ae795acd2201795af280ac0007",
							StaffName: "胡文亮",
							Account:   "kmhw1",
						},
					},
				},
				&models.ChangeDataType{
					StaffChange: &models.StaffChangeType{
						Operation:   "调出",
						ParentOrgId: "0001.02.000000000513",
						StaffId:     "0001.01.100000008723",
					},
				},
				&models.ChangeDataType{
					StaffChange: &models.StaffChangeType{
						Operation:   "调入",
						StaffId:     "e72033ae795acd2201795af2dcbc001a",
						ParentOrgId: "ee47838c23090d990123090e0f8a0002",
					},
				},
				&models.ChangeDataType{
					StaffChange: &models.StaffChangeType{
						Operation: "增加上下级关系",
						OrgId:     "e720ed6d82854ab70182f7c7a2b10d20",
						StaffId:   "e720ed6d84660a850184e6c260f20e75",
						Staff: models.StaffType{
							StaffId:   "e720ed6d84660a850184e6c260f20e75",
							StaffName: "司彦",
							Account:   "lyzdcessss",
						},
					},
				},
				&models.ChangeDataType{
					StaffChange: &models.StaffChangeType{
						Operation:   "删除上下级关系",
						ParentOrgId: "e720ed6d82854ab70182f7c7a2b10d20",
						StaffId:     "e720ed6d84660a850184e6c260f20e75",
					},
				},
				&models.ChangeDataType{
					OrgChange: &models.OrgChangeType{
						Operation: "增加",
						Org: models.OrgType{
							OrgId:   "0001.02.100000001494",
							OrgName: "新增海外机构",
						},
					},
				},
				&models.ChangeDataType{
					OrgChange: &models.OrgChangeType{
						Operation: "修改",
						Org: models.OrgType{
							OrgId:   "0001.02.100000001494",
							OrgName: "新增海外机构",
						},
					},
				},
			},
		},
	}

	return &data
}
