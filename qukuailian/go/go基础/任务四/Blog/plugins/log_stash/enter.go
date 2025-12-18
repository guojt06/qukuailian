package log_stash

import (
	"modulename/global"
	"modulename/utils/jwts"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Log struct {
	Ip     string `json:"ip"`
	Addr   string `json:"addr"`
	UserId uint   `json:"userId"`
}

func New(ip string, token string) *Log {
	// 进行解析token
	claims, err := jwts.ParseToken(token)
	var userId uint
	if err == nil {
		userId = claims.UserId
	}
	return &Log{
		Ip:     ip,
		Addr:   "内网",
		UserId: userId,
	}
}
func NewLogByGin(ctx *gin.Context) *Log {
	ip := ctx.ClientIP()
	token := ctx.Request.Header.Get("token")
	return New(ip, token)
}

func (l Log) Debug(content string) {
	l.send(DebugLevel, content)
}
func (l Log) Info(content string) {
	l.send(InfoLevel, content)
}
func (l Log) Warn(content string) {
	l.send(WarnLevel, content)
}
func (l Log) Error(content string) {
	l.send(ErrorLevel, content)
}

func (l Log) send(level Level, content string) {
	err := global.DB.Create(&LogStashModel{
		IP:      l.Ip,
		Addr:    l.Addr,
		Level:   level,
		Content: content,
		UserID:  l.UserId,
	}).Error
	if err != nil {
		logrus.Error(err)
		return
	}
}
