package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"strconv"

	"github.com/astaxie/beego/orm"
)

type Sustituto struct {
	Id                            int               `orm:"column(id);pk"`
	Beneficiario           				int 					    `orm:"column(beneficiario)"`
	Porcentaje										int					      `orm:"column(porcentaje)"`
	Estado                        string            `orm:"column(estado);null"`
	Tutor	 												int								`orm:"column(tutor)"`
}

type SustitutoAdicionales struct {	
	Proveedor											int								`orm:"column(informacion_proveedor)"`
	Beneficiario           				int 					    `orm:"column(beneficiario)"`
	Porcentaje										int					      `orm:"column(porcentaje)"`
	Estado                        string            `orm:"column(estado);null"`
	NumeroContrato								string						`orm:"column(numero_contrato);null"`
	Tutor	 												int								`orm:"column(tutor)"`
}




func (t *Sustituto) TableName() string {
	return "sustituto"
}

func init() {
	orm.RegisterModel(new(Sustituto))
}

// AddSustituto insert a new Sustituto into database and returns
// last inserted Id on success.
func AddSustituto(m *Sustituto) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func GetSustituto( idProveedorString int) (v []SustitutoAdicionales) {
	o := orm.NewOrm()
	var temp [] SustitutoAdicionales
	id_proveedor := strconv.Itoa(idProveedorString)
	_, err := o.Raw("SELECT beneficiario.informacion_proveedor,sustituto.beneficiario,sustituto.porcentaje,sustituto.estado,contrato.numero_contrato,sustituto.tutor  FROM personal.sustituto AS sustituto, personal.beneficiarios AS beneficiario, personal.informacion_pensionado AS pensionado, agora.informacion_proveedor AS informacionProveedor,argo.contrato_general as contrato WHERE sustituto.beneficiario = beneficiario.id  AND beneficiario.informacion_proveedor = informacionProveedor.id_proveedor AND beneficiario.informacion_pensionado = pensionado.informacion_proveedor AND informacionProveedor.num_documento = contrato.contratista AND beneficiario.informacion_pensionado =" +id_proveedor).QueryRows(&temp)
  //_, err := o.Raw("SELECT beneficiario.informacion_proveedor, informacionProveedor.num_documento, informacionProveedor.nom_proveedor, beneficiario.categoria_beneficiario FROM personal.beneficiario AS beneficiario, agora.informacion_proveedor AS informacionProveedor, personal.sustituto AS sustituto WHERE sustituto.beneficiario = beneficiario.id AND beneficiario.id = " + "346" +"AND beneficiario.informacion_proveedor = informacionproveedor.id_proveedor").QueryRows(&temp)
	if err == nil {
		fmt.Println("Consulta exitosa")
	}
	fmt.Println("?=)=?=)(/&)")
	fmt.Println(temp[0].Beneficiario)
	return temp
}

func GetTutorSustituto( idProveedorString int) (v []Sustituto) {
	o := orm.NewOrm()
	var temp [] Sustituto
	id_proveedor := strconv.Itoa(idProveedorString)
	_, err := o.Raw("SELECT sustituto.beneficiario,sustituto.porcentaje,sustituto.estado,sustituto.tutor,contrato.numero_contrato FROM personal.sustituto AS sustituto, personal.beneficiarios AS beneficiario, personal.informacion_persona_pensionado AS pensionado, agora.informacion_proveedor AS informacionProveedor,argo.contrato_general as contrato WHERE sustituto.beneficiario = beneficiario.id  AND sustituto.tutor = informacionProveedor.id_proveedor AND beneficiario.informacion_pensionado = pensionado.informacion_proveedor AND informacionProveedor.num_documento = contrato.contratista AND beneficiario.informacion_pensionado = " +id_proveedor).QueryRows(&temp)
  //_, err := o.Raw("SELECT beneficiario.informacion_proveedor, informacionProveedor.num_documento, informacionProveedor.nom_proveedor, beneficiario.categoria_beneficiario FROM personal.beneficiario AS beneficiario, agora.informacion_proveedor AS informacionProveedor, personal.sustituto AS sustituto WHERE sustituto.beneficiario = beneficiario.id AND beneficiario.id = " + "346" +"AND beneficiario.informacion_proveedor = informacionproveedor.id_proveedor").QueryRows(&temp)
	if err == nil {
		fmt.Println("Consulta exitosa")
	}
	fmt.Println("?=)=?=)")
	//fmt.Println(temp[0].Beneficiario)
	return temp
}
// GetSustitutoById retrieves Sustituto by Id. Returns error if
// Id doesn't exist
func GetSustitutoById(id int) (v *Sustituto, err error) {
	o := orm.NewOrm()
	v = &Sustituto{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSustituto retrieves all Sustituto matches certain condition. Returns empty list if
// no records exist
func GetAllSustituto(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Sustituto))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Sustituto
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateSustituto updates Sustituto by Id and returns error if
// the record to be updated doesn't exist
func UpdateSustitutoById(m *Sustituto) (err error) {
	o := orm.NewOrm()
	v := Sustituto{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSustituto deletes Sustituto by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSustituto(id int) (err error) {
	o := orm.NewOrm()
	v := Sustituto{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Sustituto{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
