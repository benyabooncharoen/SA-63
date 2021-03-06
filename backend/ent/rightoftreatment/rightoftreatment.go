// Code generated by entc, DO NOT EDIT.

package rightoftreatment

const (
	// Label holds the string label denoting the rightoftreatment type in the database.
	Label = "rightoftreatment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRightoftreatmentName holds the string denoting the rightoftreatmentname field in the database.
	FieldRightoftreatmentName = "rightoftreatment_name"

	// EdgePatients holds the string denoting the patients edge name in mutations.
	EdgePatients = "patients"

	// Table holds the table name of the rightoftreatment in the database.
	Table = "rightoftreatments"
	// PatientsTable is the table the holds the patients relation/edge.
	PatientsTable = "patients"
	// PatientsInverseTable is the table name for the Patient entity.
	// It exists in this package in order to avoid circular dependency with the "patient" package.
	PatientsInverseTable = "patients"
	// PatientsColumn is the table column denoting the patients relation/edge.
	PatientsColumn = "Rightoftreatment"
)

// Columns holds all SQL columns for rightoftreatment fields.
var Columns = []string{
	FieldID,
	FieldRightoftreatmentName,
}
