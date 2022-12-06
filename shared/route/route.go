package route

const (
	Register = 1 // 用户注册
	Login    = 2 // 用户登录
)

const (
	NewMail       = 101 // 新邮件
	FetchMailList = 102 // 拉取邮件列表
	ReadMail      = 103 // 读取邮件
	ReadAllMail   = 104 // 一键读取所有邮件
	DeleteMail    = 105 // 删除邮件
	DeleteAllMail = 106 // 删除所有邮件
)

const (
	FetchRooms  = 1000 // 拉取房间
	QuickStart  = 1001 // 快速开始
	SitDown     = 1002 // 坐下
	StandUp     = 1003 // 站起
	StartReady  = 1004 // 开始准备
	CancelReady = 1005 // 取消准备
	Offline     = 1006 // 玩家离线
	Online      = 1007 // 玩家上线
)
