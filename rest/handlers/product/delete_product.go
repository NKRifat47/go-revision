package product

import (
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")
	pId, err := strconv.Atoi(productID)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	err = h.productRepo.Delete(pId)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	util.SendData(w, http.StatusOK, "Successfully Deleted Product.")
}
