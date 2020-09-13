package safraroute

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/safra-team-35/backend/domain/contract"
)

var (
	instance *Controller
	once     sync.Once
)

//Controller holds user handlers
type Controller struct {
	safraService contract.SafraService
}

//NewController to handle requests
func NewController(safraService contract.SafraService) *Controller {
	once.Do(func() {
		instance = &Controller{
			safraService: safraService,
		}
	})
	return instance
}

func (s *Controller) handleGetSafraMorningCalls(c *gin.Context) {

	response, err := s.safraService.GetMorningCalls()
	if err != nil {
		c.JSON(err.StatusCode(), err)
		return
	}

	c.JSON(http.StatusOK, response)

}
