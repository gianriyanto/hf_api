package controllers

import (
	"encoding/json"
	"hf_api/models"

	"github.com/astaxie/beego"
)

// Operations about Recipe
type RecipeController struct {
	beego.Controller
}

// @Title CreateRecipe
// @Description create recipes
// @Param	body		body 	models.Recipe	true		"body for recipe content"
// @Success 200 {int} models.Recipe.Id
// @Failure 403 body is empty
// @router /create [post]
func (r *RecipeController) Post() {
	var recipe models.Recipe
	json.Unmarshal(r.Ctx.Input.RequestBody, &recipe)
	rid := models.AddRecipe(recipe)
	r.Data["json"] = map[string]int64{"rid": rid}
	r.ServeJSON()
}

// @Title GetAll
// @Description get all Recipe
// @Success 200 {object} models.Recipe
// @router / [get]
func (r *RecipeController) GetAll() {
	recipe := models.GetAllRecipes()
	r.Data["json"] = recipe
	r.ServeJSON()
}

// @Title Get
// @Description get recipe by rid
// @Param	rid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Recipe
// @Failure 403 :rid is empty
// @router /:rid [get]
func (r *RecipeController) Get() {
	rid, _ := r.GetInt64(":rid")
	recipe, err := models.GetRecipe(rid)
	if err != nil {
		r.Data["json"] = err.Error()
	} else {
		r.Data["json"] = recipe
	}
	r.ServeJSON()
}

// @Title Update
// @Description update the recipe
// @Param	rid		path 	string	true		"The rid you want to update"
// @Param	body		body 	models.Recipe	true		"body for recipe content"
// @Success 200 {object} models.Recipe
// @Failure 403 :rid is not int
// @router /:rid [put]
func (r *RecipeController) Put() {
	rid, _ := r.GetInt64(":rid")
	var recipe models.Recipe
	json.Unmarshal(r.Ctx.Input.RequestBody, &recipe)
	rr, err := models.UpdateRecipe(rid, &recipe)
	if err != nil {
		r.Data["json"] = err.Error()
	} else {
		r.Data["json"] = rr
	}
	r.ServeJSON()
}

// @Title Delete
// @Description delete the recipe
// @Param	rid		path 	string	true		"The rid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 rid is empty
// @router /:rid [delete]
func (r *RecipeController) Delete() {
	rid, _ := r.GetInt64(":rid")
	models.DeleteRecipe(rid)
	r.Data["json"] = "delete success!"
	r.ServeJSON()
}
