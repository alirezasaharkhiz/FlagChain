package controllers

import (
	"net/http"
	"strconv"

	"github.com/alirezasaharkhiz/FlagChain/services"
	"github.com/gin-gonic/gin"
)

type FeatureFlagController struct {
	Svc *services.FeatureFlagService
}

func NewFeatureFlagController(s *services.FeatureFlagService) *FeatureFlagController {
	return &FeatureFlagController{Svc: s}
}

func (c *FeatureFlagController) Create(ctx *gin.Context) {
	var req struct {
		Name         string   `json:"name" binding:"required"`
		Dependencies []string `json:"dependencies"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	flag, err := c.Svc.CreateFlag(req.Name, req.Dependencies, "system")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, flag)
}

func (c *FeatureFlagController) Toggle(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	flag, err := c.Svc.ToggleFlag(uint(id), "system")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, flag)
}

// List handles GET /flags
func (c *FeatureFlagController) List(ctx *gin.Context) {
	flags, err := c.Svc.ListFlags()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, flags)
}

// History handles GET /flags/:id/history
func (c *FeatureFlagController) History(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	hist, err := c.Svc.GetHistory(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, hist)
}
