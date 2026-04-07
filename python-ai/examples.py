"""
澎湖數位老船長 - 本地 Agent 使用示例

展示如何使用 OpenCli Agent 执行各种自动化任务
"""

import logging
from utils.local_agent import LocalAgentService

# 配置日志
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)


def example_1_simple_navigation():
    """示例 1: 简单导航和提取文本"""
    logger.info("=" * 60)
    logger.info("示例 1: 简单导航和提取文本")
    logger.info("=" * 60)
    
    agent = LocalAgentService(debug=True)
    
    # 导航到网站并提取信息
    result = agent.navigate_and_extract(
        url="https://example.com",
        extractors={
            "title": "h1",
            "body_text": "p"
        }
    )
    
    if result["success"]:
        print(f"✅ 导航成功")
        print(f"标题: {result['title']}")
        print(f"提取的数据: {result.get('extracted', {})}")
    else:
        print(f"❌ 导航失败: {result['error']}")


def example_2_form_filling():
    """示例 2: 表单填充和提交"""
    logger.info("\n" + "=" * 60)
    logger.info("示例 2: 表单填充和提交")
    logger.info("=" * 60)
    
    agent = LocalAgentService(debug=True)
    
    # 导航到表单页面
    agent.opencli.open("https://example.com/form")
    agent.opencli.wait("time", "1")
    
    # 获取页面状态以找到表单字段索引
    state = agent.opencli.get_state()
    print(f"页面元素: {state.elements}")
    
    # 填充表单字段
    # 假设: 字段 [3] = 邮箱, [4] = 密码, [5] = 提交按钮
    result = agent.fill_and_submit_form(
        fields={
            3: "user@example.com",
            4: "secure_password_123"
        },
        submit_index=5
    )
    
    if result["success"]:
        print(f"✅ {result['message']}")
    else:
        print(f"❌ 提交失败: {result['error']}")


def example_3_table_extraction():
    """示例 3: 表格数据提取"""
    logger.info("\n" + "=" * 60)
    logger.info("示例 3: 表格数据提取")
    logger.info("=" * 60)
    
    agent = LocalAgentService(debug=True)
    
    # 导航到包含表格的页面
    agent.opencli.open("https://example.com/data-table")
    agent.opencli.wait("time", "1")
    
    # 提取表格数据
    result = agent.extract_table("table.data-table")
    
    if result["success"]:
        print(f"✅ 提取了 {result['row_count']} 行数据\n")
        
        # 显示前 3 行
        for i, row in enumerate(result["data"][:3]):
            print(f"行 {i+1}: {row}")
    else:
        print(f"❌ 提取失败: {result['error']}")


def example_4_multi_step_workflow():
    """示例 4: 多步工作流 - 登录 -> 导航 -> 操作"""
    logger.info("\n" + "=" * 60)
    logger.info("示例 4: 多步工作流")
    logger.info("=" * 60)
    
    agent = LocalAgentService(debug=True)
    
    try:
        # 步骤 1: 打开登录页面
        print("步骤 1: 打开登录页面...")
        agent.opencli.open("https://example.com/login")
        agent.opencli.wait("time", "1")
        
        # 步骤 2: 获取页面状态
        print("步骤 2: 获取页面状态...")
        state = agent.opencli.get_state()
        print(f"  - 找到 {len(state.elements)} 个元素")
        
        # 步骤 3: 填充登录表单
        print("步骤 3: 填充登录信息...")
        agent.opencli.type_text(3, "admin@example.com")
        agent.opencli.type_text(4, "password123")
        
        # 步骤 4: 提交登录
        print("步骤 4: 提交登录...")
        agent.opencli.click(5)
        agent.opencli.wait("time", "2")
        
        # 步骤 5: 验证登录成功
        print("步骤 5: 验证登录...")
        new_state = agent.opencli.get_state()
        print(f"  - 当前 URL: {new_state.url}")
        print(f"  - 页面标题: {new_state.title}")
        
        print("✅ 工作流完成！")
        
    except Exception as e:
        print(f"❌ 工作流失败: {e}")


def example_5_dynamic_scrolling():
    """示例 5: 动态滚动和加载"""
    logger.info("\n" + "=" * 60)
    logger.info("示例 5: 动态滚动和加载")
    logger.info("=" * 60)
    
    agent = LocalAgentService(debug=True)
    
    try:
        # 打开长页面
        print("打开页面...")
        agent.opencli.open("https://example.com/products")
        agent.opencli.wait("time", "1")
        
        # 滚动加载更多内容
        print("滚动页面...")
        for i in range(3):
            print(f"  - 滚动 {i+1}...")
            agent.opencli.scroll("down", 500)
            agent.opencli.wait("time", "1")
        
        # 获取最终页面状态
        state = agent.opencli.get_state()
        print(f"✅ 滚动完成，加载了 {len(state.elements)} 个元素")
        
    except Exception as e:
        print(f"❌ 滚动失败: {e}")


def example_6_conditional_navigation():
    """示例 6: 条件导航 (如果元素存在则点击)"""
    logger.info("\n" + "=" * 60)
    logger.info("示例 6: 条件导航")
    logger.info("=" * 60)
    
    agent = LocalAgentService(debug=True)
    
    try:
        # 打开页面
        print("打开页面...")
        agent.opencli.open("https://example.com")
        agent.opencli.wait("time", "1")
        
        # 获取页面状态检查按钮是否存在
        state = agent.opencli.get_state()
        
        # 查找按钮 (假设按钮文本包含 "Next")
        next_button_index = None
        for idx, element in state.elements.items():
            if "Next" in element.get("text", ""):
                next_button_index = idx
                break
        
        if next_button_index:
            print(f"✅ 找到 'Next' 按钮 (索引: {next_button_index})")
            print(f"点击按钮...")
            agent.opencli.click(next_button_index)
            agent.opencli.wait("time", "1")
            print("✅ 导航成功")
        else:
            print("❌ 未找到 'Next' 按钮")
        
    except Exception as e:
        print(f"❌ 导航失败: {e}")


def example_7_javascript_extraction():
    """示例 7: 使用 JavaScript 提取复杂数据"""
    logger.info("\n" + "=" * 60)
    logger.info("示例 7: JavaScript 提取")
    logger.info("=" * 60)
    
    agent = LocalAgentService(debug=True)
    
    try:
        # 打开页面
        agent.opencli.open("https://example.com/products")
        agent.opencli.wait("time", "1")
        
        # 使用 JavaScript 提取所有产品名称
        print("提取所有产品名称...")
        js_code = """
        (function() {
            const products = document.querySelectorAll('.product-name');
            const names = [];
            products.forEach(p => names.push(p.textContent.trim()));
            return JSON.stringify(names);
        })()
        """
        
        result = agent.opencli.eval_js(js_code)
        print(f"✅ 提取结果: {result}")
        
    except Exception as e:
        print(f"❌ 提取失败: {e}")


def example_8_session_caching():
    """示例 8: 会话和缓存管理"""
    logger.info("\n" + "=" * 60)
    logger.info("示例 8: 会话和缓存")
    logger.info("=" * 60)
    
    agent = LocalAgentService(debug=True)
    
    # 获取会话信息
    session = agent.get_session_info()
    print(f"会话 ID: {session['session_id']}")
    print(f"状态: {session['status']}")
    
    # 缓存数据
    print("\n缓存数据...")
    agent.cache["user_data"] = {
        "name": "John Doe",
        "email": "john@example.com"
    }
    
    # 检索缓存
    print("检索缓存...")
    if "user_data" in agent.cache:
        print(f"✅ 缓存数据: {agent.cache['user_data']}")


# 主入口
if __name__ == "__main__":
    print("""
╔════════════════════════════════════════════════════════════════════════════╗
║           澎湖數位老船長 - 本地 Agent 使用示例                             ║
║         Penghu Digital Captain - Local Agent Examples                      ║
╚════════════════════════════════════════════════════════════════════════════╝
    """)
    
    print("选择要运行的示例:")
    print("  1. 简单导航和提取")
    print("  2. 表单填充和提交")
    print("  3. 表格数据提取")
    print("  4. 多步工作流")
    print("  5. 动态滚动")
    print("  6. 条件导航")
    print("  7. JavaScript 提取")
    print("  8. 会话和缓存")
    print("  0. 运行所有示例")
    
    choice = input("\n请输入选择 (0-8): ").strip()
    
    examples = {
        "1": example_1_simple_navigation,
        "2": example_2_form_filling,
        "3": example_3_table_extraction,
        "4": example_4_multi_step_workflow,
        "5": example_5_dynamic_scrolling,
        "6": example_6_conditional_navigation,
        "7": example_7_javascript_extraction,
        "8": example_8_session_caching,
    }
    
    if choice == "0":
        # 运行所有示例
        for example_func in examples.values():
            try:
                # example_func()
                print(f"\n⏭️  跳过 {example_func.__name__} (需要真实网站)\n")
            except Exception as e:
                logger.error(f"示例失败: {e}")
    elif choice in examples:
        try:
            # examples[choice]()
            print(f"\n⏭️  跳过示例 (需要真实网站)")
        except Exception as e:
            logger.error(f"示例失败: {e}")
    else:
        print("❌ 无效的选择")
