(function($){
    $(function(){
        loginApp.init();
    })
    var loginApp={
        init:function(){
            this.getCaptcha()
            this.captchaImgChage()
        },
        getCaptcha:function(){
            $.get("/pass/captcha?t="+Math.random(),function(response){              
                $("#captchaId").val(response.captchaId)
                $("#captchaImg").attr("src",response.captchaImage)
            })
        },
        captchaImgChage:function(){
            var that=this;
            $("#captchaImg").click(function(){
                that.getCaptcha()
            })
        }
    }
})($)

