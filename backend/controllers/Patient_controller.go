package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/benyabooncharoen/app/ent"
	"github.com/benyabooncharoen/app/ent/gender"
	"github.com/benyabooncharoen/app/ent/rightoftreatment"
	"github.com/benyabooncharoen/app/ent/systemmember"
	"github.com/gin-gonic/gin"
)

type PatientController struct {
	client *ent.Client
	router gin.IRouter
}

type Patient struct {
	HN               string
	PatientName      string
	Gender           int
	Rightoftreatment int
	Systemmember     int
}

// CreatePatient handles POST requests for adding patient entities
// @Summary Create patient
// @Description Create patient
// @ID create-patient
// @Accept   json
// @Produce  json
// @Param patient body Patient true "Patient entity"
// @Success 200 {object} Patient
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patients [post]
func (ctl *PatientController) CreatePatient(c *gin.Context) {
	obj := Patient{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "patient binding failed",
		})
		return
	}

	g, err := ctl.client.Gender.
		Query().
		Where(gender.IDEQ(int(obj.Gender))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "gender not found",
		})
		return
	}

	r, err := ctl.client.Rightoftreatment.
		Query().
		Where(rightoftreatment.IDEQ(int(obj.Rightoftreatment))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "rightoftreatment not found",
		})
		return
	}

	sm, err := ctl.client.Systemmember.
		Query().
		Where(systemmember.IDEQ(int(obj.Systemmember))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "systemmember not found",
		})
		return
	}

	p, err := ctl.client.Patient.
		Create().
		SetHn(obj.HN).
		SetPatientName(obj.PatientName).
		SetGender(g).
		SetRightoftreatment(r).
		SetSystemmember(sm).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, p)
}

// DeletePatient handles DELETE requests to delete a patient entity
// @Summary Delete a patient entity by ID
// @Description get patient by ID
// @ID delete-patient
// @Produce  json
// @Param id path int true "Patient ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patients/{id} [delete]
func (ctl *PatientController) DeletePatient(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Patient.
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

// ListPatient handles request to get a list of patient entities
// @Summary List patient entities
// @Description list patient entities
// @ID list-patient
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Patient
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patients [get]
func (ctl *PatientController) ListPatient(c *gin.Context) {
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

	patients, err := ctl.client.Patient.
		Query().
		WithGender().
		WithRightoftreatment().
		WithSystemmember().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, patients)
}

// NewPatientController creates and registers handles for the patient controller
func NewPatientController(router gin.IRouter, client *ent.Client) *PatientController {
	pic := &PatientController{
		client: client,
		router: router,
	}

	pic.register()

	return pic

}

func (ctl *PatientController) register() {
	patients := ctl.router.Group("/patients")

	patients.POST("", ctl.CreatePatient)
	patients.GET("", ctl.ListPatient)
	patients.DELETE(":id", ctl.DeletePatient)

}
