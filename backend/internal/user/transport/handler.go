package transport

import "github.com/iiincognito/CRM-detailing-studio/backend/internal/user/service"

type Handler struct {
	service service.ServInterface
}
