package controllers

import (
	"github.com/astaxie/beego"
	"phoneLocationOpenApi/models"
)

type PlController struct {
	beego.Controller
}

type pl struct {
	Data    *models.PhoneRecord
	Code    int
	Message string
}

func (p *PlController) Get() {
	num := p.GetString("phoneNum", "18000000000")

	var loc = &pl{
		Data: &models.PhoneRecord{
			PhoneNum: num,
		},
		Message: "success",
	}

	var err error
	loc.Data, err = models.Find(num)
	if err != nil {
		p.Data["json"] = &pl{
			Data: &models.PhoneRecord{
				PhoneNum: num,
			},
			Code:    -2,
			Message: "Server is busy.",
		}

	} else {
		p.Data["json"] = &loc
	}

	p.ServeJSON()
	return

}
