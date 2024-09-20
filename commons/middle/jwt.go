package middle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
	"xkginweb/commons/jwtgo"
	"xkginweb/commons/response"
	"xkginweb/global"
	"xkginweb/model/entity/jwt"
	"xkginweb/utils"
)

var jwtService = jwtgo.JwtService{}

// 定义一个JWTAuth的中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 通过http header中的token解析来认证
		// 获取token
		token := c.GetHeader("Authorization")
		if token == "" {
			response.Fail(4001, "请求未携带token，无权限访问", c)
			c.Abort()
			return
		}

		// 判断是否携带登录状态uuid
		KsdUUID := c.GetHeader("KsdUUID")
		if KsdUUID == "" {
			response.Fail(4001, "请求未携带登录标识，无权限访问", c)
			c.Abort()
			return
		}

		// 这里会判断你当前的token是不是在黑名单中
		blacklist := jwtService.IsBlacklist(token)
		if blacklist {
			response.Fail(4001, "黑名单提醒：你的token已经失效了，", c)
			c.Abort()
			return
		}

		// 生成jwt的对象
		myJwt := jwtgo.NewJWT()
		// 解析token
		customClaims, err := myJwt.ParserToken(token)
		// 如果解析失败就出现异常
		if err != nil {
			response.Fail(4001, "token失效了", c)
			c.Abort()
			return
		}

		// 从缓存中获取服务端用户uuid是否和用户传递进来的uuid是否一致，
		// 1:如果一致让操作，
		// 2:如果不一致直接提示你当前账号已经被挤下线了
		userIdStr := strconv.FormatUint(uint64(customClaims.UserId), 10)
		cacheUuid, _ := global.Cache.Get("LocalCache:Login:" + userIdStr)
		// 可能缓存被清理了
		if cacheUuid == "" {
			response.Fail(4001, "请求未携带登录标识，无权限访问", c)
			c.Abort()
			return
		}

		// 如果不相等，说明用户在别的地方登录了
		if cacheUuid != KsdUUID {
			response.Fail(4001, "你账号已被挤下线！", c)
			c.Abort()
			return
		}

		// 判断过期时间 - now  < buffertime 就开始续期 ep 1d -- no
		fmt.Println("customClaims.ExpiresAt", customClaims.ExpiresAt)
		fmt.Println("time.Now().Unix()", time.Now().Unix())
		fmt.Println("customClaims.ExpiresAt - time.Now().Unix()", customClaims.ExpiresAt-time.Now().Unix())
		fmt.Println("customClaims.BufferTime", customClaims.BufferTime)

		if customClaims.ExpiresAt-time.Now().Unix() < customClaims.BufferTime {
			// 1: 生成一个新的token
			// 2: 用c把新的token返回页面
			fmt.Println("开始续期.....")
			// 获取7天的过期时间
			eptime, _ := utils.ParseDuration("7d")
			// 用当前时间+eptime 就是 新的token过期时间
			customClaims.ExpiresAt = time.Now().Add(eptime).Unix()
			// 生成新的token
			newToken, _ := myJwt.CreateTokenByOldToken(token, *customClaims)
			// 输出给浏览器----request--header----给服务端
			// 输出给浏览器----response--header---给浏览器
			c.Header("new-authorization", newToken)
			c.Header("new-expires-at", strconv.FormatInt(customClaims.ExpiresAt, 10))
			// 如果生成新token了，旧的token怎么办？ jwt没有提供一个机制让旧token失效。
			_ = jwtService.JsonInBlacklist(jwt.JwtBlacklist{Jwt: token})
		}

		// 让后续的路由方法可以直接通过c.Get("claims")
		c.Set("claims", customClaims)
		c.Set("userId", customClaims.UserId)
		// 打印日志
		//web.Print(c)
		// 开始执行具体的router方法
		c.Next()
	}
}
