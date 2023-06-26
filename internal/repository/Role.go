package repository

import (
	"gorm.io/gorm"
	"time"
)

type Roles struct {
	ID          int `gorm:"type:bigint;primaryKey;autoIncrement"`
	Name        string
	Description string
	Status      bool           `gorm:"not null; default: True"`
	CreatedBy   int            `gorm:"index;default:null"`
	CreatedUser *Users         `gorm:"foreignKey:CreatedBy;association_foreignKey:ID"`
	UpdatedBy   int            `gorm:"index;default:null"`
	UpdatedUser *Users         `gorm:"foreignKey:UpdatedBy; association_foreignKey:ID"`
	DeletedBy   int            `gorm:"index;default:null"`
	DeletedUser *Users         `gorm:"foreignKey:DeletedBy;association_foreignKey:ID"`
	CreatedAt   time.Time      `gorm:"type:timestamp(0);default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time      `gorm:"type:timestamp(0); default:null"`
	DeletedAt   gorm.DeletedAt `gorm:"type:timestamp(0); default:null"`
}
type RoleRepoConn struct {
	DB  *gorm.DB
	Err error
}

func (conn *RoleRepoConn) GetRoleById(id int) (role *Roles, err error) {
	if conn.Err != nil {
		return nil, err
	}
	result := conn.DB.Model(&Roles{ID: id}).First(&role, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return role, nil
}
