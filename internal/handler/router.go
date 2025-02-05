package handler

import "net/http"

func (myHandler *MyHandler) InitRouter() http.Handler {

	Middleware := func(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
		return LoggingMiddleware(PanicMiddleware(http.HandlerFunc(next)))
	}
	mux := http.NewServeMux()
	mux.Handle("/", Middleware(myHandler.home))
	mux.Handle("/orders", Middleware(myHandler.order))
	mux.Handle("/menu", Middleware(myHandler.menu))
	mux.Handle("/inventory", Middleware(myHandler.inventory))
	mux.Handle("/orders/", Middleware(myHandler.specificOrder))
	mux.Handle("/menu/", Middleware(myHandler.specificMenu))
	mux.Handle("/inventory/", Middleware(myHandler.specificInventory))
	mux.Handle("/reports/", Middleware(myHandler.reports))

	return http.HandlerFunc(mux.ServeHTTP)
}
