package controllers

import (
	"encoding/json"
	"hf_api/models"

	"github.com/astaxie/beego"
)

// Operations about Ingredients
type IngredientController struct {
	beego.Controller
}

// @Title CreateIngredient
// @Description create ingredients
// @Param	body		body 	models.Ingredient	true		"body for ingredient content"
// @Success 200 {int} models.Ingredient.Id
// @Failure 403 body is empty
// @router /create [post]
func (i *IngredientController) Post() {
	var ingredient models.Ingredient
	json.Unmarshal(i.Ctx.Input.RequestBody, &ingredient)
	iid := models.AddIngredient(ingredient)
	i.Data["json"] = map[string]int64{"iid": iid}
	i.ServeJSON()
}

// @Title GetIngredient
// @Description get all Ingredients
// @Success 200 {object} models.Ingredient
// @router / [get]
func (i *IngredientController) GetAll() {
	ingredients := models.GetAllIngredients()
	i.Data["json"] = ingredients
	i.ServeJSON()
}

// @Title Get
// @Description get ingredient by iid
// @Param	iid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Ingredient
// @Failure 403 :iid is empty
// @router /:iid [get]
func (i *IngredientController) Get() {
	iid, _ := i.GetInt64(":iid")
	ingredient, err := models.GetIngredient(iid)
	if err != nil {
		i.Data["json"] = err.Error()
	} else {
		i.Data["json"] = ingredient
	}
	i.ServeJSON()
}

// @Title Update
// @Description update the ingredient
// @Param	iid		path 	string	true		"The mid you want to update"
// @Param	body		body 	models.Ingredient	true		"body for ingredient content"
// @Success 200 {object} models.Ingredient
// @Failure 403 :iid is not int
// @router /:iid [put]
func (i *IngredientController) Put() {
	iid, _ := i.GetInt64(":iid")
	var ingredient models.Ingredient
	json.Unmarshal(i.Ctx.Input.RequestBody, &ingredient)
	ii, err := models.UpdateIngredient(iid, &ingredient)
	if err != nil {
		i.Data["json"] = err.Error()
	} else {
		i.Data["json"] = ii
	}
	i.ServeJSON()
}

// @Title Delete
// @Description delete the ingredient
// @Param	iid		path 	string	true		"The iid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 iid is empty
// @router /:iid [delete]
func (i *IngredientController) Delete() {
	iid, _ := i.GetInt64(":iid")
	models.DeleteIngredient(iid)
	i.Data["json"] = "delete success!"
	i.ServeJSON()
}
