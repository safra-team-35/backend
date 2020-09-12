package userroute

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/IQ-tech/go-mapper"
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/gin-gonic/gin"
	"github.com/safra-team-35/backend/domain/contract"
	"github.com/safra-team-35/backend/server/viewmodel"
)

var (
	instance *Controller
	once     sync.Once
)

//Controller is a interface to interact with services
type Controller struct {
	userService contract.UserService
	mapper      mapper.Mapper
}

//NewController to handle requests
func NewController(userService contract.UserService, mapper mapper.Mapper) *Controller {
	once.Do(func() {
		instance = &Controller{
			userService: userService,
			mapper:      mapper,
		}
	})
	return instance
}

func (s *Controller) handleGetUserAddress(c *gin.Context) {
	userID, parseErr := strconv.Atoi(c.Param("user_id"))
	if parseErr != nil {
		err := resterrors.NewBadRequestError("user_id parameter is invalid")
		c.JSON(err.StatusCode(), err)
	}

	userAddresses, err := s.userService.GetUserAddress(int64(userID))
	if err != nil {
		c.JSON(err.StatusCode(), err)
		return
	}

	response := []viewmodel.Address{}
	mapErr := s.mapper.From(userAddresses).To(&response)
	if mapErr != nil {
		err = resterrors.NewInternalServerError("Error to do the mapper: " + fmt.Sprint(mapErr))
		c.JSON(err.StatusCode(), err)
	}

	c.JSON(http.StatusOK, response)
}
