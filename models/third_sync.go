package models

type ThirdSyncSettings struct {
	Id                   uint32 `orm:"pk;auto" json:"id"`
	ThirdName            string
	LastSyncSerialNumber int64
}
