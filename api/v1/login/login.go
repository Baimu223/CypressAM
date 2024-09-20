package login

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/patrickmn/go-cache"
	"strconv"
	"time"
	"xkginweb/commons/jwtgo"
	"xkginweb/commons/response"
	"xkginweb/global"
	msys "xkginweb/model/entity/sys"
	"xkginweb/utils"
	"xkginweb/utils/adr"
)

// 登录业务
type LoginApi struct{}

// 1: 定义验证的store --默认是存储在go内存中
var store = base64Captcha.DefaultMemStore

// 登录的接口处理
func (api *LoginApi) ToLogined(c *gin.Context) {
	type LoginParam struct {
		Account  string
		Code     string
		CodeId   string
		Password string
	}

	// 1：获取用户在页面上输入的账号和密码开始和数据库里数据进行校验
	param := LoginParam{}
	err2 := c.ShouldBindJSON(&param)
	if err2 != nil {
		response.Fail(60002, "参数绑定有误", c)
		return
	}

	if len(param.Code) == 0 {
		response.Fail(60002, "请输入验证码", c)
		return
	}

	if len(param.CodeId) == 0 {
		response.Fail(60002, "验证码获取失败", c)
		return
	}

	// 开始校验验证码是否正确
	verify := store.Verify(param.CodeId, param.Code, true)
	if !verify {
		response.Fail(60002, "你输入的验证码有误!!", c)
		return
	}

	inputAccount := param.Account
	inputPassword := param.Password

	if len(inputAccount) == 0 {
		response.Fail(60002, "请输入账号", c)
		return
	}

	if len(inputPassword) == 0 {
		response.Fail(60002, "请输入密码", c)
		return
	}

	dbUser, err := sysUserService.GetUserByAccount(inputAccount)
	if err != nil {
		response.Fail(60002, "你输入的账号和密码有误", c)
		return
	}

	// 这个时候就判断用户输入密码和数据库的密码是否一致
	// inputPassword = utils.Md5(123456) = 2ec9f77f1cde809e48fabac5ec2b8888
	// dbUser.Password = 2ec9f77f1cde809e48fabac5ec2b8888
	if dbUser != nil && dbUser.Password == adr.Md5Slat(inputPassword, dbUser.Slat) {
		// 根据用户id查询用户的角色
		userRoles, _ := sysUserRolesService.SelectUserRoles(dbUser.ID)
		if len(userRoles) > 0 {
			// 用户信息生成token -----把
			token := api.generaterToken(c, userRoles[0].RoleCode, userRoles[0].ID, dbUser)
			// 根据用户查询菜单信息
			roleMenus, _ := sysRoleMenusService.SelectRoleMenus(userRoles[0].ID)
			// 根据用户id查询用户的角色的权限
			permissions, _ := sysRoleApisService.SelectRoleApis(userRoles[0].ID)

			// 这个uuid是用于挤下线使用
			uuid := utils.GetUUID()
			userIdStr := strconv.FormatUint(uint64(dbUser.ID), 10)
			global.Cache.Set("LocalCache:Login:"+userIdStr, uuid, cache.NoExpiration)

			// 查询返回
			response.Ok(map[string]any{
				"user":        dbUser,
				"token":       token,
				"roles":       userRoles,
				"uuid":        uuid,
				"roleMenus":   sysMenuService.Tree(roleMenus, 0),
				"permissions": permissions}, c)
		} else {
			// 查询返回--
			response.Fail(80001, "你暂无授权信息", c)
		}
	} else {
		response.Fail(60002, "你输入的账号和密码有误", c)
	}
}

/*
* 根据用户信息创建一个token
 */
func (api *LoginApi) generaterToken(c *gin.Context, roleCode string, roleId uint, dbUser *msys.SysUser) string {

	// 设置token续期的缓冲时间
	bf, _ := utils.ParseDuration("1d")
	ep, _ := utils.ParseDuration("7d")

	// 1: jwt生成token
	myJwt := jwtgo.NewJWT()
	// 2: 生成token
	token, err2 := myJwt.CreateToken(jwtgo.CustomClaims{
		dbUser.ID,
		dbUser.Username,
		roleCode,
		roleId,
		int64(bf / time.Second),
		jwt.StandardClaims{
			Audience:  "KSD",                        // 受众
			Issuer:    "KSD-ADMIN",                  // 签发者
			IssuedAt:  time.Now().Unix(),            // 签发时间
			NotBefore: time.Now().Add(-1000).Unix(), // 生效时间
			ExpiresAt: time.Now().Add(ep).Unix(),    // 过期时间

		},
	})

	fmt.Println("当前时间是：", time.Now().Unix())
	fmt.Println("缓冲时间：", int64(bf/time.Second))
	fmt.Println("签发时间：" + time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("生效时间：" + time.Now().Add(-1000).Format("2006-01-02 15:04:05"))
	fmt.Println("过期时间：" + time.Now().Add(ep).Format("2006-01-02 15:04:05"))

	if err2 != nil {
		response.Fail(60002, "登录失败，token颁发不成功!", c)
	}

	return token
}
