package controllers

import (
	"encoding/json"
"titan_api_crud/models"
	"github.com/astaxie/beego"
	"fmt"
)

// InformacionPensionadoController operations for InformacionPensionado
type InformacionPensionadoController struct {
	beego.Controller
}

// URLMapping ...
func (c *InformacionPensionadoController) URLMapping() {
	c.Mapping("Pensionado_datos", c.Pensionado_datos)
}

func (c *InformacionPensionadoController) Pensionado_datos() {
	var v models.InformacionPensionado
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		respuesta := models.GetPensionado(v.Id)
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = respuesta
	} else {
		c.Data["json"] = err.Error()
		fmt.Println("error 2: ", err)
	}
	c.ServeJSON()

}
