package pool

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ReturnVMHandlerFunc turns a function with the right signature into a return VM handler
type ReturnVMHandlerFunc func(ReturnVMParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReturnVMHandlerFunc) Handle(params ReturnVMParams) middleware.Responder {
	return fn(params)
}

// ReturnVMHandler interface for that can handle valid return VM params
type ReturnVMHandler interface {
	Handle(ReturnVMParams) middleware.Responder
}

// NewReturnVM creates a new http.Handler for the return VM operation
func NewReturnVM(ctx *middleware.Context, handler ReturnVMHandler) *ReturnVM {
	return &ReturnVM{Context: ctx, Handler: handler}
}

/*ReturnVM swagger:route POST /pool/returnvm Pool returnVm

Return a VM into the pool

Return a VM into the pool

*/
type ReturnVM struct {
	Context *middleware.Context
	Handler ReturnVMHandler
}

func (o *ReturnVM) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewReturnVMParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
