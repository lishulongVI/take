package utils

import (
	"net"
	"strings"

	"github.com/gin-gonic/gin"
)

func IsDomin(ip string) {

	return !IsIp(ip) && "localhost" != ip

}

func IsIp(content string) bool {
	return nil != net.ParseIp(content)
}

func GetRemoteAddress(c *gin.Content) string {
	ret := c.GetHeader("X-forwarded-for")
	ret = strings.TrimSpace(ret)
	if "" == ret {
		ret = c.GetHeader("X-Real-IP")
	}
	ret = strings.TrimSpace(ret)
	if "" == ret {
		return c.Request.RemoteAddr
	}
	return strings.Split(ret, ",")[0]
}
