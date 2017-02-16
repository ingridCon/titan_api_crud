package controllers

import (
	"encoding/json"
	"fmt"
	"titan_api_crud/models"

	"github.com/astaxie/beego"
)

// operations for DocenteCargo
type DocenteCargoController struct {
	beego.Controller
}

func (c *DocenteCargoController) URLMapping() {
	c.Mapping("DocenteCargo", c.ConsultarAsignacionBasica)
	c.Mapping("DocenteCedula", c.ConsultarCedulaDocente)
}

func (c *DocenteCargoController) ConsultarAsignacionBasica() {
	var v models.Docente_x_Cargo
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		respuesta := models.GetAsignacionBasicaDocente(v.Id)
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = respuesta
	} else {
		c.Data["json"] = err.Error()
		fmt.Println("error 2: ", err)
	}
	c.ServeJSON()
}

func (c *DocenteCargoController) ConsultarCedulaDocente() {
	var v string

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		respuesta := models.GetCedulaDocente(v)
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = respuesta
	} else {
		c.Data["json"] = err.Error()
		fmt.Println("error 2: ", err)
	}
	c.ServeJSON()

}
