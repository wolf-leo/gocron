package notify

import (
	"github.com/ouqiang/gocron/internal/models"
	"github.com/ouqiang/gocron/internal/modules/logger"
	"github.com/ouqiang/gocron/internal/modules/utils"
	"strconv"
	"strings"
)

// @author qiang.ou<qingqianludao@gmail.com>
// @date 2017/5/1-00:19

type Mail struct {
}

func (mail *Mail) Send(msg Message) {
	model := new(models.Setting)
	mailSetting, err := model.Mail()
	logger.Debugf("%+v", mailSetting)
	if err != nil {
		logger.Error("#mail#从数据库获取mail配置失败", err)
		return
	}
	if mailSetting.Host == "" {
		logger.Error("#mail#Host为空")
		return
	}
	if mailSetting.Port == 0 {
		logger.Error("#mail#Port为空")
		return
	}
	if mailSetting.User == "" {
		logger.Error("#mail#User为空")
		return
	}
	if mailSetting.Password == "" {
		logger.Error("#mail#Password为空")
		return
	}
	msg["content"] = parseNotifyTemplate(mailSetting.Template, msg)
	toUsers := mail.getActiveMailUsers(mailSetting, msg)
	mail.send(mailSetting, toUsers, msg)
}

func (mail *Mail) send(mailSetting models.Mail, toUsers []string, msg Message) {
	body := msg["content"].(string)
	body = strings.Replace(body, "\n", "<br>", -1)

}

func (mail *Mail) getActiveMailUsers(mailSetting models.Mail, msg Message) []string {
	taskReceiverIds := strings.Split(msg["task_receiver_id"].(string), ",")
	users := []string{}
	for _, v := range mailSetting.MailUsers {
		if utils.InStringSlice(taskReceiverIds, strconv.Itoa(v.Id)) {
			users = append(users, v.Email)
		}
	}

	return users
}
