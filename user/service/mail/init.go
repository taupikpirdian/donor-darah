package send_mail

import (
	"github.com/bxcodec/go-clean-arch/cfg"
	"github.com/bxcodec/go-clean-arch/domain"
)

type serviceMail struct {
	cfg cfg.Config
}

func NewMailService(cfg cfg.Config) domain.MailService {
	return &serviceMail{cfg: cfg}
}
