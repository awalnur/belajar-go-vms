package repository

import (
	"gorm.io/gorm"
	"time"
)

type Users struct {
	Id          int       `gorm:"type:bigint;primaryKey;autoIncrement"`
	Email       string    `gorm:"uniqueIndex"`
	Username    string    `gorm:"uniqueIndex"`
	Password    string    `gorm:"not null"`
	Name        string    `gorm:"not null"`
	BirtDate    time.Time `gorm:"type:date;default:null"`
	Gender      string    `gorm:"default:null"`
	Status      bool      `gorm:"not null; default:True"`
	Banned      bool      `gorm:"not null; default:False"`
	BannedTime  time.Time `gorm:"type:date; default:null"`
	FailedAuth  int       `gorm:"type:int;default:0"`
	RoleId      int
	Role        Roles          `gorm:"foreignKey:RoleId"`
	History     []UserHistory  `gorm:"foreignKey:UserID"`
	CreateBy    int            `gorm:"index;default:null"`
	CreatedUser *Users         `gorm:"foreignKey:createdBy;association_foreignKey:Id"`
	UpdatedBy   int            `gorm:"index;default: null"`
	UpdatedUser *Users         `gorm:"foreignKey:UpdatedBy;association_foreignKey:Id"`
	DeletedBy   int            `gorm:"index;default:null"`
	DeletedUser *Users         `gorm:"foreignKey: DeletedBy; assocoation_foreignKey:Id"`
	CreatedAt   time.Time      `gorm:"type:timestamp(0);default:CURRENT_TIMESTAMP"`
	UpdatedAt   *time.Time     `gorm:"type:timestamp(0);default:null"`
	DeletedAt   gorm.DeletedAt `gorm:"type:timestamp(0);default:null"`
}

type UserRepoConn struct {
	DB  *gorm.DB
	Err error
}
