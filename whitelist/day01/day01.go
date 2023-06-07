package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

// ip白名单就是防御为未知IP的SYN攻击

// mysql 和 redis 的缓存针对已知的数据查询  缓存查不到；数据库一定是可以查到；

// 因此二者之间是对立的设计。

// 因此可以使用reidis 存储白名单；再使用原生map 来做映射关系。

// 因此关于redis 是存储关于配置信息；mysql 是存储结构化的数据资料的；不是所有的redis 都要落库；
func main() {

	var router = gin.Default()

	// 定义 IP 白名单
	var whitelist = []string{"127.0.0.1", "192.168.0.100"}

	// 使用中间件函数检查 IP 白名单
	router.Use(IPWhiteList(whitelist))

}

func IPWhiteList(whitelist []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求的 IP 地址
		ip := c.ClientIP()
		// 检查 IP 地址是否在白名单中
		allowed := false
		for _, value := range whitelist {
			if value == ip {
				allowed = true
				break
			}
		}
		// 如果 IP 地址不在白名单中，则返回错误信息
		if !allowed {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "IP address not allowed"})
			return
		}
		// 允许请求继续访问后续的处理函数
		c.Next()
	}
}

func SystemIpWhite(c *gin.Context) {
	ip := c.ClientIP()

	//白名单数据源
	blackList := models.FindConfig("SystemWhiteList")
	strings.ReplaceAll(blackList, "\r\n", "\n")
	list := strings.Split(blackList, "\n")
	exist := false
	for _, word := range list {
		word = strings.Trim(word, " ")
		if word == "" {
			continue
		}
		if ip == word {
			exist = true
			break
		}
	}

	if !exist {
		log.Println("ip whitelist forbidden", ip)
		c.String(403, "403 forbidden")
		c.Abort()
		return
	}
}
