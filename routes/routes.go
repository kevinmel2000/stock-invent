package routes

import (
	"github.com/julienschmidt/httprouter"
	"github.com/lloistborn/stockinvent/inventory"
)

// Routes object.
type Routes struct {
	Inventory *inventory.InventoryModule
}

// Route return all api endpoint from object Routes.
func Route() *Routes {
	return &Routes{
		Inventory: inventory.NewInventoryModule(),
	}
}

// RegisterAPI will register all available API Endpoint.
func (routes *Routes) RegisterAPI(r *httprouter.Router) {
	// Blouse REST endpoint.
	r.POST("/items/blouses", routes.Inventory.InsertBlouse)
	r.GET("/items/blouses", routes.Inventory.GetBlouses)
	r.GET("/items/blouses/:id", routes.Inventory.GetBlouse)
	r.PUT("/items/blouses/:id", routes.Inventory.UpdateBlouse)
	r.DELETE("/items/blouses/:id", routes.Inventory.DeleteBlouse)

	// Inbound Blouse Rest endpoint.
	r.POST("/items/inbound", routes.Inventory.InsertInboundBlouse)
	r.GET("/items/inbound/:id", routes.Inventory.GetInboundBlouse)
	r.GET("/items/inbound", routes.Inventory.GetInboundBlouses)
	r.PUT("/items/inbound/:id", routes.Inventory.UpdateInboundBlouse)
	r.DELETE("/items/inbound/:id", routes.Inventory.DeleteInboundBlouse)

	// TODO: implement Outbound Blouse Rest endpoint.

	// TODO: implement Value of Goods Rest endpoint.

	// TODO: implement Selling Report Rest endpoint.
}
