package inventory

import (
	"context"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/lloistborn/stockinvent/dao"
	"github.com/lloistborn/stockinvent/utils"
	"github.com/mholt/binding"

	"strconv"

	"github.com/julienschmidt/httprouter"
)

// InboundBlouseForm object.
type InboundBlouseForm struct {
	BlouseID       int    `valid:"required"`
	SKU            string `valid:"required"`
	Name           string `valid:"required"`
	OrderAmount    int    `valid:"required"`
	ReceivedAmount int    `valid:"required"`
	PurchasePrice  int    `valid:"required"`
	TotalPrice     int    `valid:"required"`
	ReceiptNumber  string `valid:"required"`
	Notes          string `valid:"required"`
}

// FieldMap implemented for InboundBlouseForm.
func (ibf *InboundBlouseForm) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&ibf.BlouseID:       "blouseId",
		&ibf.SKU:            "sku",
		&ibf.Name:           "name",
		&ibf.OrderAmount:    "orderAmount",
		&ibf.ReceivedAmount: "receivedAmount",
		&ibf.PurchasePrice:  "purchasePrice",
		&ibf.TotalPrice:     "totalPrice",
		&ibf.ReceiptNumber:  "receiptNumber",
		&ibf.Notes:          "notes",
	}
}

// InsertInboundBlouse will insert a record into InboundBlouse and update Blouse table.
func (inventoryModule *InventoryModule) InsertInboundBlouse(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := context.Background()

	util := utils.NewUtilsModule(ctx)
	inboundblouseform := new(InboundBlouseForm)

	if err := binding.Bind(r, inboundblouseform); err != nil {
		util.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	valid, err := govalidator.ValidateStruct(inboundblouseform)
	if !valid && err != nil {
		util.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	inboundBlouseDao := dao.NewInboundBlouseDao(ctx)

	inboundBlouse := dao.InboundBlouse{
		BlouseID:       inboundblouseform.BlouseID,
		SKU:            inboundblouseform.SKU,
		Name:           inboundblouseform.Name,
		OrderAmount:    inboundblouseform.OrderAmount,
		ReceivedAmount: inboundblouseform.ReceivedAmount,
		PurchasePrice:  inboundblouseform.PurchasePrice,
		TotalPrice:     inboundblouseform.TotalPrice,
		ReceiptNumber:  inboundblouseform.ReceiptNumber,
		Notes:          inboundblouseform.Notes,
	}

	if err := inboundBlouseDao.Insert(ctx, inboundBlouse); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// update stock of Blouse
	blouseDao := dao.NewBlouseDao(ctx)
	blouse, _ := blouseDao.FindOne(ctx, inboundBlouse.BlouseID)
	blouse.Stock += inboundBlouse.ReceivedAmount
	if err = blouseDao.Update(ctx, blouse.ID, blouse); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "success"})
}

// GetInboundBlouse return single record.
func (inventoryModule *InventoryModule) GetInboundBlouse(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := context.Background()
	dao := dao.NewInboundBlouseDao(ctx)
	util := utils.NewUtilsModule(ctx)

	id, _ := strconv.Atoi(p.ByName("id"))

	inboundBlouse, err := dao.FindOne(ctx, id)
	if err != nil {
		util.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusOK, map[string]interface{}{"status": "success", "result": inboundBlouse})
}

// GetInboundBlouses return all record.
func (inventoryModule *InventoryModule) GetInboundBlouses(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := context.Background()
	dao := dao.NewInboundBlouseDao(ctx)
	util := utils.NewUtilsModule(ctx)

	inboundBlouses, err := dao.FindAll(ctx)
	if err != nil {
		util.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusOK, map[string]interface{}{"status": "success", "result": inboundBlouses})
}

// UpdateInboundBlouse update record by id.
func (inventoryModule *InventoryModule) UpdateInboundBlouse(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := context.Background()

	inboundBlouseForm := new(InboundBlouseForm)
	util := utils.NewUtilsModule(ctx)

	if err := binding.Bind(r, inboundBlouseForm); err != nil {
		util.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	valid, err := govalidator.ValidateStruct(inboundBlouseForm)
	if !valid && err != nil {
		util.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	inboundBlouseDao := dao.NewInboundBlouseDao(ctx)

	id, _ := strconv.Atoi(p.ByName("id"))

	inboundBlouse := dao.InboundBlouse{
		ID:             id,
		BlouseID:       inboundBlouseForm.BlouseID,
		SKU:            inboundBlouseForm.SKU,
		Name:           inboundBlouseForm.Name,
		OrderAmount:    inboundBlouseForm.OrderAmount,
		ReceivedAmount: inboundBlouseForm.ReceivedAmount,
		PurchasePrice:  inboundBlouseForm.PurchasePrice,
		TotalPrice:     inboundBlouseForm.TotalPrice,
		ReceiptNumber:  inboundBlouseForm.ReceiptNumber,
		Notes:          inboundBlouseForm.Notes,
	}

	if err := inboundBlouseDao.Update(ctx, id, inboundBlouse); err != nil {
		util.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "success"})
}

// DeleteInboundBlouse delete single record by id.
func (inventoryModule *InventoryModule) DeleteInboundBlouse(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := context.Background()

	util := utils.NewUtilsModule(ctx)
	inboundBlouseDao := dao.NewInboundBlouseDao(ctx)

	id, _ := strconv.Atoi(p.ByName("id"))

	if err := inboundBlouseDao.Delete(ctx, id); err != nil {
		util.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "success"})
}
