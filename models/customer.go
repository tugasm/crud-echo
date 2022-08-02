package models

type ErrMeta struct {
	ServiceCode string
	FieldErr    string
}

type Customer struct {
	// gorm.Model `json:"-"`
	Id                    uint   `gorm:"primary_key"`
	Date                  string `gorm:"type:varchar(25)" json:"date"`
	Time                  string `gorm:"type:timestamp" json:"time"`
	Devapps               string `gorm:"type:varchar(100)" json:"devapps"`

}

type CustomerRequest struct {
	Date                  string `gorm:"type:varchar(25)" json:"date"`
	Time                  string `gorm:"type:timestamp" json:"time"`
	Devapps               string `gorm:"type:varchar(100)" json:"devapps"`

}