package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"strconv"
	"time"
	"github.com/astaxie/beego/orm"
)

type Beneficiarios struct {
	Id                    int                    `orm:"column(id);pk"`
	InformacionPensionado int                    `orm:"column(informacion_pensionado);null"`
	InformacionProveedor  int  									`orm:"column(informacion_proveedor);"`
	FechaNacBeneficiario  time.Time              `orm:"column(fecha_nac_beneficiario);type(date);null"`
	Tutor                 int                    `orm:"column(tutor);null"`
	SubFamiliar           string                 `orm:"column(sub_familiar);null"`
	CategoriaBeneficiario int 										`orm:"column(categoria_beneficiario);"`
	SubEstudios           string                 `orm:"column(aux_estudio);null"`
	Estado								string									`orm:"column(estado);null"`
}

/*type numero_beneficiarios struct {
	numero  int `orm:"column(informacion_pensionado)"`
}*/

func (t *Beneficiarios) TableName() string {
	return "beneficiarios"
}

func init() {
	orm.RegisterModel(new(Beneficiarios))
}

// AddBeneficiario insert a new Beneficiario into database and returns
// last inserted Id on success.
func AddBeneficiario(m *Beneficiarios) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetBeneficiarioById retrieves Beneficiario by Id. Returns error if
// Id doesn't exist
func GetnumBeneficiario_x_pensionado(idProveedorString int) (personas []Beneficiarios) {
	o := orm.NewOrm()
	var temp [] Beneficiarios
	fmt.Println(idProveedorString)
	id_proveedor := strconv.Itoa(idProveedorString)
	//_, err := o.Raw("SELECT count(beneficiario.informacion_pensionado) FROM personal.beneficiario AS beneficiario").QueryRows(&temp)
	_, err := o.Raw("SELECT ben.informacion_proveedor,ben.categoria_beneficiario,ben.sub_familiar,ben.aux_estudio, ben.estado FROM personal.beneficiarios AS ben, personal.informacion_persona_pensionado AS pen WHERE pen.informacion_proveedor = ben.informacion_pensionado AND ben.informacion_pensionado = " +id_proveedor).QueryRows(&temp)
	//if len(temp) == 0{
		//temp = append(temp,"0")
	//}
	if err == nil {
		fmt.Println("Consulta exitosa")
	}
	fmt.Println("Teeeeeeeeeeeemp")
return temp
}

func GetBeneficiarioById(id int) (v *Beneficiarios, err error) {
	o := orm.NewOrm()
	v = &Beneficiarios{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllBeneficiario retrieves all Beneficiario matches certain condition. Returns empty list if
// no records exist
func GetAllBeneficiario(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Beneficiarios))
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

	var l []Beneficiarios
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

// UpdateBeneficiario updates Beneficiario by Id and returns error if
// the record to be updated doesn't exist
func UpdateBeneficiarioById(m *Beneficiarios) (err error) {
	o := orm.NewOrm()
	v := Beneficiarios{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteBeneficiario deletes Beneficiario by Id and returns error if
// the record to be deleted doesn't exist
func DeleteBeneficiario(id int) (err error) {
	o := orm.NewOrm()
	v := Beneficiarios{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Beneficiarios{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
