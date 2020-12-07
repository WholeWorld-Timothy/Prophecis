// Code generated by go-swagger; DO NOT EDIT.

package roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetRoleByIDHandlerFunc turns a function with the right signature into a get role by Id handler
type GetRoleByIDHandlerFunc func(GetRoleByIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRoleByIDHandlerFunc) Handle(params GetRoleByIDParams) middleware.Responder {
	return fn(params)
}

// GetRoleByIDHandler interface for that can handle valid get role by Id params
type GetRoleByIDHandler interface {
	Handle(GetRoleByIDParams) middleware.Responder
}

// NewGetRoleByID creates a new http.Handler for the get role by Id operation
func NewGetRoleByID(ctx *middleware.Context, handler GetRoleByIDHandler) *GetRoleByID {
	return &GetRoleByID{Context: ctx, Handler: handler}
}

/*GetRoleByID swagger:route GET /cc/v1/roles/id/{id} Roles getRoleById

get roles.

Optional extended description in Markdown.

*/
type GetRoleByID struct {
	Context *middleware.Context
	Handler GetRoleByIDHandler
}

func (o *GetRoleByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetRoleByIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}