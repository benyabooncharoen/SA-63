// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// GendersColumns holds the columns for the "genders" table.
	GendersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "gender_name", Type: field.TypeString},
	}
	// GendersTable holds the schema information for the "genders" table.
	GendersTable = &schema.Table{
		Name:        "genders",
		Columns:     GendersColumns,
		PrimaryKey:  []*schema.Column{GendersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// PatientsColumns holds the columns for the "patients" table.
	PatientsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "hn", Type: field.TypeString, Unique: true},
		{Name: "patient_name", Type: field.TypeString},
		{Name: "Gender", Type: field.TypeInt, Nullable: true},
		{Name: "Rightoftreatment", Type: field.TypeInt, Nullable: true},
		{Name: "Systemmember", Type: field.TypeInt, Nullable: true},
	}
	// PatientsTable holds the schema information for the "patients" table.
	PatientsTable = &schema.Table{
		Name:       "patients",
		Columns:    PatientsColumns,
		PrimaryKey: []*schema.Column{PatientsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "patients_genders_patients",
				Columns: []*schema.Column{PatientsColumns[3]},

				RefColumns: []*schema.Column{GendersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "patients_rightoftreatments_patients",
				Columns: []*schema.Column{PatientsColumns[4]},

				RefColumns: []*schema.Column{RightoftreatmentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "patients_systemmembers_patients",
				Columns: []*schema.Column{PatientsColumns[5]},

				RefColumns: []*schema.Column{SystemmembersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// RightoftreatmentsColumns holds the columns for the "rightoftreatments" table.
	RightoftreatmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "rightoftreatment_name", Type: field.TypeString},
	}
	// RightoftreatmentsTable holds the schema information for the "rightoftreatments" table.
	RightoftreatmentsTable = &schema.Table{
		Name:        "rightoftreatments",
		Columns:     RightoftreatmentsColumns,
		PrimaryKey:  []*schema.Column{RightoftreatmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// SystemmembersColumns holds the columns for the "systemmembers" table.
	SystemmembersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "systemmember_name", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
	}
	// SystemmembersTable holds the schema information for the "systemmembers" table.
	SystemmembersTable = &schema.Table{
		Name:        "systemmembers",
		Columns:     SystemmembersColumns,
		PrimaryKey:  []*schema.Column{SystemmembersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		GendersTable,
		PatientsTable,
		RightoftreatmentsTable,
		SystemmembersTable,
	}
)

func init() {
	PatientsTable.ForeignKeys[0].RefTable = GendersTable
	PatientsTable.ForeignKeys[1].RefTable = RightoftreatmentsTable
	PatientsTable.ForeignKeys[2].RefTable = SystemmembersTable
}
