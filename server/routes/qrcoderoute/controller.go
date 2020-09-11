package qrcoderoute

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/IQ-tech/go-mapper"
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
	qrcodeService contract.QRCodeService
	mapper        mapper.Mapper
}

//NewController to handle requests
func NewController(qrcodeService contract.QRCodeService, mapper mapper.Mapper) *Controller {
	once.Do(func() {
		instance = &Controller{
			qrcodeService: qrcodeService,
			mapper:        mapper,
		}
	})
	return instance
}

// handlePing - handle a qrcode create request
func (s *Controller) handleCreateQRCode(c *gin.Context) {

	var input viewmodel.QRCode

	err := c.ShouldBindJSON(&input)
	if err != nil {
		log.Println(err)
		restErr := resterrors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.StatusCode(), restErr)
		return
	}
	uuid := c.Param("identifier_id")

	var QRCode entity.QRCode

	mapErr := s.mapper.From(input).To(&QRCode)
	if mapErr != nil {
		restErr := resterrors.NewInternalServerError("Error to do the mapper: " + fmt.Sprint(mapErr))
		c.JSON(restErr.StatusCode(), restErr)
	}

	expirationMinutes := time.Duration(input.ExpirationTime) * time.Minute
	expirationTime := time.Now().UTC().Add(expirationMinutes)
	QRCode.ExpirationTime = expirationTime.Format("2006-01-02T15:04:05")

	hash, createErr := s.qrcodeService.CreateQRCode(QRCode, uuid)
	if createErr != nil {
		c.JSON(createErr.StatusCode(), createErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{"hash": hash})
}
