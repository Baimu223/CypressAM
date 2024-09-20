package global

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type GVA_MODEL struct {
	ID        uint                  `gorm:"primarykey;comment:主键ID" json:"id" structs:"-"` // 主键ID
	CreatedAt time.Time             `gorm:"type:datetime(0);autoCreateTime;comment:创建时间" json:"createdAt" structs:"-"`
	UpdatedAt time.Time             `gorm:"type:datetime(0);autoUpdateTime;comment:更新时间" json:"updatedAt" structs:"-"`
	IsDeleted soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt;default:0" json:"isDeleted" structs:"is_deleted"`
}
