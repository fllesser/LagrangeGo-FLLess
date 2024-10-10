package bot

import (
	"os"
	"time"

	"github.com/ExquisiteCore/LagrangeGo-Template/config"
	"github.com/LagrangeDev/LagrangeGo/client"
	"github.com/LagrangeDev/LagrangeGo/client/auth"
	"github.com/sirupsen/logrus"
)

// Bot 全局 Bot
type Bot struct {
	*client.QQClient
}

// Bot 实例

var QQClient *Bot

func Init() {
	appInfo := auth.AppList["linux"]["3.2.10-25765"]
	deviceInfo := auth.NewDeviceInfo(114514)
	qqClientInstance := client.NewClient(config.GlobalConfig.Bot.Account, appInfo, "https://sign.lagrangecore.org/api/sign/25765")
	//qqClientInstance.SetLogger(utils.Logger)
	qqClientInstance.UseDevice(deviceInfo)

	data, err := os.ReadFile("sig.bin")
	if err != nil {
		logrus.Warnln("read sig error:", err)
	} else {
		sig, err := auth.UnmarshalSigInfo(data, true)
		if err != nil {
			logrus.Warnln("load sig error:", err)
		} else {
			qqClientInstance.UseSig(sig)
		}
	}
	QQClient = &Bot{QQClient: qqClientInstance}
	CheckAlive()
}

// Login 登录
func Login() {
	// 声明 err 变量并进行错误处理
	err := QQClient.Login(config.GlobalConfig.Bot.Password, "qrcode.png")
	if err != nil {
		logrus.Errorln("login err:", err)
	}
}

// 保存sign

func Dumpsig() {
	data, err := QQClient.Sig().Marshal()
	if err != nil {
		logrus.Errorln("marshal sig.bin err:", err)
		return
	}
	err = os.WriteFile("sig.bin", data, 0644)
	if err != nil {
		logrus.Errorln("write sig.bin err:", err)
		return
	}
	logrus.Infoln("sig saved into sig.bin")
}

func CheckAlive() {
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				logrus.Errorf("定时器发生错误，%v", r)
			}
			ticker.Stop() // 意外退出时关闭定时器
		}()
		status, lastStatus := true, false
		statusContent := map[bool]string{true: "online", false: "offline"}
		for range ticker.C {
			if lastStatus != status {
				logrus.Errorf("Lgr[%v] %v", QQClient.Uin, statusContent[status])
			}
			lastStatus, status = status, QQClient.Online.Load()
		}
	}()
}
