package walletB

import (
	"net/http"
	"strconv"
	"wallet/util"
)

func (h *Handler) GetTransactions(w http.ResponseWriter, r *http.Request) {
	val := r.Context().Value("user_id")
	var userID uint64
	switch v := val.(type) {
	case int:
		userID = uint64(v)
	case int64:
		userID = uint64(v)
	case uint64:
		userID = v
	default:
		util.SendError(w, http.StatusInternalServerError, "invalid user id")
		return
	}

	page, _ := strconv.ParseInt(r.URL.Query().Get("page"), 10, 64)
	if page <= 0 {
		page = 1
	}
	limit, _ := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 64)
	if limit <= 0 {
		limit = 10
	}

	txns, err := h.svc.GetTransactionHistory(r.Context(), userID)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	total := int64(len(txns))
	start := (page - 1) * limit
	end := start + limit
	if start > total {
		start = total
	}
	if end > total {
		end = total
	}
	paginatedTxns := txns[start:end]

	resp := map[string]interface{}{
		"count":        total,
		"page":         page,
		"limit":        limit,
		"transactions": paginatedTxns,
	}

	util.SendData(w, http.StatusOK, resp)
}
