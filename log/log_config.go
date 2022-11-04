package log

type LogSimpleConfig struct {
	//输出到控制台
	Console bool `xml:"Console"`

	//输出彩色的字符
	Colored bool `xml:"Colored"`

	//存储路径
	//	eg: log/xxx.log
	Path string `xml:"Path"`

	//记录等级
	// trace < debug < info < warn(warning) < error
	Level string `xml:"Level"`

	//文件轮转类型
	//	none: 单日志文件
	//	size: 按文件大小
	//	hourly(hour): 按每小时
	//	daily(day): 按每天
	RotateType string `xml:"RotateType"`

	//存储文件最大数量
	//	0: 表示不限制文件数量
	MaxStoreFiles uint `xml:"MaxStoreFiles"`

	//存储文件最大字节数  仅 RotateType=size 时有效
	//	eg: 10m, 500k
	MaxFileSize string `xml:"MaxFileSize"`

	//启用压缩
	Compress bool `xml:"Compress"`

	//日志输出前缀
	Prefix string `xml:"Prefix"`
}
