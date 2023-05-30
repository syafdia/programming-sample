package internal

import "net/http"

type HTTPHandler struct {
	appManagementUC AppManagementUseCase
	authUC          AuthUseCase
}

func NewHTTPHandler(
	appManagementUC AppManagementUseCase,
	authUC AuthUseCase,
) *HTTPHandler {
	return &HTTPHandler{
		appManagementUC: appManagementUC,
		authUC:          authUC,
	}
}

func (h *HTTPHandler) DestroyApplication(w http.ResponseWriter, r *http.Request) {
	// TODO
}
