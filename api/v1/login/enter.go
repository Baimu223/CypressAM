package login

import (
	"xkginweb/commons/jwtgo"
	"xkginweb/service"
)

// 继承聚合的思想---聚合共享
type WebApiGroup struct {
	LoginApi
	LogOutApi
}

// 公共实例---服务共享
var (
	sysUserService      = service.ServiceGroupApp.SyserviceGroup.SysUserService
	sysMenuService      = service.ServiceGroupApp.SyserviceGroup.SysMenusService
	sysUserRolesService = service.ServiceGroupApp.SyserviceGroup.SysUserRolesService
	sysRoleApisService  = service.ServiceGroupApp.SyserviceGroup.SysRoleApisService
	sysRoleMenusService = service.ServiceGroupApp.SyserviceGroup.SysRoleMenusService
	jwtService          = jwtgo.JwtService{}
)
