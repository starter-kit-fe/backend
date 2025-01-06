package email

import (
	"bytes"
	"fmt"
	"html/template"
	"time"
	"admin/pkg/utils"

	"github.com/resend/resend-go/v2"
)

// Config 存储邮件服务配置
type Config struct {
	APIKey     string
	FromEmail  string
	Domain     string
	TimeFormat string
}

// Service 处理邮件发送逻辑
type Service struct {
	config Config
	client *resend.Client
	tmpl   *template.Template
}

// EmailData 存储邮件模板数据
type EmailData struct {
	Domain    string
	Code      string
	Timestamp string
}

// NewService 创建新的邮件服务实例
func NewService(cfg Config) (*Service, error) {
	if err := validateConfig(cfg); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	tmpl, err := template.New("verification").Parse(verificationEmailTemplate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse template: %w", err)
	}

	return &Service{
		config: cfg,
		client: resend.NewClient(cfg.APIKey),
		tmpl:   tmpl,
	}, nil
}

// SendVerificationCode 发送验证码邮件
func (s *Service) SendVerificationCode(to, code string) error {
	if err := validateEmail(to); err != nil {
		return fmt.Errorf("invalid recipient email: %w", err)
	}
	if err := validateCode(code); err != nil {
		return fmt.Errorf("invalid verification code: %w", err)
	}

	html, err := s.renderTemplate(code)
	if err != nil {
		return fmt.Errorf("failed to render template: %w", err)
	}

	params := &resend.SendEmailRequest{
		From:    s.config.FromEmail,
		To:      []string{to},
		Subject: fmt.Sprintf("%s 是您的验证码", code),
		Html:    html,
	}

	_, err = s.client.Emails.Send(params)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

// renderTemplate 渲染邮件模板
func (s *Service) renderTemplate(code string) (string, error) {
	data := EmailData{
		Domain:    s.config.Domain,
		Code:      code,
		Timestamp: time.Now().Format(s.config.TimeFormat),
	}

	var buf bytes.Buffer
	if err := s.tmpl.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// 验证配置
func validateConfig(cfg Config) error {
	if cfg.APIKey == "" {
		return fmt.Errorf("API key is required")
	}
	if cfg.FromEmail == "" {
		return fmt.Errorf("from email is required")
	}
	if cfg.Domain == "" {
		return fmt.Errorf("domain is required")
	}
	if cfg.TimeFormat == "" {
		return fmt.Errorf("time format is required")
	}
	return nil
}

// 验证邮箱格式
func validateEmail(email string) error {
	// 这里可以添加更复杂的邮箱验证逻辑
	if !utils.IsValidEmail(email) {
		return fmt.Errorf("email not valid")
	}
	return nil
}

// 验证验证码格式
func validateCode(code string) error {
	// 这里可以添加更复杂的验证码验证逻辑
	if code == "" {
		return fmt.Errorf("code cannot be empty")
	}
	return nil
}

// 验证码邮件模板
const verificationEmailTemplate = `
<table role="presentation" style="width:100%;max-width:405px;margin:0 auto;border-spacing:0">
    <tbody>
        <tr>
            <td>
                <p style="font-size:18px;margin:0;line-height:32px;text-align:center">
                    您的
                    <a href="https://{{.Domain}}" style="font-weight:bold">
                        {{.Domain}}
                    </a>
                    验证码
                </p>
                <p style="margin-top:20px;margin-bottom:0;background:#f6f6f6;height:40px;line-height:40px;font-size:20px;color:#000000;letter-spacing:5px;font-weight:bold;text-align:center">
                    {{.Code}}    
                </p>
            </td>
        </tr>
        <tr>
            <td>
                <p style="font-size:12px;margin-top:10px;margin-bottom:0;color:#13151a">
                    {{.Timestamp}}
                </p>
            </td>
        </tr>
        <tr>
            <td>
                <p style="font-size:12px;margin-top:20px;margin-bottom:0;color:#13151a">
                    您的验证码有效期为五分钟。
                </p>
            </td>
        </tr>
        <tr style="margin:0;padding:0;border:none">
            <td style="margin:0;padding:0;border:none">
                <p style="font-size:12px;border-bottom:1px solid #e6e8eb;margin-top:10px;padding-bottom:20px;color:#13151a;padding-top:0;margin-bottom:0">
                    不要把您的验证码告诉任何人
                </p>
            </td>
        </tr>
        <tr style="margin:0;padding:0;border:none">
            <td style="padding:0;margin:0;border:none">
                <div>
                    <p style="font-size:12px;padding-top:20px;line-height:18px;color:#848b96;margin:0;border:none">
                        不是您发送的验证码？投诉建议请联系
                        <a href="mailto:support@tigerzh.com"
                            style="text-decoration:none;color:#0059da;font-weight:bold"
                            target="_blank">
                            客服中心
                        </a>
                    </p>
                </div>
            </td>
        </tr>
    </tbody>
</table>`
