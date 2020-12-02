package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Recipe struct {
	Id              int64            `orm:"pk;auto;column(id)" json:"id"`
	RecipeName      string           `orm:"column(recipename)" json:"recipename"`
	Instruction     string           `orm:"column(instruction)" json:"instruction"`
	Ratings         string           `orm:"column(ratings)" json:"ratings"`
	Nutritionalinfo *Nutritionalinfo `orm:"rel(one)" json:"nutritionalinfo"` // OneToOne relation
}

type Nutritionalinfo struct {
	Id      int64   `orm:"pk;auto;column(id)" json:"id"`
	Carbs   string  `orm:"column(carbs)" json:"carbs"`
	Protein string  `orm:"column(protein)" json:"protein"`
	Fat     string  `orm:"column(fat)" json:"fat"`
	Recipe  *Recipe `orm:"reverse(one)" json:"-"` // Reverse relationship (optional)
}

func AddRecipe(r Recipe) int64 {
	o := orm.NewOrm()

	_, err := o.Insert(r.Nutritionalinfo)
	if err != nil {
		return 0
	}

	id, err := o.Insert(&r)
	if err == nil {
		return id
	}

	return r.Id
}

func GetRecipe(rid int64) (r *Recipe, err error) {
	o := orm.NewOrm()
	var recipe Recipe
	err = o.QueryTable("recipe").Filter("id", rid).RelatedSel().One(&recipe)
	if err == nil {
		r = &recipe
		return
	}

	return
}

func GetAllRecipes() (recipe []Recipe) {
	o := orm.NewOrm()
	_, err := o.QueryTable("recipe").RelatedSel().All(&recipe)
	if err == nil {
		return
	}

	recipe = []Recipe{}

	return
}

func UpdateRecipe(rid int64, rr *Recipe) (r *Recipe, err error) {
	o := orm.NewOrm()

	r, err = GetRecipe(rid)
	if err != nil {
		err = errors.New("Recipe Not Exist")
		return
	}

	if rr.RecipeName != "" {
		r.RecipeName = rr.RecipeName
	}
	if rr.Instruction != "" {
		r.Instruction = rr.Instruction
	}
	if rr.Nutritionalinfo.Carbs != "" {
		r.Nutritionalinfo.Carbs = rr.Nutritionalinfo.Carbs
	}
	if rr.Nutritionalinfo.Protein != "" {
		r.Nutritionalinfo.Protein = rr.Nutritionalinfo.Protein
	}
	if rr.Nutritionalinfo.Fat != "" {
		r.Nutritionalinfo.Fat = rr.Nutritionalinfo.Fat
	}

	if _, err = o.Update(r.Nutritionalinfo); err == nil {
		return
	}
	if _, err = o.Update(r); err == nil {
		return
	}

	return nil, errors.New("Recipe Does Not Exist")
}

func DeleteRecipe(rid int64) {
	o := orm.NewOrm()
	if num, err := o.Delete(&Recipe{Id: rid}); err == nil {
		fmt.Println(num)
	}
}
