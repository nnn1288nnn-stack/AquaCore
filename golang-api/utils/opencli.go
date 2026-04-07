package opencli_sdk

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

// OpenCliClient - OpenCli 操作客户端
type OpenCliClient struct {
	DEBUG bool
}

// NewOpenCliClient - 创建一个新的 OpenCli 客户端
func NewOpenCliClient(debug bool) *OpenCliClient {
	return &OpenCliClient{DEBUG: debug}
}

// PageState - 页面状态响应
type PageState struct {
	URL      string                     `json:"url"`
	Title    string                     `json:"title"`
	Elements map[int]PageElement        `json:"elements"`
	RawHTML  string                     `json:"rawHtml,omitempty"`
}

// PageElement - 页面元素定义
type PageElement struct {
	Index    int    `json:"index"`
	Tag      string `json:"tag"`
	Text     string `json:"text"`
	Value    string `json:"value"`
	Selector string `json:"selector"`
	Type     string `json:"type"`
}

// OperateResult - 操作结果
type OperateResult struct {
	Success bool        `json:"success"`
	Result  interface{} `json:"result,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// executeCommand - 执行 opencli 命令
func (c *OpenCliClient) executeCommand(args ...string) (string, error) {
	cmd := exec.Command("opencli", args...)
	output, err := cmd.CombinedOutput()
	
	if c.DEBUG {
		fmt.Printf("[OPENCLI DEBUG] 命令: opencli %s\n", strings.Join(args, " "))
		fmt.Printf("[OPENCLI DEBUG] 输出: %s\n", string(output))
		if err != nil {
			fmt.Printf("[OPENCLI DEBUG] 错误: %v\n", err)
		}
	}
	
	if err != nil {
		return "", fmt.Errorf("opencli 命令失败: %v, 输出: %s", err, string(output))
	}
	
	return string(output), nil
}

// Open - 打开 URL
func (c *OpenCliClient) Open(url string) error {
	_, err := c.executeCommand("operate", "open", url)
	return err
}

// GetState - 获取当前页面状态
func (c *OpenCliClient) GetState() (*PageState, error) {
	output, err := c.executeCommand("operate", "state")
	if err != nil {
		return nil, err
	}
	
	var state PageState
	if err := json.Unmarshal([]byte(output), &state); err != nil {
		// 如果 JSON 解析失败，返回原始文本
		return &PageState{RawHTML: output}, nil
	}
	
	return &state, nil
}

// Click - 点击元素
func (c *OpenCliClient) Click(elementIndex int) error {
	_, err := c.executeCommand("operate", "click", fmt.Sprintf("%d", elementIndex))
	return err
}

// Type - 输入文本
func (c *OpenCliClient) Type(elementIndex int, text string) error {
	_, err := c.executeCommand("operate", "type", fmt.Sprintf("%d", elementIndex), text)
	return err
}

// Select - 选择下拉选项
func (c *OpenCliClient) Select(elementIndex int, value string) error {
	_, err := c.executeCommand("operate", "select", fmt.Sprintf("%d", elementIndex), value)
	return err
}

// GetValue - 获取元素值
func (c *OpenCliClient) GetValue(elementIndex int) (string, error) {
	output, err := c.executeCommand("operate", "get", "value", fmt.Sprintf("%d", elementIndex))
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(output), nil
}

// GetText - 获取元素文本
func (c *OpenCliClient) GetText(elementIndex int) (string, error) {
	output, err := c.executeCommand("operate", "get", "text", fmt.Sprintf("%d", elementIndex))
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(output), nil
}

// GetTitle - 获取页面标题
func (c *OpenCliClient) GetTitle() (string, error) {
	output, err := c.executeCommand("operate", "get", "title")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(output), nil
}

// GetURL - 获取当前 URL
func (c *OpenCliClient) GetURL() (string, error) {
	output, err := c.executeCommand("operate", "get", "url")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(output), nil
}

// Scroll - 滚动页面
func (c *OpenCliClient) Scroll(direction string, amount int) error {
	args := []string{"operate", "scroll", direction}
	if amount > 0 {
		args = append(args, "--amount", fmt.Sprintf("%d", amount))
	}
	_, err := c.executeCommand(args...)
	return err
}

// Wait - 等待条件满足
func (c *OpenCliClient) Wait(waitType string, value string) error {
	_, err := c.executeCommand("operate", "wait", waitType, value)
	return err
}

// WaitTime - 等待指定时间（秒）
func (c *OpenCliClient) WaitTime(seconds int) error {
	_, err := c.executeCommand("operate", "wait", "time", fmt.Sprintf("%d", seconds))
	return err
}

// Back - 返回上一页
func (c *OpenCliClient) Back() error {
	_, err := c.executeCommand("operate", "back")
	return err
}

// Screenshot - 保存屏幕截图
func (c *OpenCliClient) Screenshot(filePath string) error {
	_, err := c.executeCommand("operate", "screenshot", filePath)
	return err
}

// Eval - 执行 JavaScript 代码 (仅用于读取)
func (c *OpenCliClient) Eval(jsCode string) (string, error) {
	output, err := c.executeCommand("operate", "eval", jsCode)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(output), nil
}

// Keys - 按下键盘按键
func (c *OpenCliClient) Keys(keys string) error {
	_, err := c.executeCommand("operate", "keys", keys)
	return err
}

// Doctor - 诊断 OpenCli 连接状态
func (c *OpenCliClient) Doctor() (string, error) {
	return c.executeCommand("doctor")
}
