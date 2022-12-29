package models

import "encoding/xml"

//全量数据结构
type HxDataStruct struct {
	XMLName         xml.Name `xml:"组织机构数据"`
	Result          bool     `xml:"结果"`
	Info            string   `xml:"信息"`
	MaxSerialNumber int64    `xml:"最大流水号"`
	NowSerialNumber int64    `xml:"当前流水号"`
	//该字段只存在于全量数据
	Org OrgType `xml:"机构"`
	//该字段仅存在于增量数据
	DataIncr DataIncrType `xml:"数据增量"`
}

//机构
type OrgType struct {
	OrgId   string       `xml:"机构ID"`
	OrgName string       `xml:"名称"`
	Staff   *[]StaffType `xml:"人员"`
	Org     *[]OrgType   `xml:"机构"`
}

//人员
type StaffType struct {
	StaffId   string `xml:"个人ID"`
	StaffName string `xml:"姓名"`
	Account   string `xml:"对应帐号"`
}

type DataIncrType struct {
	ChangeData []*ChangeDataType `xml:"机构人员变动"`
}

type ChangeDataType struct {
	SerialNumber  int64              `xml:"流水号"`
	Time          string             `xml:"时间"`
	StaffChange   *StaffChangeType   `xml:"个人变更"`
	OrgChange     *OrgChangeType     `xml:"机构变更"`
	AccountChange *AccountChangeType `xml:"帐号变更"`
}

type StaffChangeType struct {
	Operation string    `xml:"操作"`
	StaffId   string    `xml:"个人ID"`
	OrgId     string    `xml:"机构ID"`
	Staff     StaffType `xml:"人员"`
	//仅用于<个人变更>类型中的<调入><调出>
	ParentOrgId string `xml:"上级机构ID"`
}
type OrgChangeType struct {
	Operation string  `xml:"操作"`
	OrgId     string  `xml:"机构ID"`
	PId       string  `xml:"上级机构ID"`
	Org       OrgType `xml:"机构"`
}
type AccountChangeType struct {
	Operation string `xml:"操作"`
	StaffId   string `xml:"个人ID"`
	Account   string `xml:"帐号"`
}

//个人变更
const (
	//调整现有员工的所属机构，会产生<调入><调出>数据；
	STAFF_IN  = "调入"
	STAFF_OUT = "调出"
	//不处理（处理逻辑包含在增加上下级关系逻辑中）
	STAFF_ADD          = "增加"
	STAFF_MOD          = "修改"
	STAFF_DEL          = "删除"
	STAFF_RELATION_ADD = "增加上下级关系"
	STAFF_RELATION_DEL = "删除上下级关系"
)

//机构变更
const (
	//不处理（处理逻辑包含在 增加上下级关系逻辑中）
	ORG_ADD          = "增加"
	ORG_MOD          = "修改"
	ORG_DEL          = "删除"
	ORG_RELATION_ADD = "增加上下级关系"
	ORG_RELATION_DEL = "删除上下级关系"
)

//帐号变更
const (
	ACCOUNT_ADD = "增加"
	ACCOUNT_DEL = "删除"
)

const (
	THIRD_CN_NAME = "华夏银行"
	THIRD_TAG     = "hxbk"
)

const SYS_CODE = "Qaxbrowser"
