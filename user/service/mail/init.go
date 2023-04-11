package send_mail

import (
	"github.com/bxcodec/go-clean-arch/domain"
)

type serviceMail struct{}

func NewMailService() domain.MailService {
	return &serviceMail{}
}
