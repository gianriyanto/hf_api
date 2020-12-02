package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["hf_api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["hf_api/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["hf_api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["hf_api/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["hf_api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["hf_api/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           "/:objectId",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["hf_api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["hf_api/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:objectId",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["hf_api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["hf_api/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:objectId",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["hf_api/controllers:RecipeController"] = append(beego.GlobalControllerRouter["hf_api/controllers:RecipeController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["hf_api/controllers:RecipeController"] = append(beego.GlobalControllerRouter["hf_api/controllers:RecipeController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           "/:rid",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["hf_api/controllers:RecipeController"] = append(beego.GlobalControllerRouter["hf_api/controllers:RecipeController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:rid",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["hf_api/controllers:RecipeController"] = append(beego.GlobalControllerRouter["hf_api/controllers:RecipeController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:rid",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["hf_api/controllers:RecipeController"] = append(beego.GlobalControllerRouter["hf_api/controllers:RecipeController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/create",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["hf_api/controllers:UserController"] = append(beego.GlobalControllerRouter["hf_api/controllers:UserController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["hf_api/controllers:UserController"] = append(beego.GlobalControllerRouter["hf_api/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           "/:uid",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["hf_api/controllers:UserController"] = append(beego.GlobalControllerRouter["hf_api/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:uid",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["hf_api/controllers:UserController"] = append(beego.GlobalControllerRouter["hf_api/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:uid",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["hf_api/controllers:UserController"] = append(beego.GlobalControllerRouter["hf_api/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/create",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
