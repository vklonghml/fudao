package router

import (
	"github.com/emicklei/go-restful"
)

func Register() {
	ws := new(restful.Web)
	apiVersion := "/v1"
	ws.Path(apiVersion).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	registerRoutes(ws)
	restful.Add(ws)
}

func registerRoutes(ws *restful.WebService) {
	ws.Route(ws.GET("/courses/{course_id}").To(getOne)).Operation("get one course")
	ws.Route(ws.GET("/courses").To(getAll)).Operation("get one course")
}

func getOne(req *restful.Request, resp *restful.Response) {

}

func getAll(req *restful.Request, resp *restful.Response) {

}
