package entity

type Access struct {
	Url          string `json:"endpoint" gorm:"column:url;type:varchar(100)"`
	Method       string `json:"method" gorm:"column:method;type:varchar(10)"`
	RequestBody  string `json:"request_body" gorm:"column:request_body;type:text"`
	ResponseBody string `json:"response_body" gorm:"column:response_body;type:text"`
	BaseEntity   `gorm:"embedded"`
}

func (e *Access) TableName() string {
	return "dth_access_log"
}
