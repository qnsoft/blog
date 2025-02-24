package conf

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/zxysilent/logs"
)

type appcfg struct {
	Title        string `toml:"title" json:"title"`                   //
	Intro        string `toml:"intro" json:"intro"`                   //
	Mode         string `toml:"mode" json:"mode"`                     //
	Addr         string `toml:"addr" json:"addr"`                     //
	Srv          string `toml:"srv" json:"srv"`                       //
	TokenKey     string `toml:"token_key" json:"token_key"`           //token关键词
	TokenExp     int    `toml:"token_exp" json:"token_exp"`           //过期时间 h
	TokenKeep    bool   `toml:"token_keep" json:"token_keep"`         //保持在线
	TokenSso     bool   `toml:"token_sso" json:"token_sso"`           //单点登录
	TokenSecret  string `toml:"token_secret" json:"token_secret"`     //加密私钥
	ImageCut     bool   `toml:"image_cut" json:"image_cut"`           //图片裁剪
	ImageWidth   int    `toml:"image_width" json:"image_width"`       //图片宽度
	ImageHeight  int    `toml:"image_height" json:"image_height"`     //图片高度
	PageMin      int    `toml:"page_min" json:"page_min"`             //最小分页大小
	PageMax      int    `toml:"page_max" json:"page_max"`             //最大分页大小
	DbHost       string `toml:"db_host" json:"db_host"`               //数据库地址
	DbPort       int    `toml:"db_port" json:"db_port"`               //数据库端口
	DbUser       string `toml:"db_user" json:"db_user"`               //数据库账号
	DbPasswd     string `toml:"db_passwd" json:"db_passwd"`           //数据库密码
	DbName       string `toml:"db_name" json:"db_name"`               //数据库名称
	DbParams     string `toml:"db_params" json:"db_params"`           //数据库参数
	OrmIdle      int    `toml:"orm_idle" json:"orm_idle"`             //
	OrmOpen      int    `toml:"orm_open" json:"orm_open"`             //
	OrmShow      bool   `toml:"orm_show" json:"orm_show"`             //显示sql
	OrmSync      bool   `toml:"orm_sync" json:"orm_sync"`             //同步表结构
	OrmCacheUse  bool   `toml:"orm_cache_use" json:"orm_cache_use"`   //是否使用缓存
	OrmCacheSize int    `toml:"orm_cache_size" json:"orm_cache_size"` //缓存数量
	OrmHijackLog bool   `toml:"orm_hijack_log" json:"orm_hijack_log"` //劫持日志
	Author       struct {
		Name    string `toml:"name" json:"name"`
		Website string `toml:"website" json:"website"`
	} `toml:"author" json:"author"`
	Wechat struct {
		GzhAppid  string `toml:"gzh_appid"`  //公众号
		GzhSecret string `toml:"gzh_secret"` //公众号
		MpgAppid  string `toml:"mpg_appid"`  //小程序
		MpgSecret string `toml:"mpg_secret"` //小程序
		WebAppid  string `toml:"web_appid"`  //web
	} `toml:"wechat" json:"wechat"`
}

func (app *appcfg) IsProd() bool {
	return app.Mode == "prod"
}
func (app *appcfg) IsDev() bool {
	return app.Mode == "dev"
}

// uid:pass@tcp(host:port)/dbname?charset=utf8&parseTime=true
// 用户名:密码@tcp(主机:端口)/数据库名称?charset=utf8&parseTime=true
const _dsn = "%s:%s@tcp(%s:%d)/%s?%s"

// MySQL链接字符串
func (app *appcfg) Dsn() string {
	return fmt.Sprintf(_dsn, app.DbUser, app.DbPasswd, app.DbHost, app.DbPort, app.DbName, app.DbParams)
}

var (
	App       *appcfg              //运行配置实体
	defConfig = "./conf/conf.toml" //配置文件路径，方便测试
)

func Init() {
	var err error
	App, err = initCfg()
	if err != nil {
		logs.Fatal("config init error : ", err.Error())
	}
	logs.Debug("conf init")
}

func initCfg() (*appcfg, error) {
	app := &appcfg{}
	_, err := toml.DecodeFile(defConfig, &app)
	if err != nil {
		return nil, err
	}
	return app, nil
}
