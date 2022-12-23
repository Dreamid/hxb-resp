package models

//自定义外部数据源相关结构体

//自定义外部身份数据源
type ThirdCommonIdp struct {
	Id int64 `orm:"pk;auto" json:"id"`
	//外部数据源名称
	Name string `json:"name"`
	//服务器地址
	Addr string `json:"addr"`
	//数据源简称
	Tag string `json:"tag"`
	//zeus数据源ID
	IdpId string `orm:"type(text)" json:"-"`
	//zeus数据源服务器ID
	IdpServerId string `orm:"type(text)" json:"-"`
	//同步设置相关配置
	//同步设置类型 1每月  2每周   3每天
	SyncSetupType uint32 `json:"syncSetupType"`
	//同步方式 1: 每隔N个小时周期同步   2: 在某几个时间点同步
	SyncType uint32 `json:"syncType"`
	//同步周期间隔N个小时，如果syncType=1 则该值不能为空dd
	SyncPeriod uint32 `json:"syncPeriod"`
	//同步时间点，整数0-24之间，如果syncType=2 则该值不能为空
	SyncTime []uint32 `orm:"type(json)" json:"syncTime"`
	//每月的日期1-31/每周的日期1-7，如果syncType=1 or syncType=2 则该值不能为空
	SyncSetupDay []uint32 `orm:"type(json)" json:"syncSetupDay"`
	//同步时间点，00:00 -  23:59
	SyncSetupTime string `orm:"type(text)" json:"syncSetupTime"`
	Created       int64  `json:"-"`
	Modified      int64  `json:"-"`
}

//自定义外部数据源-组织
type ThirdGroups struct {
	Id          int64  `orm:"pk;auto" json:"id"`
	GroupId     string `orm:"index" json:"groupId"`
	Name        string `orm:"" json:"name"`
	IsRoot      int8   `orm:"" json:"isRoot"`
	NamePinyin  string `orm:"" json:"namePinyin"`
	ParentId    string `orm:"index" json:"parentId"`
	Parents     string `orm:"type(text)" json:"parents"`
	Order       int64  `orm:"" json:"order"`
	State       int8   `orm:"" json:"state"`
	Desc        string `orm:"" json:"desc"`
	IdpId       string `orm:"" json:"idpId"`
	ExtraFields string `orm:"type(text)" json:"extraFields"`
	Created     int64  ` json:"created"`
	Modified    int64  `orm:"index" json:"modified"`
	Deleted     int64  `json:"deleted"`
}

//自定义外部数据源-人员
type ThirdStaffs struct {
	Id                uint64 `orm:"pk;auto" json:"id"`
	StaffId           string `orm:"index" json:"staff_id"`
	RealName          string `json:"real_name"`
	UserName          string `json:"user_name"`
	Password          string `json:"password"`
	Email             string `json:"email"`
	CountryCode       int16  `json:"country_code"`
	PhoneNumber       string `json:"phone_number"`
	State             int8   `json:"state"`
	Order             int64  `json:"order"`
	Department        string `json:"department"`
	ExpiredTime       int64  `json:"expired_time"`
	Duty              string `json:"duty"`
	EmployeeNumber    string `json:"employee_number"`
	Desc              string `json:"desc"`
	FixedNumber       string `json:"fixed_number"`
	FixedCountry_code int16  `json:"fixed_country_code"`
	IdpId             string `json:"idp_id"`
	ExtraFields       string `orm:"type(text)" json:"extra_fields"`
	Created           int64  ` json:"created"`
	Modified          int64  `orm:"index" json:"modified"`
	Deleted           int64  `json:"deleted"`
}

//自定义外部数据源 组织-人员关系
type ThirdGroupStaffs struct {
	Id       uint64 `orm:"pk;auto" json:"id"`
	MemberId string `json:"member_id"` //三方的组织-人员对应关系的唯一id
	GroupId  string `json:"group_id"`
	StaffId  string `json:"staff_id"`
	State    int64  `json:"state"`
	IdpId    string `json:"idp_id"`
	Created  int64  ` json:"created"`
	Modified int64  `orm:"index" json:"modified"`
	Deleted  int64  `json:"deleted"`
}

func (u *ThirdGroupStaffs) TableIndex() [][]string {
	return [][]string{{"GroupId", "StaffId"}}
}

//数据源同步情况表
type ZeusSyncTask struct {
	Id       int64
	IdpId    int64
	DataType SyncDataType
	Version  int64
}

type SyncDataType int

const (
	SYNCGROUP SyncDataType = iota + 1
	SYNCSTAFF
	SYNCGROUPANDSTAFF
)

const (
	//全量同步标识
	SYNC_ALL = "sync_all"
	//增量同步标识
	SYNC_FAST = "sync_fast"
	//定时同步标识
	SYNC_CRON = "sync_cron"
)
