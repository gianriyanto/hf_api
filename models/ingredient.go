package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Ingredient struct {
	Id   int64  `orm:"pk;auto;column(id)" json:"id"`
	Name string `orm:"column(name)" json:"name"`
}

func AddIngredient(i Ingredient) int64 {
	o := orm.NewOrm()

	id, err := o.Insert(&i)
	if err == nil {
		return id
	}

	return i.Id
}

func GetIngredient(iid int64) (i *Ingredient, err error) {
	o := orm.NewOrm()
	var ingredient Ingredient
	err = o.QueryTable("ingredient").Filter("id", iid).RelatedSel().One(&ingredient)
	if err == nil {
		i = &ingredient
		return
	}

	return
}

func GetAllIngredients() (ingredient []Ingredient) {
	m := orm.NewOrm()
	_, err := m.QueryTable("ingredient").RelatedSel().All(&ingredient)
	if err == nil {
		return
	}

	ingredient = []Ingredient{}

	return
}

func UpdateIngredient(iid int64, ii *Ingredient) (i *Ingredient, err error) {
	o := orm.NewOrm()

	i, err = GetIngredient(iid)
	if err != nil {
		err = errors.New("Ingredient Does Not Exist")
		return
	}

	if ii.Name != "" {
		i.Name = ii.Name
	}

	if _, err = o.Update(i); err == nil {
		return
	}

	return nil, errors.New("Ingredient Does Not Exist")
}

func DeleteIngredient(iid int64) {
	o := orm.NewOrm()
	if num, err := o.Delete(&Ingredient{Id: iid}); err == nil {
		fmt.Println(num)
	}
}
