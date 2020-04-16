// Code generated by go-swagger; DO NOT EDIT.

package flag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetFlagEntityTypesHandlerFunc turns a function with the right signature into a get flag entity types handler
type GetFlagEntityTypesHandlerFunc func(GetFlagEntityTypesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetFlagEntityTypesHandlerFunc) Handle(params GetFlagEntityTypesParams) middleware.Responder {
	return fn(params)
}

// GetFlagEntityTypesHandler interface for that can handle valid get flag entity types params
type GetFlagEntityTypesHandler interface {
	Handle(GetFlagEntityTypesParams) middleware.Responder
}

// NewGetFlagEntityTypes creates a new http.Handler for the get flag entity types operation
func NewGetFlagEntityTypes(ctx *middleware.Context, handler GetFlagEntityTypesHandler) *GetFlagEntityTypes {
	return &GetFlagEntityTypes{Context: ctx, Handler: handler}
}

/*GetFlagEntityTypes swagger:route GET /flags/entity_types flag getFlagEntityTypes

GetFlagEntityTypes get flag entity types API

*/
type GetFlagEntityTypes struct {
	Context *middleware.Context
	Handler GetFlagEntityTypesHandler
}

func (o *GetFlagEntityTypes) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetFlagEntityTypesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
