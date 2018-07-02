package controller

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"LearGoProject/Seckill/SecProxy/service"
)

type SkillController struct {
	beego.Controller
}

func (p *SkillController) SecKill()  {
	p.Data["json"] = "sec Kill"
	p.ServeJSON()
}

func (p *SkillController) SecInfo()  {
	//p.Data["json"] = "sec Info"
	productId,err := p.GetInt("product_id")
	result := make(map[string]interface{})
	result["code"] = 0
	result["message"] = "success"
	defer func() {
		p.Data["json"] = result
		p.ServeJSON()
	}()

	if err != nil {
		data,code,err := service.SecInfoList(productId)
		if err != nil{
			result["code"] = code
			result["message"] = err.Error()
			logs.Error("invaild request.get product_id failed,err:%v",err)
			return
		}
		result["code"] = code
		result["data"] = data
		return
	}else {
		data,code,err := service.SecInfo(productId)
		if err != nil{
			result["code"] = code
			result["message"] = err.Error()
			logs.Error("invaild request.get product_id failed,err:%v",err)
			return
		}
		result["data"] = data
	}


}

