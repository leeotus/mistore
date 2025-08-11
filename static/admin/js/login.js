// $(function(){
//     loginApp.init();
// })
// var loginApp={
//     init:function(){
//         this.getCaptcha()
//         this.captchaImgChage()
//     },
//     getCaptcha:function(){
//         $.get("/admin/code?t="+Math.random(),function(response){
//             console.log(response)
//             $("#captchaId").val(response.captchaId)
//             $("#captchaImg").attr("src",response.captchaImage)
//         })
//     },
//     captchaImgChage:function(){
//         var that=this;
//         $("#captchaImg").click(function(){
//             that.getCaptcha()
//         })
//     }
// }

$(function(){
    loginApp.init();
})
var loginApp={
    init:function(){
        this.getCaptcha()
        this.captchaImgChage()
        this.setupLoginForm()
    },
    getCaptcha:function(){
        $.get("/admin/code?t="+Math.random(),function(response){
            $("#captchaId").val(response.captchaId)
            $("#captchaImg").attr("src",response.captchaImage)
        })
    },
    captchaImgChage:function(){
        var that=this;
        $("#captchaImg").click(function(){
            that.getCaptcha()
        })
    },
    setupLoginForm:function(){
        $("#loginForm").on("submit", function(e){
            e.preventDefault();
            $.ajax({
                url: "/admin/doLogin",
                type: "POST",
                data: $(this).serialize(),
                success: function(response){
                    // 检查响应头中的认证信息
                    const authHeader = response.getResponseHeader("Authorization");
                    const refreshToken = response.getResponseHeader("X-Refresh-Token");

                    if(authHeader) {
                        // 将token存储在localStorage中
                        localStorage.setItem("authToken", authHeader);
                    }
                    if(refreshToken) {
                        localStorage.setItem("refreshToken", refreshToken);
                    }

                    // 执行重定向
                    window.location.href = "/admin";
                },
                error: function(xhr){
                    if(xhr.status === 401){
                        window.location.href = "/admin/login";
                    }
                }
            });
        });
    }
}
