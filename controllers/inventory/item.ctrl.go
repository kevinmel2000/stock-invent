package inventory

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/lloistborn/stockinvent/models"
)

// GetItems return all array of items[] or status `internal server error`
func (inventoryModule *InventoryModule) GetItems(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := context.Background()
	itemModel := model.NewItemModel(ctx)

	items, err := itemModel.GetMany(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"Status" : "Internal Server error"}`))
		return
	}
	res, _ := json.Marshal(items)

	w.Write(res)
}
