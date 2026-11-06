package handler

import (
	"net/http"
	"strconv"

	"github.com/example/person-crud-go/internal/entity"
	"github.com/example/person-crud-go/internal/service"
	"github.com/gin-gonic/gin"
)

type PersonHandler struct {
	service *service.PersonService
}

func NewPersonHandler(svc *service.PersonService) *PersonHandler {
	return &PersonHandler{service: svc}
}

func (h *PersonHandler) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/persons")
	{
		api.GET("", h.GetAll)
		api.GET("/:id", h.GetByID)
		api.POST("", h.Create)
		api.PUT("/:id", h.Update)
		api.DELETE("/:id", h.Delete)
	}
}

func (h *PersonHandler) GetAll(c *gin.Context) {
	persons, err := h.service.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, persons)
}

func (h *PersonHandler) GetByID(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	person, err := h.service.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "person not found"})
		return
	}
	c.JSON(http.StatusOK, person)
}

func (h *PersonHandler) Create(c *gin.Context) {
	var person entity.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.Save(&person); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, person)
}

func (h *PersonHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	var person entity.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	person.ID = uint(id)
	if err := h.service.Save(&person); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, person)
}

func (h *PersonHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
