package main

import (
	"fmt"
	"log"
	"time"

	"project/golang-api/utils"
)

// 示例 1: 简单导航和提取
func example1SimpleNavigation() {
	fmt.Println("\n" + "="*60)
	fmt.Println("示例 1: 简单导航和提取")
	fmt.Println("="*60)

	// 创建 OpenCli 客户端
	client := utils.NewOpenCliClient()

	// 打开网页
	if err := client.Open("https://example.com"); err != nil {
		log.Printf("❌ 打开页面失败: %v", err)
		return
	}

	// 获取页面状态
	state, err := client.GetState()
	if err != nil {
		log.Printf("❌ 获取页面状态失败: %v", err)
		return
	}

	fmt.Printf("✅ 页面加载成功\n")
	fmt.Printf("   标题: %s\n", state.Title)
	fmt.Printf("   URL: %s\n", state.URL)
	fmt.Printf("   元素数: %d\n", len(state.Elements))
}

// 示例 2: 表单填充
func example2FormFilling() {
	fmt.Println("\n" + "="*60)
	fmt.Println("示例 2: 表单填充")
	fmt.Println("="*60)

	client := utils.NewOpenCliClient()

	// 打开登录页面
	if err := client.Open("https://example.com/login"); err != nil {
		log.Printf("❌ 打开页面失败: %v", err)
		return
	}

	time.Sleep(1 * time.Second)

	// 获取页面状态
	state, err := client.GetState()
	if err != nil {
		log.Printf("❌ 获取页面状态失败: %v", err)
		return
	}

	// 显示可用的表单字段
	fmt.Printf("✅ 页面加载成功\n")
	fmt.Printf("   找到 %d 个元素\n", len(state.Elements))

	// 填充邮箱字段 (假设索引为 1)
	if err := client.Type(1, "user@example.com"); err != nil {
		log.Printf("❌ 输入邮箱失败: %v", err)
		return
	}
	fmt.Println("   ✓ 邮箱已填充")

	// 填充密码字段 (假设索引为 2)
	if err := client.Type(2, "password123"); err != nil {
		log.Printf("❌ 输入密码失败: %v", err)
		return
	}
	fmt.Println("   ✓ 密码已填充")
}

// 示例 3: 点击和导航
func example3ClickNavigation() {
	fmt.Println("\n" + "="*60)
	fmt.Println("示例 3: 点击和导航")
	fmt.Println("="*60)

	client := utils.NewOpenCliClient()

	// 打开页面
	if err := client.Open("https://example.com"); err != nil {
		log.Printf("❌ 打开页面失败: %v", err)
		return
	}

	time.Sleep(1 * time.Second)

	// 获取初始状态
	state1, err := client.GetState()
	if err != nil {
		log.Printf("❌ 获取页面状态失败: %v", err)
		return
	}

	fmt.Printf("✅ 初始页面\n")
	fmt.Printf("   URL: %s\n", state1.URL)

	// 查找按钮元素并点击
	// (假设有一个"Next"按钮在索引 5 处)
	if err := client.Click(5); err != nil {
		log.Printf("❌ 点击按钮失败: %v", err)
		return
	}

	time.Sleep(2 * time.Second)

	// 获取新页面状态
	state2, err := client.GetState()
	if err != nil {
		log.Printf("❌ 获取页面状态失败: %v", err)
		return
	}

	fmt.Printf("✅ 导航后页面\n")
	fmt.Printf("   URL: %s\n", state2.URL)
	fmt.Printf("   状态已改变: %v\n", state1.URL != state2.URL)
}

// 示例 4: 使用 Agent 服务
func example4AgentService() {
	fmt.Println("\n" + "="*60)
	fmt.Println("示例 4: 使用 Agent 服务")
	fmt.Println("="*60)

	// 创建 Agent 服务
	service := utils.NewAgentService()

	// 初始化服务
	if err := service.Initialize(); err != nil {
		log.Printf("❌ Agent 初始化失败: %v", err)
		return
	}

	fmt.Println("✅ Agent 初始化成功")
	fmt.Printf("   Session ID: %s\n", service.SessionID)

	// 导航并提取数据
	extractors := map[string]string{
		"title": "h1",
		"body":  "p",
	}

	result, err := service.NavigateAndExtract("https://example.com", extractors)
	if err != nil {
		log.Printf("❌ 导航和提取失败: %v", err)
		return
	}

	fmt.Printf("✅ 导航和提取成功\n")
	fmt.Printf("   提取的数据: %v\n", result)
}

// 示例 5: 滚动页面
func example5Scrolling() {
	fmt.Println("\n" + "="*60)
	fmt.Println("示例 5: 滚动页面")
	fmt.Println("="*60)

	client := utils.NewOpenCliClient()

	// 打开长页面
	if err := client.Open("https://example.com/long-page"); err != nil {
		log.Printf("❌ 打开页面失败: %v", err)
		return
	}

	time.Sleep(1 * time.Second)

	// 向下滚动
	fmt.Println("✅ 开始滚动...")

	for i := 1; i <= 3; i++ {
		if err := client.Scroll("down"); err != nil {
			log.Printf("❌ 滚动失败: %v", err)
			return
		}
		fmt.Printf("   - 滚动 %d 次\n", i)
		time.Sleep(500 * time.Millisecond)
	}

	// 获取最终状态
	state, err := client.GetState()
	if err != nil {
		log.Printf("❌ 获取页面状态失败: %v", err)
		return
	}

	fmt.Printf("✅ 滚动完成，加载了 %d 个元素\n", len(state.Elements))
}

// 示例 6: 等待条件
func example6WaitFor() {
	fmt.Println("\n" + "="*60)
	fmt.Println("示例 6: 等待条件")
	fmt.Println("="*60)

	client := utils.NewOpenCliClient()

	// 打开页面
	if err := client.Open("https://example.com"); err != nil {
		log.Printf("❌ 打开页面失败: %v", err)
		return
	}

	fmt.Println("✅ 等待页面加载...")

	// 等待指定时间
	if err := client.Wait("time", "2"); err != nil {
		log.Printf("❌ 等待失败: %v", err)
		return
	}

	fmt.Println("✅ 等待完成")

	// 等待特定选择器出现
	if err := client.Wait("selector", ".product-item"); err != nil {
		log.Printf("❌ 等待选择器失败: %v", err)
		return
	}

	fmt.Println("✅ 产品项已加载")
}

// 示例 7: 使用 JavaScript 求值
func example7JavaScript() {
	fmt.Println("\n" + "="*60)
	fmt.Println("示例 7: JavaScript 求值")
	fmt.Println("="*60)

	client := utils.NewOpenCliClient()

	// 打开页面
	if err := client.Open("https://example.com/products"); err != nil {
		log.Printf("❌ 打开页面失败: %v", err)
		return
	}

	time.Sleep(1 * time.Second)

	// 使用 JavaScript 获取所有产品名称
	jsCode := `
(function() {
    const products = document.querySelectorAll('.product-name');
    const names = [];
    products.forEach(p => names.push(p.textContent.trim()));
    return JSON.stringify(names);
})()
`

	result, err := client.Eval(jsCode)
	if err != nil {
		log.Printf("❌ JavaScript 执行失败: %v", err)
		return
	}

	fmt.Printf("✅ JavaScript 执行成功\n")
	fmt.Printf("   结果: %s\n", result)
}

// 示例 8: 多步工作流
func example8Workflow() {
	fmt.Println("\n" + "="*60)
	fmt.Println("示例 8: 多步工作流")
	fmt.Println("="*60)

	client := utils.NewOpenCliClient()

	steps := []string{
		"打开登录页面",
		"填充用户名",
		"填充密码",
		"提交表单",
		"验证登录",
	}

	// 步骤 1: 打开页面
	fmt.Printf("步骤 1: %s\n", steps[0])
	if err := client.Open("https://example.com/login"); err != nil {
		log.Printf("❌ 失败: %v", err)
		return
	}

	time.Sleep(1 * time.Second)

	// 步骤 2-4: 表单交互
	fmt.Printf("步骤 2: %s\n", steps[1])
	if err := client.Type(1, "admin"); err != nil {
		log.Printf("❌ 失败: %v", err)
		return
	}

	fmt.Printf("步骤 3: %s\n", steps[2])
	if err := client.Type(2, "password123"); err != nil {
		log.Printf("❌ 失败: %v", err)
		return
	}

	fmt.Printf("步骤 4: %s\n", steps[3])
	if err := client.Click(3); err != nil {
		log.Printf("❌ 失败: %v", err)
		return
	}

	time.Sleep(2 * time.Second)

	// 步骤 5: 验证
	fmt.Printf("步骤 5: %s\n", steps[4])
	state, err := client.GetState()
	if err != nil {
		log.Printf("❌ 失败: %v", err)
		return
	}

	fmt.Printf("✅ 工作流完成\n")
	fmt.Printf("   当前 URL: %s\n", state.URL)
}

func main() {
	fmt.Println(`
╔════════════════════════════════════════════════════════════════════════════╗
║           澎湖數位老船長 - Go OpenCli 使用示例                             ║
║         Penghu Digital Captain - Go OpenCli Examples                       ║
╚════════════════════════════════════════════════════════════════════════════╝
	`)

	examples := map[string]func(){
		"1": example1SimpleNavigation,
		"2": example2FormFilling,
		"3": example3ClickNavigation,
		"4": example4AgentService,
		"5": example5Scrolling,
		"6": example6WaitFor,
		"7": example7JavaScript,
		"8": example8Workflow,
	}

	fmt.Println("选择要运行的示例:")
	fmt.Println("  1. 简单导航和提取")
	fmt.Println("  2. 表单填充")
	fmt.Println("  3. 点击和导航")
	fmt.Println("  4. 使用 Agent 服务")
	fmt.Println("  5. 滚动页面")
	fmt.Println("  6. 等待条件")
	fmt.Println("  7. JavaScript 求值")
	fmt.Println("  8. 多步工作流")

	var choice string
	fmt.Print("\n请输入选择 (1-8): ")
	fmt.Scanln(&choice)

	if fn, ok := examples[choice]; ok {
		fn()
	} else {
		fmt.Println("❌ 无效的选择")
	}
}
