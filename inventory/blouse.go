package inventory

import (
	"context"
	"net/http"

	"github.com/asaskevich/govalidator"

	"github.com/julienschmidt/httprouter"
	"github.com/lloistborn/stockinvent/dao"
	"github.com/lloistborn/stockinvent/utils"
	"github.com/mholt/binding"
)

// BlouseForm object.
type BlouseForm struct {
	SKU   string `valid:"required"`
	Name  string `valid:"required"`
	Stock int    `valid:"required"`
}

// FieldMap implemented for BlouseForm.
func (bf *BlouseForm) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&bf.Name:  "name",
		&bf.SKU:   "sku",
		&bf.Stock: "stock",
	}
}

// InsertBlouse used to insert one record to table.
func (inventoryModule *InventoryModule) InsertBlouse(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := context.Background()

	blouseform := new(BlouseForm)
	util := utils.NewUtilsModule(ctx)

	if err := binding.Bind(r, blouseform); err != nil {
		util.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	valid, err := govalidator.ValidateStruct(blouseform)
	if !valid && err != nil {
		util.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	blouseDao := dao.NewBlouseDao(ctx)

	blouse := dao.Blouse{
		Name:  blouseform.Name,
		SKU:   blouseform.SKU,
		Stock: blouseform.Stock,
	}
	if err := blouseDao.Insert(ctx, blouse); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

// GetBlouses return all array of blouse[] or status `internal server error`
func (inventoryModule *InventoryModule) GetBlouses(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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
