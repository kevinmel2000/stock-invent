package inventory

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/lloistborn/stockinvent/dao"
	"github.com/lloistborn/stockinvent/utils"
)

// GetBlouse return all array of blouse[] or status `internal server error`
func (inventoryModule *InventoryModule) GetBlouse(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := context.Background()
	blouseDao := dao.NewBlouseDao(ctx)
	util := utils.NewUtilsModule(ctx)

	blouse, err := blouseDao.FindAll(ctx)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusOK, blouse)
}
