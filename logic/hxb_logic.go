package logic

import (
	"encoding/xml"
	"fmt"
	"hxbk_resp/models"
	"hxbk_resp/utils"
	"os"
)

const (
	ROOT = iota
	LEVEL1
	LEVEL2
	LEVEL3
	LEVEL4
	LEVEL5
)

const (
	LEVEL5_ORG_NUM = (iota + 1) * 100
	LEVEL4_ORG_NUM
	LEVEL3_ORG_NUM
	LEVEL2_ORG_NUM
	LEVEL1_ORG_NUM
	ROOT_ORG_NUM
)

var StaffNums = 2
var OrgLevel = 5

func GetFullData() []byte {
	data := models.HxDataStruct{
		Result:          true,
		MaxSerialNumber: 1248448,
		NowSerialNumber: 1248448,
		Org: models.OrgType{
			OrgId:   "ee47838c23090d990123090e0f8a0002",
			OrgName: "华夏银行",
			//Org: &models.OrgType{
			//	OrgId:   "0001.02.000000000001",
			//	OrgName: "总行",
			//	Staff: []*models.StaffType{
			//		&models.StaffType{
			//			StaffId:   "e72033ae795acd2201795af280ac0007",
			//			StaffName: "张三",
			//		},
			//		&models.StaffType{
			//			StaffId:   "e72033ae795acd2201795af2dcbc001a",
			//			StaffName: "李四",
			//		},
			//		&models.StaffType{
			//			StaffId:   "e72033ae795acd2201795af356a002d",
			//			StaffName: "王五",
			//		},
			//		&models.StaffType{
			//			StaffId:   "e72033ae795acd2201795af3738d0040",
			//			StaffName: "成燕红",
			//		},
			//		&models.StaffType{
			//			StaffId:   "e72033ae795acd2201795af392db0053",
			//			StaffName: "邱悦",
			//		},
			//		&models.StaffType{
			//			StaffId:   "e72033ae795acd2201795af39e690066",
			//			StaffName: "邵卓",
			//		},
			//		&models.StaffType{
			//			StaffId:   "e72033ae795acd2201795af3a5380079",
			//			StaffName: "李新",
			//		},
			//		&models.StaffType{
			//			StaffId:   "e72033ae795acd2201795af3b1370098",
			//			StaffName: "王海洁",
			//		},
			//		&models.StaffType{
			//			StaffId:   "e72033ae795acd2201795af3c23d00c2",
			//			StaffName: "王一平",
			//		},
			//	},
			//	Org: &models.OrgType{
			//		OrgId:   "e720ed6d82854ab70182f7c6de710d0a",
			//		OrgName: "测试机构",
			//		Org: &models.OrgType{
			//			OrgId:   "e720ed6d82854ab70182f7c7a2b10d20",
			//			OrgName: "测试机构2",
			//			Staff: []*models.StaffType{
			//				&models.StaffType{
			//					StaffId:   "e720ed6d82854ab70182f7cd38730d5d",
			//					StaffName: "测试人员2",
			//				},
			//			},
			//		},
			//	},
			//},
		},
	}
	//b, _ := xml.MarshalIndent(data, "", "	")
	b, _ := xml.Marshal(data)
	//xmlHeader := `<?xml version="1.0" encoding="GBK"?>` + "\n"
	xmlHeader := `<?xml version="1.0" encoding="GBK"?>`
	resp := append([]byte(xmlHeader), b...)
	//UTF-8转换GBK
	respGbk, err := utils.Utf8ToGbk(resp)
	if err != nil {
		fmt.Printf("err=%v", err)
		return nil
	}
	return respGbk
}

func GetFullData2() []byte {

	b, err := os.ReadFile("./example/data-full-1.xml")
	if err != nil {
		return nil
	}
	xmlHeader := `<?xml version="1.0" encoding="GBK"?>`
	resp := append([]byte(xmlHeader), b...)
	//UTF-8转换GBK
	respGbk, err := utils.Utf8ToGbk(resp)
	if err != nil {
		fmt.Printf("err=%v", err)
		return nil
	}
	return respGbk
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

func GetFullData3() []byte {
	/**
	整体思路：
	1. 根节点为 华夏银行
	2. 子节点共有 orgLevel 级，每一级共有 OrgNums个机构，每个机构有 StaffNums 个人。
	*/
	//组织机构数据
	rootOrg := models.HxDataStruct{
		Result:          true,
		MaxSerialNumber: 1248448,
		NowSerialNumber: 1248448,
	}

	// 生成根节点人员
	rootStaffs := GenerateStaff(ROOT)
	//生成根节点子组织数据
	rootSubOrgs := GenerateOrg(ROOT, ROOT_ORG_NUM)

	//生成根节点
	rootOrg.Org = models.OrgType{
		OrgId:   utils.GeneraterIntIdStr(),
		OrgName: "华夏银行",
		Staff:   rootStaffs,
		Org:     rootSubOrgs,
	}
	//生成1级组织数据
	for i := 0; i < len(*rootOrg.Org.Org); i++ {
		(*rootOrg.Org.Org)[i].Staff = GenerateStaff(1)
		(*rootOrg.Org.Org)[i].Org = GenerateOrg(1, LEVEL1_ORG_NUM)
		for j := 0; j < len(*(*rootOrg.Org.Org)[i].Org); j++ {
			(*(*rootOrg.Org.Org)[i].Org)[j].Staff = GenerateStaff(2)
			(*(*rootOrg.Org.Org)[i].Org)[j].Org = GenerateOrg(2, LEVEL2_ORG_NUM)
			for k := 0; k < len(*(*(*rootOrg.Org.Org)[i].Org)[j].Org); k++ {
				(*(*(*rootOrg.Org.Org)[i].Org)[j].Org)[k].Staff = GenerateStaff(3)
				(*(*(*rootOrg.Org.Org)[i].Org)[j].Org)[k].Org = GenerateOrg(3, LEVEL3_ORG_NUM)
				for l := 0; l < len(*(*(*(*rootOrg.Org.Org)[i].Org)[j].Org)[k].Org); l++ {
					(*(*(*(*rootOrg.Org.Org)[i].Org)[j].Org)[k].Org)[l].Staff = GenerateStaff(4)
					(*(*(*(*rootOrg.Org.Org)[i].Org)[j].Org)[k].Org)[l].Org = GenerateOrg(4, LEVEL4_ORG_NUM)
					for n := 0; n < len(*(*(*(*(*rootOrg.Org.Org)[i].Org)[j].Org)[k].Org)[l].Org); n++ {
						(*(*(*(*(*rootOrg.Org.Org)[i].Org)[j].Org)[k].Org)[l].Org)[n].Staff = GenerateStaff(5)
					}
				}
			}
		}
	}

	//--------------------------------
	b, err := xml.Marshal(rootOrg)
	if err != nil {
		fmt.Printf("err = %v", err)
		return nil
	}
	//UTF-8转换GBK
	respGbk, err := utils.Utf8ToGbk(b)
	if err != nil {
		fmt.Printf("err=%v", err)
		return nil
	}
	return respGbk
}

func GenerateStaff(level int) *[]models.StaffType {
	//生成人员数据
	staffsTemp := []models.StaffType{}
	for i := 0; i < StaffNums; i++ {
		staff := models.StaffType{
			StaffId:   utils.GenIDStr(),
			StaffName: utils.GetRandUserName(level),
			Account:   utils.GetRandUserAccount(6),
		}
		staffsTemp = append(staffsTemp, staff)
	}
	return &staffsTemp
}

func GenerateOrg(level, orgNums int) *[]models.OrgType {
	//生成组织数据
	orgTemp := []models.OrgType{}
	for i := 0; i < orgNums; i++ {
		org := models.OrgType{
			OrgId:   utils.GenIDStr(),
			OrgName: utils.GetRandOrgName(level),
		}
		orgTemp = append(orgTemp, org)
	}
	return &orgTemp
}
