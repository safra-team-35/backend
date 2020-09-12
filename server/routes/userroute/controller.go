package userroute

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/IQ-tech/go-mapper"
	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/gin-gonic/gin"
	"github.com/safra-team-35/backend/domain/contract"
	"github.com/safra-team-35/backend/domain/entity"
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

func (s *Controller) handleCreateNewOrder(c *gin.Context) {

	userID, parseErr := strconv.Atoi(c.Param("user_id"))
	if parseErr != nil {
		err := resterrors.NewBadRequestError("user_id parameter is invalid")
		c.JSON(err.StatusCode(), err)
	}

	input := viewmodel.Order{}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		logger.Error("handleCreateNewOrder", err)
		restErr := resterrors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.StatusCode(), restErr)
		return
	}

	order := entity.Order{}

	mapErr := s.mapper.From(input).To(&order)
	if mapErr != nil {
		restErr := resterrors.NewInternalServerError("Error to do the mapper: " + fmt.Sprint(mapErr))
		c.JSON(restErr.StatusCode(), restErr)
	}

	order.UserID = int64(userID)

	orderNumber, createErr := s.userService.CreateOrder(order)
	if createErr != nil {
		c.JSON(createErr.StatusCode(), createErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{"orderNumber": orderNumber})
}
