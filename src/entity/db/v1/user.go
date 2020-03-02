package v1

import "time"

// Users Mapping
type Users struct {
	ID             uint       `gorm:"primary_key" json:"id"`
	Name           string     `gorm:"column:name;not null;type:varchar(50)" json:"name"`
	Email          string     `gorm:"column:email;unique;not null;type:varchar(100)" json:"email"`
	Password       string     `gorm:"column:password;not null;type:varchar(100)" json:"password"`
	DateOfBirth    string     `gorm:"column:date_of_birth;not null;type:date" json:"date_of_birt"`
	PhoneNumber    string     `gorm:"column:phone_number;not null;type:varchar(15)" json:"phone_number"`
	CurrentAddress string     `gorm:"column:current_address;type:varchar(255)" json:"current_address" form:"current_address"`
	City           string     `gorm:"column:city;type:varchar(50)" json:"city" form:"city"`
	Province       string     `gorm:"column:province;type:varchar(50)" json:"province" form:"province"`
	District       string     `gorm:"column:district;type:varchar(50)" json:"district" form:"district"`
	Lat            float64    `gorm:"column:lat;type:double" json:"lat" form:"lat"`
	Lng            float64    `gorm:"column:lng;type:double" json:"lng" form:"lng"`
	Job            string     `gorm:"column:job;type:varchar(255)" json:"job" form:"job"`
	Image          string     `gorm:"column:image;type:varchar(255)" json:"image" form:"image"`
	StartIncome    int        `gorm:"column:start_income;type:int unsigned" json:"start_income" form:"start_income"`
	EndIncome      int        `gorm:"column:end_income;type:int unsigned" json:"end_income" form:"end_income"`
	Interest       string     `gorm:"column:interests;type:varchar(255)" json:"interests" form:"interests"`
	Habit          string     `gorm:"column:habit;type:varchar(100)" json:"habit" form:"habit"`
	IsOrganizer    bool       `gorm:"column:is_organizer;type:bool" json:"is_organizer" form:"is_organizer"`
	CreatedAt      *time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      *time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt      *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName User Return Table
// return: tablename
func (Users) TableName() string {
	return "users"
}
