var target_Module = "com.snapone.lumaview";

	if (ObjC.available){
	    console.log("===hooking!======")
	    var class_name = "SuperCamHD_ProAppDelegate";
	    var method_name = "- iphoneIsJailbroken";
	    var hook = eval('ObjC.classes.' + class_name + '["' + method_name + '"]');
	    Interceptor.attach(hook.implementation, {
	        onEnter:function(args){
	            console.log("==== 函数执行前执行 ====")
	        },
	        onLeave:function(retval){
	            console.log("==== 函数返回前执行 ====")
	            retval.replace(0x0);    //修改函数的返回值
	            console.log("Return: " + retval)
	        }
	    });
	}