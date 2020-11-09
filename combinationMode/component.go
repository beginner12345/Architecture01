package main

import "fmt"

// Component
type Component interface {
	PrintComponent()
}

// WinForm
type WinForm struct {
	v []Component
	tag string
}

func (wf WinForm) PrintComponent() {
	fmt.Printf("print WinForm (%s)\r\n", wf.tag)
	for i:= range wf.v {
		wf.v[i].PrintComponent()
	}
}

func (wf *WinForm) AddSubComponent(c Component) {
	wf.v = append(wf.v, c)
}

// Picture
type Picture struct {
	tag string
}

func (p Picture) PrintComponent() {
	fmt.Printf("print Picture (%s)\r\n", p.tag)
}

// Button
type Button struct {
	tag string
}

func (b Button) PrintComponent() {
	fmt.Printf("print Button (%s)\r\n", b.tag)
}

// Frame
type Frame struct {
	v []Component
	tag string
}

func (f Frame) PrintComponent() {
	fmt.Printf("print Frame (%s)\r\n", f.tag)
	for i:= range(f.v) {
		f.v[i].PrintComponent()
	}
}

func (f *Frame) AddSubComponent(c Component) {
	f.v = append(f.v, c)
}

// Lable
type Lable struct {
	tag string
}

func (l Lable) PrintComponent() {
	fmt.Printf("print Lable (%s)\r\n", l.tag)
}

// TextBox
type TextBox struct {
	tag string
}

func (t TextBox) PrintComponent() {
	fmt.Printf("print TextBox (%s)\r\n", t.tag)
}

// PasswordBox
type PasswordBox struct {
	tag string
}

func (p PasswordBox) PrintComponent() {
	fmt.Printf("print PasswordBox (%s)\r\n", p.tag)
}

// CheckBox
type CheckBox struct {
	tag string
}

func (cb CheckBox) PrintComponent() {
	fmt.Printf("print CheckBox (%s)\r\n", cb.tag)
}

// LinkLable
type LinkLable struct {
	tag string
}

func (ll LinkLable) PrintComponent() {
	fmt.Printf("print LinkLable (%s)\r\n", ll.tag)
}
