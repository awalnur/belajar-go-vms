package repository

import (
	"gorm.io/gorm"
	"time"
)

type UserHistory struct {
	ID          int    `gorm:"type:bigint;primaryKey;autoIncrement"`
	UserID      int    `gorm:"index;default:null"`
	User        Users  `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`
	Action      string `gorm:"not null;index"`
	Description string
	CreatedDate time.Time      `gorm:"type:date;default:CURRENT_DATE"`
	CreatedAt   time.Time      `gorm:"type:timestamp(0);default:CURRENT_TIMESTAMP"`
	UpdatedAt   *time.Time     `gorm:"type:timestamp(0);default:null"`
	DeletedAt   gorm.DeletedAt `gorm:"type:timestamp(0);default:null"`
}

type UserHistoryRepoConn struct {
	DB  *gorm.DB
	Err error
}
