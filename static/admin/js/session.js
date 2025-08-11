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
// static/admin/js/session.js
function monitorAjaxRequests() {
    $(document).ajaxComplete(function(event, xhr, settings) {
        // 401状态码表示未授权（JWT失效）
        if (xhr.status === 401) {
            window.location.href = "/admin/login";
        }
        // 检查响应头中的认证信息
        const authHeader = xhr.getResponseHeader("Authorization");
        const refreshToken = xhr.getResponseHeader("X-Refresh-Token");
        if (authHeader) {
            localStorage.setItem('authToken', authHeader);
        }
        if (refreshToken) {
            localStorage.setItem('refreshToken', refreshToken);
        }
    });
}

// 页面加载时检查认证状态
function checkAuthStatus() {
    const token = localStorage.getItem('authToken');
    if (!token) {
        window.location.href = "/admin/login";
    }
}

// 页面加载完成后初始化
$(function() {
    checkAuthStatus();
    monitorIframeSession();
    monitorAjaxRequests();
});


// 页面加载完成后初始化监听
$(function() {
    checkAuthStatus()
    monitorIframeSession();
    monitorAjaxRequests();
});
