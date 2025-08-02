// 监听 iframe 加载事件，检查登录状态
function monitorIframeSession() {
    try {
        // 获取 iframe 元素
        const iframe = window.frames["rightMain"];
        if (iframe) {
            iframe.onload = function() {
                // 检查 iframe 中是否有登录提示标记
                if (iframe.document.getElementById("need-login")) {
                    window.location.href = "/admin/login";
                }
            };
        }
    } catch (e) {
        // 跨域或其他错误时，直接强制跳转登录页
        window.location.href = "/admin/login";
    }
}

// 监听所有 AJAX 请求的完成事件
function monitorAjaxRequests() {
    $(document).ajaxComplete(function(event, xhr, settings) {
        // 401 状态码表示未授权（Session 失效）
        if (xhr.status === 401) {
            window.location.href = "/admin/login";
        }
        // 也可以处理 302 重定向的情况（如果中间件返回重定向）
        else if (xhr.status === 302 && xhr.getResponseHeader("Location") === "/admin/login") {
            window.location.href = "/admin/login";
        }
    });
}

// 页面加载完成后初始化监听
$(function() {
    monitorIframeSession();
    monitorAjaxRequests();
});
