package walletB

import (
	"net/http"

	middleware "wallet/rest/middlewares"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /wallets/my-balance",
		manager.With(
			http.HandlerFunc(h.GetBalance),
			h.middlewares.AuthenticateJWT,
		),
	)

	mux.Handle("POST /wallets/topup",
		manager.With(
			http.HandlerFunc(h.TopUp),
			h.middlewares.AuthenticateJWT,
		),
	)

	mux.Handle("POST /wallets/transfer",
		manager.With(
			http.HandlerFunc(h.Transfer),
			h.middlewares.AuthenticateJWT,
		),
	)

	mux.Handle("GET /wallets/transactions",
		manager.With(
			http.HandlerFunc(h.GetTransactions),
			h.middlewares.AuthenticateJWT,
		),
	)
}
