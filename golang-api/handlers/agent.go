package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	opencli "github.com/penghu-digital-captain/golang-api/opencli"
	"github.com/penghu-digital-captain/golang-api/utils"
)

// AgentHandler - Agent 处理器
type AgentHandler struct {
	agentService *utils.AgentService
}

// NewAgentHandler - 创建新的 Agent 处理器
func NewAgentHandler(agentService *utils.AgentService) *AgentHandler {
	return &AgentHandler{
		agentService: agentService,
	}
}

// NavigateRequest - 导航请求
type NavigateRequest struct {
	URL        string                 `json:"url" binding:"required"`
	Extractors map[string]string      `json:"extractors,omitempty"`
}

// NavigateResponse - 导航响应
type NavigateResponse struct {
	Success bool                   `json:"success"`
	Data    map[string]interface{} `json:"data,omitempty"`
	Error   string                 `json:"error,omitempty"`
}

// Navigate - 导航端点
func (h *AgentHandler) Navigate(c *gin.Context) {
	var req NavigateRequest
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, NavigateResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}
	
	data, err := h.agentService.NavigateAndExtract(req.URL, req.Extractors)
	if err != nil {
		c.JSON(http.StatusInternalServerError, NavigateResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, NavigateResponse{
		Success: true,
		Data:    data,
	})
}

// ClickRequest - 点击请求
type ClickRequest struct {
	ElementIndex int `json:"element_index" binding:"required"`
}

// Click - 点击端点
func (h *AgentHandler) Click(c *gin.Context) {
	var req ClickRequest
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, NavigateResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}
	
	data, err := h.agentService.ClickAndNavigate(req.ElementIndex)
	if err != nil {
		c.JSON(http.StatusInternalServerError, NavigateResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, NavigateResponse{
		Success: true,
		Data:    data,
	})
}

// FillFormRequest - 填充表单请求
type FillFormRequest struct {
	Fields map[string]string `json:"fields" binding:"required"`
}

// FillForm - 填充表单端点
func (h *AgentHandler) FillForm(c *gin.Context) {
	var req FillFormRequest
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, NavigateResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}
	
	// 转换字符串 key 为 int
	fields := make(map[int]string)
	for k, v := range req.Fields {
		var i int
		if _, err := scanInt(k, &i); err != nil {
			c.JSON(http.StatusBadRequest, NavigateResponse{
				Success: false,
				Error:   "字段 key 必须是整数",
			})
			return
		}
		fields[i] = v
	}
	
	if err := h.agentService.FillForm(fields); err != nil {
		c.JSON(http.StatusInternalServerError, NavigateResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, NavigateResponse{
		Success: true,
		Data:    map[string]interface{}{"message": "表单已填充"},
	})
}

// SubmitFormRequest - 提交表单请求
type SubmitFormRequest struct {
	SubmitButtonIndex int `json:"submit_button_index" binding:"required"`
}

// SubmitForm - 提交表单端点
func (h *AgentHandler) SubmitForm(c *gin.Context) {
	var req SubmitFormRequest
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, NavigateResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}
	
	if err := h.agentService.SubmitForm(req.SubmitButtonIndex); err != nil {
		c.JSON(http.StatusInternalServerError, NavigateResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, NavigateResponse{
		Success: true,
		Data:    map[string]interface{}{"message": "表单已提交"},
	})
}

// GetPageStateResponse - 获取页面状态响应
type GetPageStateResponse struct {
	Success bool                   `json:"success"`
	Data    *opencli.PageState     `json:"data,omitempty"`
	Error   string                 `json:"error,omitempty"`
}

// GetPageState - 获取页面状态端点
func (h *AgentHandler) GetPageState(c *gin.Context) {
	state, err := h.agentService.OpenCli.GetState()
	if err != nil {
		c.JSON(http.StatusInternalServerError, GetPageStateResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, GetPageStateResponse{
		Success: true,
		Data:    state,
	})
}

// GetSessionResponse - 获取会话响应
type GetSessionResponse struct {
	SessionID string `json:"session_id"`
	Status    string `json:"status"`
}

// GetSession - 获取会话信息端点
func (h *AgentHandler) GetSession(c *gin.Context) {
	c.JSON(http.StatusOK, GetSessionResponse{
		SessionID: h.agentService.GetSessionID(),
		Status:    "initialied",
	})
}

// 辅助函数

func scanInt(s string, i *int) (int, error) {
	_, err := fmt.Sscanf(s, "%d", i)
	return *i, err
}
