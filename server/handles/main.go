package handles

import (
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"

	"github.com/petaki/inertia-go"
)

type Handler struct {
	inertiaManager *inertia.Inertia
}

func NewHandler(inertiaManager *inertia.Inertia) *Handler {
	return &Handler{
		inertiaManager: inertiaManager,
	}
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		fmt.Fprint(w, "404 Not found")
		return
	}

	err := h.inertiaManager.Render(w, r, "home/Index", map[string]interface{}{
		"random": rand.IntN(500),
	})
	if err != nil {
		log.Panicln(err)
	}
}

func (h *Handler) About(w http.ResponseWriter, r *http.Request) {
	err := h.inertiaManager.Render(w, r, "about/Index", nil)
	if err != nil {
		log.Panicln(err)
	}
}
