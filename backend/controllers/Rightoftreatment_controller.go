package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/benyabooncharoen/app/ent"
	"github.com/benyabooncharoen/app/ent/rightoftreatment"
)

type RightoftreatmentController struct {
	client *ent.Client
	router gin.IRouter
}

type Rightoftreatment struct {
	RightoftreatmentName	 string
}

// CreateRightoftreatment handles POST requests for adding rightoftreatment entities
// @Summary Create rightoftreatment
// @Description Create rightoftreatment
// @ID create-rightoftreatment
// @Accept   json
// @Produce  json
// @Param rightoftreatment body ent.Rightoftreatment true "Rightoftreatment entity"
// @Success 200 {object} ent.Rightoftreatment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /rightoftreatments [post]
func (ctl *RightoftreatmentController) CreateRightoftreatment(c *gin.Context) {
	obj := Rightoftreatment{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "rightoftreatment binding failed",
		})
		return
	}

	r, err := ctl.client.Rightoftreatment.
		Create().
		SetRightoftreatmentName(obj.RightoftreatmentName).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, r)
}

// GetRightoftreatment handles GET requests to retrieve a rightoftreatment entity
// @Summary Get a rightoftreatment entity by ID
// @Description get rightoftreatment by ID
// @ID get-rightoftreatment
// @Produce  json
// @Param id path int true "Rightoftreatment ID"
// @Success 200 {object} ent.Rightoftreatment
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /rightoftreatments/{id} [get]
func (ctl *RightoftreatmentController) GetRightoftreatment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	r, err := ctl.client.Rightoftreatment.
		Query().
		Where(rightoftreatment.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, r)
}

// ListRightoftreatment handles request to get a list of rightoftreatment entities
// @Summary List rightoftreatment entities
// @Description list rightoftreatment entities
// @ID list-rightoftreatment
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Rightoftreatment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /rightoftreatments [get]
func (ctl *RightoftreatmentController) ListRightoftreatment(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	rightoftreatments, err := ctl.client.Rightoftreatment.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, rightoftreatments)
}

// DeleteRightoftreatment handles DELETE requests to delete a rightoftreatment entity
// @Summary Delete a rightoftreatment entity by ID
// @Description get rightoftreatment by ID
// @ID delete-rightoftreatment
// @Produce  json
// @Param id path int true "Rightoftreatment ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /rightoftreatment/{id} [delete]
func (ctl *RightoftreatmentController) DeleteRightoftreatment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Rightoftreatment.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// NewRightoftreatmentController creates and registers handles for the rightoftreatment controller
func NewRightoftreatmentController(router gin.IRouter, client *ent.Client) *RightoftreatmentController {
	rc := &RightoftreatmentController{
		client: client,
		router: router,
	}

	rc.register()

	return rc

}

func (ctl *RightoftreatmentController) register() {
	rightoftreatments := ctl.router.Group("/rightoftreatments")

	rightoftreatments.POST("", ctl.CreateRightoftreatment)
	rightoftreatments.GET(":id", ctl.GetRightoftreatment)
	rightoftreatments.GET("", ctl.ListRightoftreatment)
	rightoftreatments.DELETE("", ctl.DeleteRightoftreatment)

}