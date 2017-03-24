// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"titan_api_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/funcionario_cargo",
				  beego.NSInclude(
				    &controllers.FuncionarioCargoController{},
				  ),
				),
		beego.NSNamespace("/tipo_vinculacion",
			beego.NSInclude(
				&controllers.TipoVinculacionController{},
			),
		),
		beego.NSNamespace("/preliquidacion",
		  beego.NSInclude(
		    &controllers.PreliquidacionController{},
		  ),
		),
		beego.NSNamespace("/concepto",
		  beego.NSInclude(
		    &controllers.ConceptoController{},
		  ),
		),

		beego.NSNamespace("/nomina",
		  beego.NSInclude(
		    &controllers.NominaController{},
		  ),
		),

		beego.NSNamespace("/concepto_por_persona",
		  beego.NSInclude(
		    &controllers.ConceptoPorPersonaController{},
		  ),
		),

		beego.NSNamespace("/tipo_nomina",
		  beego.NSInclude(
		    &controllers.TipoNominaController{},
		  ),
		),
		beego.NSNamespace("/acta_inicio",
		  beego.NSInclude(
		    &controllers.ActaInicioController{},
		  ),
		),

		beego.NSNamespace("/parametro_estandar",
		  beego.NSInclude(
		    &controllers.ParametroEstandarController{},
		  ),
		),

		beego.NSNamespace("/contrato_general",
		  beego.NSInclude(
		    &controllers.ContratoGeneralController{},
		  ),
		),

		beego.NSNamespace("/argo_ordenadores",
		  beego.NSInclude(
		    &controllers.ArgoOrdenadoresController{},
		  ),
		),

		beego.NSNamespace("/relacion_parametro",
		  beego.NSInclude(
		    &controllers.RelacionParametroController{},
		  ),
		),

		beego.NSNamespace("/unidad_ejecutora",
		  beego.NSInclude(
		    &controllers.UnidadEjecutoraController{},
		  ),
		),

		beego.NSNamespace("/tipo_contrato",
		  beego.NSInclude(
		    &controllers.TipoContratoController{},
		  ),
		),

		beego.NSNamespace("/departamento",
		  beego.NSInclude(
		    &controllers.DepartamentoController{},
		  ),
		),

		beego.NSNamespace("/unidad_ejecutora",
		  beego.NSInclude(
		    &controllers.UnidadEjecutoraController{},
		  ),
		),

		beego.NSNamespace("/pais",
		  beego.NSInclude(
		    &controllers.PaisController{},
		  ),
		),

		beego.NSNamespace("/informacion_proveedor",
		  beego.NSInclude(
		    &controllers.InformacionProveedorController{},
		  ),
		),

		beego.NSNamespace("/lugar_ejecucion",
		  beego.NSInclude(
		    &controllers.LugarEjecucionController{},
		  ),
		),

		beego.NSNamespace("/ciudad",
		  beego.NSInclude(
		    &controllers.CiudadController{},
		  ),
		),

		beego.NSNamespace("/parametros",
		  beego.NSInclude(
		    &controllers.ParametrosController{},
		  ),
		),

		beego.NSNamespace("/supervisor_contrato",
		  beego.NSInclude(
		    &controllers.SupervisorContratoController{},
		  ),
		),

		beego.NSNamespace("/concepto",
		  beego.NSInclude(
		    &controllers.ConceptoController{},
		  ),
		),

		beego.NSNamespace("/nomina",
		  beego.NSInclude(
		    &controllers.NominaController{},
		  ),
		),

		beego.NSNamespace("/concepto_por_persona",
		  beego.NSInclude(
		    &controllers.ConceptoPorPersonaController{},
		  ),
		),

		beego.NSNamespace("/detalle_preliquidacion",
		  beego.NSInclude(
		    &controllers.DetallePreliquidacionController{},
		  ),
		),

		beego.NSNamespace("/liquidacion",
		  beego.NSInclude(
		    &controllers.LiquidacionController{},
		  ),
		),

		beego.NSNamespace("/detalle_liquidacion",
		  beego.NSInclude(
		    &controllers.DetalleLiquidacionController{},
		  ),
		),
		beego.NSNamespace("/concepto",
		  beego.NSInclude(
		    &controllers.ConceptoController{},
		  ),
		),
		beego.NSNamespace("/docente_cargo",
	beego.NSInclude(
		&controllers.DocenteCargoController{},
	),
),

		beego.NSNamespace("/funcionario_proveedor",
		  beego.NSInclude(
		    &controllers.FuncionarioProveedorController{},
		  ),
		),

		beego.NSNamespace("/estado_civil",
			beego.NSInclude(
				&controllers.EstadoCivilController{},
			),
		),

		beego.NSNamespace("/categoria_beneficiario",
			beego.NSInclude(
				&controllers.CategoriaBeneficiarioController{},
			),
		),

		beego.NSNamespace("/beneficiarios",
			beego.NSInclude(
				&controllers.BeneficiariosController{},
			),
		),

		beego.NSNamespace("/tipo_pension",
			beego.NSInclude(
				&controllers.TipoPensionController{},
			),
		),

		beego.NSNamespace("/tipo_pensionado",
			beego.NSInclude(
				&controllers.TipoPensionadoController{},
			),
		),

		beego.NSNamespace("/resolucion",
			beego.NSInclude(
				&controllers.ResolucionController{},
			),
		),

		beego.NSNamespace("/informacion_pensionado",
			beego.NSInclude(
				&controllers.InformacionPensionadoController{},
			),
		),

		beego.NSNamespace("/sustituto",
			beego.NSInclude(
				&controllers.SustitutoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
