package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Menu struct {
	Id      int64  `orm:"pk;auto;column(id)" json:"id"`
	Name    string `orm:"column(name)" json:"name"`
	WeekNum int    `orm:"column(weeknum)" json:"weeknum"`
}

func AddMenu(m Menu) int64 {
	o := orm.NewOrm()

	id, err := o.Insert(&m)
	if err == nil {
		return id
	}

	return m.Id
}

func GetMenu(mid int64) (m *Menu, err error) {
	o := orm.NewOrm()
	var menu Menu
	err = o.QueryTable("menu").Filter("id", mid).RelatedSel().One(&menu)
	if err == nil {
		m = &menu
		return
	}

	return
}

func GetAllMenus() (menu []Menu) {
	m := orm.NewOrm()
	_, err := m.QueryTable("menu").RelatedSel().All(&menu)
	if err == nil {
		return
	}

	menu = []Menu{}

	return
}

func UpdateMenu(mid int64, mm *Menu) (m *Menu, err error) {
	o := orm.NewOrm()

	m, err = GetMenu(mid)
	if err != nil {
		err = errors.New("Menu Does Not Exist")
		return
	}

	if mm.Name != "" {
		m.Name = mm.Name
	}
	if mm.WeekNum != 0 {
		m.WeekNum = mm.WeekNum
	}

	if _, err = o.Update(m); err == nil {
		return
	}

	return nil, errors.New("Menu Does Not Exist")
}

func DeleteMenu(mid int64) {
	o := orm.NewOrm()
	if num, err := o.Delete(&Menu{Id: mid}); err == nil {
		fmt.Println(num)
	}
}
