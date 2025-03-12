package email

import (
	"admin/pkg/utils"
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
	"time"

	"github.com/resend/resend-go/v2"
)

//go:embed verification_email_template.html
var verificationEmailTemplate string

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
