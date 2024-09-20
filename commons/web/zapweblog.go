package web

import (
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"time"
	"xkginweb/global"
)

type ResponseWriterWrapper struct {
	gin.ResponseWriter
	Body *bytes.Buffer // 缓存
}

func (w ResponseWriterWrapper) Write(b []byte) (int, error) {
	w.Body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w ResponseWriterWrapper) WriteString(s string) (int, error) {
	w.Body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

/**
 * 打印日志
 */
func Print(c *gin.Context) {
	//reqBody, err := io.ReadAll(c.Request.Body)
	//var str string
	//if err == nil {
	//	str = string(reqBody)
	//}
	global.Log.Debug("HTTP CALL START",
		zap.String("callTime", time.Now().Format("2006-01-02 15:04:05")),
		zap.String("schema ", c.Request.Proto),
		zap.String("host", c.Request.Host),
		zap.String("requestUri", c.Request.RequestURI),
		zap.String("method", c.Request.Method),
		zap.String("ip", c.ClientIP()),
		zap.String("remoteAddr", c.Request.RemoteAddr),
		zap.Any("requestHeaders", c.Request.Header),
		zap.Any("requestParams", GetParams(c)),
		//zap.Any("body", str),
	)

	blw := &ResponseWriterWrapper{Body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw
	c.Next()
	global.Log.Debug("HTTP CALL END",
		zap.Int("status", c.Writer.Status()),
		zap.Any("responseHeaders", c.Writer.Header()),
		zap.String("responseData", blw.Body.String()),
		zap.Int("responseSize", c.Writer.Size()),
		zap.String("resolveTime", time.Now().Format("2006-01-02 15:04:05")),
	)
}

func GetParams(c *gin.Context) map[string]any {
	if strings.ToLower(c.Request.Method) == "get" {
		return GetQueryParams(c)
	}
	if strings.ToLower(c.Request.Method) == "post" {
		return GetPostFormParams(c)
	}
	return nil
}

func GetQueryParams(c *gin.Context) map[string]any {
	query := c.Request.URL.Query()
	var queryMap = make(map[string]any, len(query))
	for k := range query {
		queryMap[k] = c.Query(k)
	}
	return queryMap
}

func GetPostFormParams(c *gin.Context) map[string]any {
	if err := c.Request.ParseMultipartForm(32 << 20); err != nil {
		if !errors.Is(err, http.ErrNotMultipart) {
			return nil
		}
	}
	var postMap = make(map[string]any, len(c.Request.PostForm))
	for k, v := range c.Request.PostForm {
		if len(v) > 1 {
			postMap[k] = v
		} else if len(v) == 1 {
			postMap[k] = v[0]
		}
	}

	return postMap
}
