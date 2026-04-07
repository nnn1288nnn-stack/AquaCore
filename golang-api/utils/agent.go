package utils

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"
)

// AgentService - 本地 Agent 服务
type AgentService struct {
	OpenCli      *OpenCliClient
	cache        map[string]interface{}
	cacheMutex   sync.RWMutex
	sessionID    string
	isInitialized bool
}

// NewAgentService - 创建新的 Agent 服务
func NewAgentService(debug bool) *AgentService {
	return &AgentService{
		OpenCli:    NewOpenCliClient(debug),
		cache:      make(map[string]interface{}),
		sessionID:  generateSessionID(),
	}
}

// Initialize - 初始化 Agent 服务
func (s *AgentService) Initialize() error {
	// 检查 OpenCli 连接
	status, err := s.OpenCli.Doctor()
	if err != nil {
		return fmt.Errorf("❌ OpenCli 不可用: %v", err)
	}
	
	log.Printf("✅ OpenCli 状态: %s", status)
	s.isInitialized = true
	return nil
}

// NavigateAndExtract - 导航到 URL 并提取数据
func (s *AgentService) NavigateAndExtract(url string, extractors map[string]string) (map[string]interface{}, error) {
	if !s.isInitialized {
		return nil, fmt.Errorf("Agent 服务未初始化")
	}
	
	// 打开 URL
	if err := s.OpenCli.Open(url); err != nil {
		return nil, fmt.Errorf("打开 URL 失败: %v", err)
	}
	
	// 获取页面状态
	state, err := s.OpenCli.GetState()
	if err != nil {
		return nil, fmt.Errorf("获取页面状态失败: %v", err)
	}
	
	result := make(map[string]interface{})
	result["url"] = state.URL
	result["title"] = state.Title
	result["elements"] = state.Elements
	
	// 执行自定义提取器
	for key, selector := range extractors {
		if value, err := s.OpenCli.Eval(fmt.Sprintf(
			"(function(){ const el = document.querySelector('%s'); return el ? el.textContent : ''; })()",
			selector,
		)); err == nil {
			result[key] = value
		}
	}
	
	return result, nil
}

// ClickAndNavigate - 点击元素并导航
func (s *AgentService) ClickAndNavigate(elementIndex int) (map[string]interface{}, error) {
	if !s.isInitialized {
		return nil, fmt.Errorf("Agent 服务未初始化")
	}
	
	// 点击元素
	if err := s.OpenCli.Click(elementIndex); err != nil {
		return nil, fmt.Errorf("点击元素失败: %v", err)
	}
	
	// 等待页面加载
	s.OpenCli.WaitTime(1)
	
	// 获取新页面状态
	state, err := s.OpenCli.GetState()
	if err != nil {
		return nil, fmt.Errorf("获取页面状态失败: %v", err)
	}
	
	return map[string]interface{}{
		"url":      state.URL,
		"title":    state.Title,
		"elements": state.Elements,
	}, nil
}

// FillForm - 填充表单
func (s *AgentService) FillForm(fields map[int]string) error {
	if !s.isInitialized {
		return fmt.Errorf("Agent 服务未初始化")
	}
	
	for elementIndex, value := range fields {
		if err := s.OpenCli.Type(elementIndex, value); err != nil {
			return fmt.Errorf("填充表单失败 (元素 %d): %v", elementIndex, err)
		}
	}
	
	return nil
}

// SubmitForm - 提交表单
func (s *AgentService) SubmitForm(submitButtonIndex int) error {
	if !s.isInitialized {
		return fmt.Errorf("Agent 服务未初始化")
	}
	
	// 点击提交按钮
	if err := s.OpenCli.Click(submitButtonIndex); err != nil {
		return fmt.Errorf("提交表单失败: %v", err)
	}
	
	// 等待响应
	s.OpenCli.WaitTime(2)
	
	return nil
}

// ScrollPage - 滚动页面
func (s *AgentService) ScrollPage(direction string, amount int) error {
	if !s.isInitialized {
		return fmt.Errorf("Agent 服务未初始化")
	}
	
	return s.OpenCli.Scroll(direction, amount)
}

// ExtractTableData - 提取表格数据
func (s *AgentService) ExtractTableData(tableSelector string) ([]map[string]string, error) {
	if !s.isInitialized {
		return nil, fmt.Errorf("Agent 服务未初始化")
	}
	
	// 使用 JavaScript 提取表格数据
	jsCode := fmt.Sprintf(`
		(function() {
			const rows = document.querySelectorAll('%s tbody tr');
			const data = [];
			rows.forEach(row => {
				const cells = row.querySelectorAll('td');
				const rowData = {};
				cells.forEach((cell, idx) => {
					rowData['col_' + idx] = cell.textContent.trim();
				});
				data.push(rowData);
			});
			return JSON.stringify(data);
		})()
	`, tableSelector)
	
	output, err := s.OpenCli.Eval(jsCode)
	if err != nil {
		return nil, fmt.Errorf("提取表格数据失败: %v", err)
	}
	
	// 解析 JSON 结果
	var data []map[string]string
	if err := json.Unmarshal([]byte(output), &data); err != nil {
		return nil, fmt.Errorf("解析表格数据失败: %v", err)
	}
	
	return data, nil
}

// SetCache - 设置缓存
func (s *AgentService) SetCache(key string, value interface{}) {
	s.cacheMutex.Lock()
	defer s.cacheMutex.Unlock()
	s.cache[key] = value
}

// GetCache - 获取缓存
func (s *AgentService) GetCache(key string) (interface{}, bool) {
	s.cacheMutex.RLock()
	defer s.cacheMutex.RUnlock()
	val, ok := s.cache[key]
	return val, ok
}

// ClearCache - 清除缓存
func (s *AgentService) ClearCache() {
	s.cacheMutex.Lock()
	defer s.cacheMutex.Unlock()
	s.cache = make(map[string]interface{})
}

// GetSessionID - 获取会话 ID
func (s *AgentService) GetSessionID() string {
	return s.sessionID
}

// LogActivity - 记录活动
func (s *AgentService) LogActivity(action string, details map[string]interface{}) {
	log.Printf("[AGENT] 会话 %s - 操作: %s, 详情: %+v", s.sessionID, action, details)
}

// 辅助函数

func generateSessionID() string {
	b := make([]byte, 8)
	rand.Read(b)
	return hex.EncodeToString(b) + "-" + fmt.Sprintf("%d", time.Now().Unix())
}
