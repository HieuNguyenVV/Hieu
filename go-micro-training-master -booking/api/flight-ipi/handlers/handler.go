package handlers

import (
	"gin-training/api/flight-ipi/requests"
	"gin-training/api/flight-ipi/responses"
	"gin-training/pb"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FlightHandler interface {
	CreateFlight(c *gin.Context)
	UpdateFlight(c *gin.Context)
	SearchFly(c *gin.Context)
}

type flightHandler struct {
	flightClient pb.FlightServiceClient
}

func NewFlightHandler(flightClient pb.FlightServiceClient) FlightHandler {
	return &flightHandler{
		flightClient: flightClient,
	}
}

func (h *flightHandler) CreateFlight(c *gin.Context) {
	req := requests.CreateFlightRequest{}

	if err := c.ShouldBind(&req); err != nil {
		if validateErrors, ok := err.(validator.ValidationErrors); ok {
			errMessages := make([]string, 0)
			for _, v := range validateErrors {
				errMessages = append(errMessages, v.Error())
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusText(http.StatusBadRequest),
				"error":  errMessages,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})

		return
	}

	pReq := &pb.Flight{
		Name:          req.Name,
		From:          req.From,
		To:            req.To,
		Date:          timestamppb.New(req.Date),
		Status:        req.Status,
		AvailableSlot: req.AvailableSlot,
	}

	// fReq := &pb.Flight{}
	// err := copier.Copy(&pReq, &req)
	// if err != nil {
	// 	return
	// }
	pRes, err := h.flightClient.CreateFlight(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.FlightResponse{
		Id:            pRes.Id,
		Name:          pRes.Name,
		From:          pRes.From,
		To:            pRes.To,
		Date:          pRes.Date.AsTime(),
		Status:        pRes.Status,
		AvailableSlot: pRes.AvailableSlot,
	}
	// dto := &responses.FlightResponse{}
	// err = copier.Copy(&dto, &pRes)
	// if err != nil {
	// 	return
	// }
	c.JSON(http.StatusOK, gin.H{
		"payload": dto,
		"status":  http.StatusOK,
	})
}

func (h *flightHandler) UpdateFlight(c *gin.Context) {
	req := requests.CreateFlightRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		if validatoErrors, ok := err.(validator.ValidationErrors); ok {
			errMessages := make([]string, 0)
			for _, v := range validatoErrors {
				errMessages = append(errMessages, v.Error())
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusText(http.StatusBadRequest),
				"error":  errMessages,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})
		return
	}

	pReq := pb.Flight{
		Date: timestamppb.New(req.Date),
	}

	err := copier.Copy(&pReq, &req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	pRes, err := h.flightClient.UpdateFlight(c.Request.Context(), &pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.FlightResponse{
		Date: pRes.Date.AsTime(),
	}

	err = copier.Copy(&dto, &pRes)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"payload": dto,
		"status":  http.StatusOK,
	})
}

func (h *flightHandler) SearchFly(c *gin.Context) {
	req := requests.SearchFlightRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		if validatoErrors, ok := err.(validator.ValidationErrors); ok {
			errMessages := make([]string, 0)
			for _, v := range validatoErrors {
				errMessages = append(errMessages, v.Error())
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusText(http.StatusBadRequest),
				"error":  errMessages,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})
		return
	}

	pReq := pb.SearchFlightRequest{}

	err := copier.Copy(&pReq, &req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	pRes, err := h.flightClient.SearchFlight(c.Request.Context(), &pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.ListFlights{}

	err = copier.CopyWithOption(&dto, &pRes, copier.Option{
		IgnoreEmpty: true,
		DeepCopy:    true,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"payload": dto,
		"status":  http.StatusOK,
	})
}
