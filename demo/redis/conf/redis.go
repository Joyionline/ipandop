package conf

type RdsConfig struct {
	Host      string
	Port      int
	User      string
	Pwd       string
	IsRunning bool // 是否正常运行
}

var RdsCacheList = []RdsConfig{
	{
		Host:      "127.0.0.1",
		Port:      6379,
		User:      "",
		Pwd:       "",
		IsRunning: true,
	},
}

var RdsCache RdsConfig = RdsCacheList[0]
