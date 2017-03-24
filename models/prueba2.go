package models

import (
	//"errors"
	//"fmt"
	//"reflect"
	//"strings"
	//"strconv"

	"github.com/astaxie/beego/orm"
)

type Prueba2 struct {
	Id                    int                    `orm:"column(id);pk"`
	TipoDocumento int                    `orm:"column(tipo_documento);"`
}

func (t *Prueba2) TableName() string {
	return "prueba2"
}

func init() {
	orm.RegisterModel(new(Prueba2))
}
