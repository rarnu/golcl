package main

import (
	"github.com/rarnu/golcl/lcl"
	"github.com/rarnu/golcl/lcl/types"
)

func main() {
	lcl.Application.Initialize()
	lcl.Application.SetMainFormOnTaskBar(true)
	lcl.Application.CreateForm(&MainForm)
	lcl.Application.CreateForm(&AboutForm)
	lcl.Application.Run()
}

type TMainForm struct {
	*lcl.TForm
	Btn1 *lcl.TButton
}

type TAboutForm struct {
	*lcl.TForm
	Btn1 *lcl.TButton
}

var MainForm *TMainForm
var AboutForm *TAboutForm

func (f *TMainForm) OnFormCreate(sender lcl.IObject) {

	f.SetCaption("Main")
	f.SetPosition(types.PoScreenCenter)
	f.SetWidth(600)
	f.SetHeight(400)

	f.Btn1 = lcl.NewButton(f)
	f.Btn1.SetParent(f)
	f.Btn1.SetCaption("About")
	f.Btn1.SetLeft(10)
	f.Btn1.SetTop(10)
	f.Btn1.SetOnClick(f.OnBtn1Click)
}

func (f *TMainForm) OnBtn1Click(sender lcl.IObject) {
	AboutForm.ShowModal()
}

func (f *TAboutForm) OnFormCreate(sender lcl.IObject) {
	f.SetCaption("About")
	f.SetPosition(types.PoMainFormCenter)
	f.SetWidth(300)
	f.SetHeight(200)

	f.Btn1 = lcl.NewButton(f)
	f.Btn1.SetParent(f)
	f.Btn1.SetCaption("OK")
	f.Btn1.SetLeft(10)
	f.Btn1.SetTop(10)
	f.Btn1.SetOnClick(f.OnBtn1Click)
}

func (f *TAboutForm) OnBtn1Click(sender lcl.IObject) {
	f.Close()
}
