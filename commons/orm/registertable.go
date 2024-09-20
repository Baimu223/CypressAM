package orm

import (
	"xkginweb/global"
	bbs2 "xkginweb/model/entity/bbs"
	"xkginweb/model/entity/jwt"
	sys2 "xkginweb/model/entity/sys"
	user2 "xkginweb/model/entity/user"
	video2 "xkginweb/model/entity/video"
)

func RegisterTable() {
	db := global.KSD_DB
	// 注册和声明model
	db.AutoMigrate(user2.XkUser{})
	db.AutoMigrate(user2.XkUserAuthor{})
	// 系统用户，角色，权限表
	db.AutoMigrate(sys2.SysApis{})
	db.AutoMigrate(sys2.SysMenus{})
	db.AutoMigrate(sys2.SysRoleApis{})
	db.AutoMigrate(sys2.SysRoleMenus{})
	db.AutoMigrate(sys2.SysRoles{})
	db.AutoMigrate(sys2.SysUserRoles{})
	db.AutoMigrate(sys2.SysUser{})
	// 视频表
	db.AutoMigrate(video2.XkVideo{})
	db.AutoMigrate(video2.XkVideoCategory{})
	db.AutoMigrate(video2.XkVideoChapterLesson{})
	// 社区
	db.AutoMigrate(bbs2.XkBbs{})
	db.AutoMigrate(bbs2.BbsCategory{})

	// 声明一下jwt模型
	db.AutoMigrate(jwt.JwtBlacklist{})
}
