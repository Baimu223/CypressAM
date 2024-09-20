package sys

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type SysRoles struct {
	ID        uint                  `gorm:"primarykey;comment:主键ID" json:"id" structs:"-"` // 主键ID
	CreatedAt time.Time             `gorm:"type:datetime(0);autoCreateTime;comment:创建时间" json:"createdAt" structs:"-"`
	UpdatedAt time.Time             `gorm:"type:datetime(0);autoUpdateTime;comment:更新时间" json:"updatedAt" structs:"-"`
	IsDeleted soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt;default:0" json:"isDeleted" structs:"is_deleted"`
	RoleName  string                `json:"roleName" gorm:"comment:角色名"`  // 角色名
	RoleCode  string                `json:"roleCode" gorm:"comment:角色代号"` // 角色代号
}

func (s *SysRoles) TableName() string {
	return "sys_roles"
}
