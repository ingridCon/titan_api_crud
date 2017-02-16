package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["titan_api_crud/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ActaInicioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ActaInicioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ActaInicioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ActaInicioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ActaInicioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ArgoOrdenadoresController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ArgoOrdenadoresController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ArgoOrdenadoresController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ArgoOrdenadoresController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ArgoOrdenadoresController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ArgoOrdenadoresController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ArgoOrdenadoresController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ArgoOrdenadoresController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ArgoOrdenadoresController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ArgoOrdenadoresController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:CategoriaBeneficiarioController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:CategoriaBeneficiarioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:CategoriaBeneficiarioController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:CategoriaBeneficiarioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:CategoriaBeneficiarioController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:CategoriaBeneficiarioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:CategoriaBeneficiarioController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:CategoriaBeneficiarioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:CategoriaBeneficiarioController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:CategoriaBeneficiarioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:CiudadController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:CiudadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:CiudadController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:CiudadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:CiudadController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:CiudadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:CiudadController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:CiudadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:CiudadController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:CiudadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ConceptoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ConceptoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ConceptoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ConceptoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ConceptoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ConceptoPorPersonaController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ConceptoPorPersonaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ConceptoPorPersonaController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ConceptoPorPersonaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ConceptoPorPersonaController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ConceptoPorPersonaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ConceptoPorPersonaController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ConceptoPorPersonaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ConceptoPorPersonaController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ConceptoPorPersonaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:DetalleLiquidacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:DetalleLiquidacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:DetalleLiquidacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:DetalleLiquidacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:DetalleLiquidacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:DetalleLiquidacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:DetalleLiquidacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:DetalleLiquidacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:DetalleLiquidacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:DetalleLiquidacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:DetallePreliquidacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:DetallePreliquidacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:DetallePreliquidacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:DetallePreliquidacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:DetallePreliquidacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:DetallePreliquidacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:DetallePreliquidacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:DetallePreliquidacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:DetallePreliquidacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:DetallePreliquidacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:EstadoCivilController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:EstadoCivilController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:EstadoCivilController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:EstadoCivilController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:EstadoCivilController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:EstadoCivilController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:EstadoCivilController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:EstadoCivilController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:EstadoCivilController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:EstadoCivilController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:LiquidacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:LiquidacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:LiquidacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:LiquidacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:LiquidacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:LiquidacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:LiquidacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:LiquidacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:LiquidacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:LiquidacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:LugarEjecucionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:LugarEjecucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:LugarEjecucionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:LugarEjecucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:LugarEjecucionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:LugarEjecucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:LugarEjecucionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:LugarEjecucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:LugarEjecucionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:LugarEjecucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:NominaController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:NominaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:NominaController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:NominaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:NominaController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:NominaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:NominaController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:NominaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:NominaController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:NominaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ObjectController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ObjectController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ObjectController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ObjectController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ObjectController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:PaisController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:PaisController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:PaisController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:PaisController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:PaisController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:PaisController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:PaisController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:PaisController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:PaisController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:PaisController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:PreliquidacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:PreliquidacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:PreliquidacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:PreliquidacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:PreliquidacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:PreliquidacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:PreliquidacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:PreliquidacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:PreliquidacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:PreliquidacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:RelacionParametroController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:RelacionParametroController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:RelacionParametroController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:RelacionParametroController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:RelacionParametroController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:SustitutoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:SustitutoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:SustitutoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:SustitutoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:SustitutoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:SustitutoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:SustitutoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:SustitutoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:SustitutoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:SustitutoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoNominaController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoNominaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoNominaController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoNominaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoNominaController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoNominaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoNominaController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoNominaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoNominaController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoNominaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoPensionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoPensionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoPensionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoPensionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoPensionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoPensionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoPensionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoPensionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoPensionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoPensionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoPensionadoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoPensionadoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoPensionadoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoPensionadoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoPensionadoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoPensionadoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoPensionadoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoPensionadoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoPensionadoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoPensionadoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoVinculacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoVinculacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoVinculacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoVinculacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoVinculacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoVinculacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoVinculacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoVinculacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoVinculacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoVinculacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:UserController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:UserController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:UserController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:UserController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:UserController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:UserController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:UserController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
