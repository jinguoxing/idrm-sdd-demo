package utils

// SMSProvider 短信服务提供者接口
// 当前为接口定义，具体实现待后续集成真实短信服务（阿里云/腾讯云等）
type SMSProvider interface {
	// SendCode 发送验证码
	// phone: 手机号
	// code: 验证码（6位数字）
	// codeType: 验证码类型（register/reset）
	SendCode(phone, code, codeType string) error
}

// MockSMSProvider Mock短信服务提供者（用于开发和测试）
type MockSMSProvider struct{}

// NewMockSMSProvider 创建Mock短信服务提供者
func NewMockSMSProvider() SMSProvider {
	return &MockSMSProvider{}
}

// SendCode 发送验证码（Mock实现，仅打印日志）
func (m *MockSMSProvider) SendCode(phone, code, codeType string) error {
	// TODO: 后续集成真实短信服务时，替换此实现
	// 当前Mock实现仅用于开发测试
	// logx.Infof("Mock SMS: 发送验证码到 %s, 验证码: %s, 类型: %s", phone, code, codeType)
	return nil
}

