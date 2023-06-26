package repository

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
	"time"
)

type Privileges struct {
	ID           int `gorm:"bigint;primaryKey;autoIncrement"`
	Name         string
	NameCode     string
	Description  string
	StatusAccess pq.BoolArray   `gorm:"type:boolean[]"`
	Status       bool           `gorm:"not null; default:true"`
	CreatedBy    int            `gorm:"index; default:nulll"`
	CreatedUser  *Users         `gorm:"foreignKey: CreatedBy"`
	UpdatedBy    int            `gorm:"index;default:null"`
	UpdatedUser  *Users         `gorm:"foreignKey: UpdatedBy"`
	DeletedBy    int            `gorm:"index;default:null"`
	DeletedUser  *Users         `gorm:"foreignKey:DeletedBy"`
	CreatedAt    time.Time      `gorm:"type:timestamp(0); default:CURRENT_TIMESTAMP"`
	UpdatedAt    *time.Time     `gorm:"type:timestamp(0); default:null"`
	DeletedAt    gorm.DeletedAt `gorm:"type:timestamp(0);default:null"`
}

type PrivilegesRepoCon struct {
	DB  *gorm.DB
	Err error
}
