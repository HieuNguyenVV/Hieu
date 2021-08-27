package handlers

import (
	"gin-training/api/booking-ipi/requests"
	"gin-training/api/booking-ipi/responses"
	"gin-training/pb"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
)

type BookingHandler interface {
	Booking(c *gin.Context)
	ViewBooking(c *gin.Context)
	CancelBooking(c *gin.Context)
}

type bookingHandler struct {
	bookingClient pb.BookingsClient
}

func NewBookingHandler(bookingClient pb.BookingsClient) BookingHandler {
	return &bookingHandler{
		bookingClient: bookingClient,
	}
}

func (h *bookingHandler) Booking(c *gin.Context) {
	req := requests.CreateBookingRequest{}
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

	pReq := &pb.Info{}
	err := copier.Copy(&pReq, &req)
	if err != nil {
		return
	}

	pRes, err := h.bookingClient.Booking(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.CreateBookingResponses{}
	err = copier.Copy(&dto, &pRes)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}
func (h *bookingHandler) ViewBooking(c *gin.Context) {
	req := requests.ViewRequest{}
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

	pReq := &pb.ViewRequest{}
	err := copier.Copy(&pReq, &req)
	if err != nil {
		return
	}

	pRes, err := h.bookingClient.ViewBooking(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.ViewResponse{}
	err = copier.Copy(&dto, &pRes)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}
func (h *bookingHandler) CancelBooking(c *gin.Context) {
	req := requests.ViewRequest{}
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

	pReq := &pb.ViewRequest{}
	err := copier.Copy(&pReq, &req)
	if err != nil {
		return
	}

	pRes, err := h.bookingClient.CancelBooking(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.Cancel{}
	err = copier.Copy(&dto, &pRes)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}
