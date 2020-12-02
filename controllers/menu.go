package controllers

import (
	"encoding/json"
	"hf_api/models"

	"github.com/astaxie/beego"
)

// Operations about object
type MenuController struct {
	beego.Controller
}

// @Title CreateMenu
// @Description create menus
// @Param	body		body 	models.Menu	true		"body for menu content"
// @Success 200 {int} models.Menu.Id
// @Failure 403 body is empty
// @router /create [post]
func (m *MenuController) Post() {
	var menu models.Menu
	json.Unmarshal(m.Ctx.Input.RequestBody, &menu)
	mid := models.AddMenu(menu)
	m.Data["json"] = map[string]int64{"mid": mid}
	m.ServeJSON()
}

// @Title GetMenu
// @Description get all Menus
// @Success 200 {object} models.Menu
// @router / [get]
func (m *MenuController) GetAll() {
	users := models.GetAllMenus()
	m.Data["json"] = users
	m.ServeJSON()
}

// @Title Get
// @Description get user by mid
// @Param	mid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Menu
// @Failure 403 :mid is empty
// @router /:mid [get]
func (m *MenuController) Get() {
	mid, _ := m.GetInt64(":mid")
	menu, err := models.GetMenu(mid)
	if err != nil {
		m.Data["json"] = err.Error()
	} else {
		m.Data["json"] = menu
	}
	m.ServeJSON()
}

// @Title Update
// @Description update the menu
// @Param	mid		path 	string	true		"The mid you want to update"
// @Param	body		body 	models.Menu	true		"body for menu content"
// @Success 200 {object} models.Menu
// @Failure 403 :mid is not int
// @router /:mid [put]
func (m *MenuController) Put() {
	mid, _ := m.GetInt64(":mid")
	var menu models.Menu
	json.Unmarshal(m.Ctx.Input.RequestBody, &menu)
	mm, err := models.UpdateMenu(mid, &menu)
	if err != nil {
		m.Data["json"] = err.Error()
	} else {
		m.Data["json"] = mm
	}
	m.ServeJSON()
}

// @Title Delete
// @Description delete the menu
// @Param	mid		path 	string	true		"The mid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 mid is empty
// @router /:mid [delete]
func (m *MenuController) Delete() {
	mid, _ := m.GetInt64(":uid")
	models.DeleteMenu(mid)
	m.Data["json"] = "delete success!"
	m.ServeJSON()
}
