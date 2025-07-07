package controllers

import (
	"github.com/alirezasaharkhiz/FlagChain/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
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
		if strings.Contains(err.Error(), "EOF") {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Request body is required"})
		} else {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		}
		return
	}
	flag, err := c.Svc.CreateFlag(req.Name, req.Dependencies, "system")
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "dependency flag not found"})
		} else {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		}
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

func (c *FeatureFlagController) List(ctx *gin.Context) {
	flags, err := c.Svc.ListFlags()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, flags)
}

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

func (c *FeatureFlagController) AddDependency(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid flag id"})
		return
	}

	var req struct {
		DependsOnID uint `json:"depends_on_id" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		if strings.Contains(err.Error(), "EOF") {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "request body is required"})
		} else {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		}
		return
	}

	err = c.Svc.AddDependency(uint(id), req.DependsOnID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "dependency added successfully"})
}
