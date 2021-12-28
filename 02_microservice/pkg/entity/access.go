package entity

type Access struct {
	Url          string `json:"endpoint" gorm:"type:varchar(100)"`
	Method       string `json:"method" gorm:"type:varchar(10)"`
	RequestBody  string `json:"request_body" gorm:"type:text;"`
	ResponseBody string `json:"response_body" gorm:"type:text"`
	BaseEntity   `gorm:"embedded"`
}

func (e *Access) TableName() string {
	return "dth_access_log"
}
