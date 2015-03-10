package setting

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
	"github.com/beego/compress"
	"github.com/beego/i18n"
	//"github.com/beego/social-auth"
	//"github.com/beego/social-auth/apps"
	"fmt"
	"github.com/howeyc/fsnotify"
	"github.com/thanzen/eq/cachemanager"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const (
	APP_VER = "0.0.0.1"
)

const (
	LangEnUS = iota
	LangZhCN
)

//todo:add logic
var (
	Captcha   *captcha.Captcha
	CacheTime int64
)

//todo: add logic
var (
	Langs []string
)

//todo:add logic
var (
	GlobalConfPath   = "conf/app.conf"
	CompressConfPath = "conf/compress.json"
)

var (
	AppHost                   string
	AppUrl                    string
	AppLogo                   string
	EnforceRedirect           bool
	IsProMode                 bool
	AppVer                    string
	DateFormat                string
	DateTimeFormat            string
	DateTimeShortFormat       string
	SecretKey                 string
	PostgresConnection        string
	PostgresMigrateConnection string

	//password
	ActiveCodeLives   int
	ResetPwdCodeLives int

	//email
	MailHost     string
	MailFrom     string
	MailUser     string
	MailAuthUser string
	MailAuthPass string

	//login
	LoginRememberDays int
	LoginMaxRetries   int
	LoginFailedBlocks int

	//cookies
	CookieRememberName string
	CookieUsername     string

	AvatarUrl string

	//robot.txt
	RobotUas      string
	RobotDisallow string
)

func Initialize() {
	var err error

	var cac cache.Cache
	cac, err = cache.NewCache("memory", `{"interval":360}`)

	cachemanager.Cache = cac
	// session settings
	beego.SessionOn = true
	//beego.SessionProvider = Cfg.MustValue("session", "session_provider", "file")
	//beego.SessionSavePath = Cfg.MustValue("session", "session_path", "sessions")
	//beego.SessionName = Cfg.MustValue("session", "session_name", "wetalk_sess")
	//beego.SessionCookieLifeTime = Cfg.MustInt("session", "session_life_time", 0)
	//beego.SessionGCMaxLifetime = Cfg.MustInt64("session", "session_gc_time", 86400)
	//load custom configurations
	loadConfig()
	//disable live reload for db connections
	PostgresConnection = beego.AppConfig.DefaultString(beego.RunMode+"::"+"pg_conn", "user=postgres password=root dbname=eq sslmode=disable")
    PostgresMigrateConnection =  beego.AppConfig.DefaultString(beego.RunMode+"::"+"pg_migrate", "postgres://postgres:root@localhost:5432/test?sslmode=disable")
	// cache system

	Captcha = captcha.NewCaptcha("/captcha/", cac)
	Captcha.FieldIdName = "CaptchaId"
	Captcha.FieldCaptchaName = "Captcha"

	settingCompress()
	settingLocales()
	//watch conf files change
	configWatcher()
	if err != nil {
		beego.Error(err)
	}
}

func settingCompress() {
	setting, err := compress.LoadJsonConf(CompressConfPath, IsProMode, AppUrl)
	if err != nil {
		beego.Error(err)
		return
	}

	setting.RunCommand()

	if IsProMode {
		setting.RunCompress(true, false, true)
	}

	beego.AddFuncMap("compress_js", setting.Js.CompressJs)
	beego.AddFuncMap("compress_css", setting.Css.CompressCss)
}

func loadConfig() {
	AppVer = strings.Join(strings.Split(APP_VER, ".")[:3], ".")

	IsProMode = beego.RunMode == "pro"

	CacheTime = beego.AppConfig.DefaultInt64("cache_time", 300)
	cachemanager.ExpireTime = CacheTime
	AppHost = beego.HttpAddr + ":" + strconv.Itoa(beego.HttpPort)
	//	AppUrl = beego.AppConfig.DefaultString("app_url", "http://localhost:8080/")
	AppUrl = "http://" + AppHost + "/"
	beego.Info(PostgresConnection)
	EnforceRedirect = beego.AppConfig.DefaultBool("enforce_redirect", false)
	AppLogo = beego.AppConfig.DefaultString("app_logo", "/static/img/logo.gif")
	//todo change later
	MailFrom = beego.AppConfig.DefaultString("mail_from", "some@gmail.com")
	//todo change later
	MailUser = beego.AppConfig.DefaultString("mail_name", "tech company")
	// set mailer connect args
	MailHost = beego.AppConfig.DefaultString("mail_host", "127.0.0.1:25")
	MailAuthUser = beego.AppConfig.DefaultString("mail_user", "example@example.com")
	MailAuthPass = beego.AppConfig.DefaultString("mail_pass", "******")

	DateFormat = beego.AppConfig.String("date_format")
	DateTimeFormat = beego.AppConfig.String("datetime_format")
	DateTimeShortFormat = beego.AppConfig.String("datetime_short_format")

	ActiveCodeLives = beego.AppConfig.DefaultInt("acitve_code_live_minutes", 180)
	ResetPwdCodeLives = beego.AppConfig.DefaultInt("resetpwd_code_live_minutes", 180)

	LoginRememberDays = beego.AppConfig.DefaultInt("login_remember_days", 7)
	LoginMaxRetries = beego.AppConfig.DefaultInt("login_max_retries", 5)
	LoginFailedBlocks = beego.AppConfig.DefaultInt("login_failed_blocks", 10)

	CookieRememberName = beego.AppConfig.DefaultString("cookie_remember_name", "eq_magic")
	CookieUsername = beego.AppConfig.DefaultString("cookie_user_name", "eq_powerful")

	AvatarUrl = beego.AppConfig.DefaultString("avatar_url", "http://1.gravatar.com/avatar/")

	RobotUas = beego.AppConfig.DefaultString("robot::uas", "")
	RobotDisallow = beego.AppConfig.DefaultString("robot::disallow", "")

	SecretKey = beego.AppConfig.String("secret_key")
	if len(SecretKey) == 0 {
		fmt.Println("Please set your secret_key in app.conf file")
	}

}

func settingLocales() {
	// load locales with locale_LANG.ini files
	langs := "en-US|zh-CN"
	for _, lang := range strings.Split(langs, "|") {
		lang = strings.TrimSpace(lang)
		files := []string{"conf/" + "locale_" + lang + ".ini"}
		if fh, err := os.Open(files[0]); err == nil {
			fh.Close()
		} else {
			files = nil
		}
		if err := i18n.SetMessage(lang, "conf/global/"+"locale_"+lang+".ini", files...); err != nil {
			beego.Error("Fail to set message file: " + err.Error())
			os.Exit(2)
		}
	}
	Langs = i18n.ListLangs()
}

var eventTime = make(map[string]int64)

func configWatcher() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic("Failed start app watcher: " + err.Error())
	}

	go func() {
		for {
			select {
			case event := <-watcher.Event:
				switch filepath.Ext(event.Name) {
				case ".conf":
					if checkEventTime(event.Name) {
						continue
					}

					if err = beego.ParseConfig(); err != nil {
						beego.Error("Conf Reload: ", err)
					}
					loadConfig()
					if err := i18n.ReloadLangs(); err != nil {
						beego.Error("i18n Reload: ", err)
					}

					beego.Info("Config Reloaded")
					//todo:extend this functionality
				case ".json":
					if checkEventTime(event.Name) {
						continue
					}
					if event.Name == CompressConfPath {
						settingCompress()
						beego.Info("Beego Compress Reloaded")
					}
				}
			}
		}
	}()

	if err := watcher.WatchFlags("conf", fsnotify.FSN_MODIFY); err != nil {
		beego.Error(err)
	}

	if err := watcher.WatchFlags("conf/", fsnotify.FSN_MODIFY); err != nil {
		beego.Error(err)
	}
}

// checkEventTime returns true if FileModTime does not change.
func checkEventTime(name string) bool {
	mt := getFileModTime(name)
	if eventTime[name] == mt {
		return true
	}

	eventTime[name] = mt
	return false
}

// getFileModTime returns unix timestamp of `os.File.ModTime` by given path.
func getFileModTime(path string) int64 {
	path = strings.Replace(path, "\\", "/", -1)
	f, err := os.Open(path)
	if err != nil {
		beego.Error("Fail to open file[ %s ]\n", err)
		return time.Now().Unix()
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		beego.Error("Fail to get file information[ %s ]\n", err)
		return time.Now().Unix()
	}

	return fi.ModTime().Unix()
}

func IsMatchHost(uri string) bool {
	if len(uri) == 0 {
		return false
	}

	u, err := url.ParseRequestURI(uri)
	if err != nil {
		return false
	}

	if u.Host != beego.AppPath {
		return false
	}
	return true
}
