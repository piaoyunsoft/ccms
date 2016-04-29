package v2

import (
	"api/models/v2"
	"github.com/astaxie/beego"
	"api/util"
)


type PersonController struct {
	BaseController
}

func (p *PersonController) List() {
	persons :=  v2.List()
	p.Data["json"] = persons
	p.ServeJSON()
}

func (p *PersonController) Get() {
	// userId := p.GetString("userId")
	// beego.Informational("userId",userId)
	// persons :=  v2.Get(userId)
	requestParams := p.params()
	person := v2.Get(requestParams["userId"])
	beego.Informational(p.params())

	p.RespResult(util.SUCCESS,util.SUCCESS_MSG,person)
}