package middleware

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var keywords = []string{
	// 基本 DML/DQL/DCL 操作
	"select",
	"insert",
	"update",
	"delete",
	"drop",
	"truncate",
	"exec",
	"execute",
	"union",
	"create",
	"alter",
	"rename",
	"grant",
	"revoke",
	"show",
	"describe",

	// 系统库
	"information_schema",
	"table_schema",

	// 时间延迟攻击
	"sleep",
	"benchmark",

	// 注释语法
	"--",
	"#",
	";",

	// 布尔型注入
	"or ",
	"and ",
	"1=1",
	"1 = 1",
	"' or",
	"\" or",
	"' and",
	"\" and",

	// 高危函数和系统对象
	"xp_cmdshell",
	"sysobjects",
	"syscolumns",
	"char(",
	"concat(",
	"cast(",
	"convert(",

	// 文件操作
	"declare",
	"set global",
	"load_file(",
	"outfile",
	"load data",
	"into outfile",
}

type RateLimiter struct {
	ips      map[string][]time.Time
	mutex    sync.Mutex
	limit    int
	window   time.Duration
	cleanup  time.Duration
	stopChan chan struct{}
}

func NewRateLimiter(limit int, window time.Duration, cleanup time.Duration) *RateLimiter {
	limiter := &RateLimiter{
		ips:      make(map[string][]time.Time),
		limit:    limit,
		window:   window,
		cleanup:  cleanup,
		stopChan: make(chan struct{}),
	}

	go limiter.cleanupLoop()

	return limiter
}

func (rl *RateLimiter) cleanupLoop() {
	ticker := time.NewTicker(rl.cleanup)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			rl.cleanupExpired()
		case <-rl.stopChan:
			return
		}
	}
}

func (rl *RateLimiter) cleanupExpired() {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()

	now := time.Now()
	for ip, times := range rl.ips {
		var validTimes []time.Time
		for _, t := range times {
			if now.Sub(t) < rl.window {
				validTimes = append(validTimes, t)
			}
		}
		if len(validTimes) == 0 {
			delete(rl.ips, ip)
		} else {
			rl.ips[ip] = validTimes
		}
	}
}

func (rl *RateLimiter) Allow(ip string) bool {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()

	now := time.Now()
	times := rl.ips[ip]

	var validTimes []time.Time
	for _, t := range times {
		if now.Sub(t) < rl.window {
			validTimes = append(validTimes, t)
		}
	}

	if len(validTimes) >= rl.limit {
		rl.ips[ip] = validTimes
		return false
	}

	validTimes = append(validTimes, now)
	rl.ips[ip] = validTimes
	return true
}

func (rl *RateLimiter) Stop() {
	close(rl.stopChan)
}

func containsMaliciousKeyword(s string) bool {
	s = strings.ToLower(s)
	for _, keyword := range keywords {
		if strings.Contains(s, keyword) {
			return true
		}
	}
	return false
}

func SecurityMiddleware() gin.HandlerFunc {
	rateLimiter := NewRateLimiter(60, time.Minute, time.Minute*5)

	return func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 1<<20)

		ip := c.ClientIP()

		if !rateLimiter.Allow(ip) {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"status":  "error",
				"message": "请求过于频繁，请稍后再试",
			})
			c.Abort()
			return
		}

		// 检查URL参数
		for _, values := range c.Request.URL.Query() {
			for _, value := range values {
				if containsMaliciousKeyword(value) {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  "error",
						"message": "请求包含恶意关键词",
					})
					c.Abort()
					return
				}
			}
		}

		// 检查请求头
		for name, values := range c.Request.Header {
			if name != "User-Agent" && name != "Content-Type" {
				for _, value := range values {
					if containsMaliciousKeyword(value) {
						c.JSON(http.StatusBadRequest, gin.H{
							"status":  "error",
							"message": "请求包含恶意关键词",
						})
						c.Abort()
						return
					}
				}
			}
		}

		// 检查请求体
		if c.Request.Body != nil {
			body, err := io.ReadAll(c.Request.Body)
			if err == nil {
				if containsMaliciousKeyword(string(body)) {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  "error",
						"message": "请求包含恶意关键词",
					})
					c.Abort()
					return
				}
				// 重置请求体
				c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
			}
		}

		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")

		c.Next()
	}
}

func BlockSQLInjectionInParmametersAndBody() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 检查URL参数
		for _, values := range c.Request.URL.Query() {
			for _, value := range values {
				if containsMaliciousKeyword(value) {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  "error",
						"message": "请求包含恶意关键词",
					})
					c.Abort()
					return
				}
			}
		}

		// 检查请求头
		for name, values := range c.Request.Header {
			if name != "User-Agent" && name != "Content-Type" {
				for _, value := range values {
					if containsMaliciousKeyword(value) {
						c.JSON(http.StatusBadRequest, gin.H{
							"status":  "error",
							"message": "请求包含恶意关键词",
						})
						c.Abort()
						return
					}
				}
			}
		}

		// 检查请求体
		if c.Request.Body != nil {
			body, err := io.ReadAll(c.Request.Body)
			if err == nil {
				if containsMaliciousKeyword(string(body)) {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  "error",
						"message": "请求包含恶意关键词",
					})
					c.Abort()
					return
				}
				// 重置请求体
				c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
			}
		}

		c.Next()
	}
}