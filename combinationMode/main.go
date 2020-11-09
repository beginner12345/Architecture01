package main

func main() {
	// 创建winform，并添加子元素
	wf := WinForm {tag: "WINDOWS窗口", v: make([]Component, 0, 10)}
	pic := Picture{tag: "LOGO图片"}
	wf.AddSubComponent(pic)
	bt1 := Button{tag: "登录"}
	wf.AddSubComponent(bt1)
	bt2 := Button{tag: "注册"}
	wf.AddSubComponent(bt2)
	//创建frame，并添加子元素
	f := Frame{tag: "FRAME1", v: make([]Component, 0, 10)}
	lb1 := Lable{tag: "用户名"}
	f.AddSubComponent(lb1)
	tb := TextBox{tag: "文本框"}
	f.AddSubComponent(tb)
	lb2 := Lable{tag: "密码"}
	f.AddSubComponent(lb2)
	pb := PasswordBox{tag: "密码框"}
	f.AddSubComponent(pb)
	cb := CheckBox{tag: "复选框"}
	f.AddSubComponent(cb)
	lb3 := Lable{tag: "记住用户名"}
	f.AddSubComponent(lb3)
	ll := LinkLable{tag: "忘记密码"}
	f.AddSubComponent(ll)
	//将frame添加到winform
	wf.AddSubComponent(f)
	//打印winform
	wf.PrintComponent()
}
