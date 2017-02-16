package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["titan_api_crud/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ActaInicioController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ActaInicioController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ActaInicioController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ActaInicioController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ActaInicioController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ArgoOrdenadoresController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ArgoOrdenadoresController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ArgoOrdenadoresController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ArgoOrdenadoresController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ArgoOrdenadoresController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ArgoOrdenadoresController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ArgoOrdenadoresController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ArgoOrdenadoresController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ArgoOrdenadoresController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ArgoOrdenadoresController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:CategoriaBeneficiarioController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:CategoriaBeneficiarioController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:CategoriaBeneficiarioController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:CategoriaBeneficiarioController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:CategoriaBeneficiarioController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:CategoriaBeneficiarioController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:CategoriaBeneficiarioController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:CategoriaBeneficiarioController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:CategoriaBeneficiarioController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:CategoriaBeneficiarioController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:CiudadController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:CiudadController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:CiudadController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:CiudadController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:CiudadController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:CiudadController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:CiudadController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:CiudadController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:CiudadController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:CiudadController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ConceptoController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ConceptoController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ConceptoController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ConceptoController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ConceptoController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ConceptoPorPersonaController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ConceptoPorPersonaController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ConceptoPorPersonaController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ConceptoPorPersonaController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ConceptoPorPersonaController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ConceptoPorPersonaController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ConceptoPorPersonaController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ConceptoPorPersonaController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ConceptoPorPersonaController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ConceptoPorPersonaController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:DepartamentoController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:DepartamentoController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:DepartamentoController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:DepartamentoController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:DepartamentoController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:DetalleLiquidacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:DetalleLiquidacionController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:DetalleLiquidacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:DetalleLiquidacionController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:DetalleLiquidacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:DetalleLiquidacionController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:DetalleLiquidacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:DetalleLiquidacionController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:DetalleLiquidacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:DetalleLiquidacionController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:DetallePreliquidacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:DetallePreliquidacionController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:DetallePreliquidacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:DetallePreliquidacionController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:DetallePreliquidacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:DetallePreliquidacionController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:DetallePreliquidacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:DetallePreliquidacionController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:DetallePreliquidacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:DetallePreliquidacionController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:EstadoCivilController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:EstadoCivilController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:EstadoCivilController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:EstadoCivilController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:EstadoCivilController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:EstadoCivilController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:EstadoCivilController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:EstadoCivilController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:EstadoCivilController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:EstadoCivilController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:LiquidacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:LiquidacionController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:LiquidacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:LiquidacionController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:LiquidacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:LiquidacionController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:LiquidacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:LiquidacionController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:LiquidacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:LiquidacionController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:LugarEjecucionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:LugarEjecucionController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:LugarEjecucionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:LugarEjecucionController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:LugarEjecucionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:LugarEjecucionController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:LugarEjecucionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:LugarEjecucionController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:LugarEjecucionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:LugarEjecucionController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:NominaController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:NominaController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:NominaController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:NominaController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:NominaController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:NominaController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:NominaController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:NominaController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:NominaController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:NominaController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ObjectController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ObjectController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ObjectController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ObjectController"],
		beego.ControllerComments{
			"Get",
			`/:objectId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ObjectController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ObjectController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ObjectController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ObjectController"],
		beego.ControllerComments{
			"Put",
			`/:objectId`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ObjectController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ObjectController"],
		beego.ControllerComments{
			"Delete",
			`/:objectId`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:PaisController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:PaisController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:PaisController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:PaisController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:PaisController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:PaisController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:PaisController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:PaisController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:PaisController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:PaisController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ParametrosController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ParametrosController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ParametrosController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ParametrosController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ParametrosController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:PreliquidacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:PreliquidacionController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:PreliquidacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:PreliquidacionController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:PreliquidacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:PreliquidacionController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:PreliquidacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:PreliquidacionController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:PreliquidacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:PreliquidacionController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:RelacionParametroController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:RelacionParametroController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:RelacionParametroController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:RelacionParametroController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:RelacionParametroController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ResolucionController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ResolucionController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ResolucionController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ResolucionController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ResolucionController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:SustitutoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:SustitutoController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:SustitutoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:SustitutoController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:SustitutoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:SustitutoController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:SustitutoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:SustitutoController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:SustitutoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:SustitutoController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoContratoController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoContratoController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoContratoController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoContratoController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoContratoController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoNominaController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoNominaController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoNominaController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoNominaController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoNominaController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoNominaController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoNominaController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoNominaController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoNominaController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoNominaController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoPensionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoPensionController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoPensionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoPensionController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoPensionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoPensionController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoPensionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoPensionController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoPensionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoPensionController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoPensionadoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoPensionadoController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoPensionadoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoPensionadoController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoPensionadoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoPensionadoController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoPensionadoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoPensionadoController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoPensionadoController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoPensionadoController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoVinculacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoVinculacionController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoVinculacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoVinculacionController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoVinculacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoVinculacionController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoVinculacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoVinculacionController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:TipoVinculacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:TipoVinculacionController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:UserController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:UserController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:UserController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:UserController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:UserController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:UserController"],
		beego.ControllerComments{
			"Get",
			`/:uid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:UserController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:UserController"],
		beego.ControllerComments{
			"Put",
			`/:uid`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:UserController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:UserController"],
		beego.ControllerComments{
			"Delete",
			`/:uid`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:UserController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:UserController"],
		beego.ControllerComments{
			"Login",
			`/login`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["titan_api_crud/controllers:UserController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:UserController"],
		beego.ControllerComments{
			"Logout",
			`/logout`,
			[]string{"get"},
			nil})

}
