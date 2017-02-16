package controllers

import (
  	"encoding/json"
//	"errors"
	//"strconv"
	//"strings"
	"titan_api_crud/models"

	"github.com/astaxie/beego"
  "fmt"
)

// ConceptoController operations for Concepto
type FuncionarioCargoController struct {
	beego.Controller
}

// URLMapping ...
func (c *FuncionarioCargoController) URLMapping() {
	c.Mapping("FuncionarioCargo", c.ConsultarAsignacionBasica)

}


func (c *FuncionarioCargoController) ConsultarAsignacionBasica() {
  var v models.Funcionario_x_Cargo
  if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {

    respuesta := models.GetAsignacionBasica(v.Id)
    c.Ctx.Output.SetStatus(201)
  
    c.Data["json"] = respuesta


  } else {
    c.Data["json"] = err.Error()
		fmt.Println("error 2: ", err)
  }
  c.ServeJSON()
  }
