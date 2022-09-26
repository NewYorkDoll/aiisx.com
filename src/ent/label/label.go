// DO NOT EDIT, CODE GENERATED BY entc. yiziluoying

package label

const (
	// Label holds the string label denoting the label type in the database.
	Label = "label"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// Table holds the table name of the label in the database.
	Table = "labels"
)

// Columns holds all SQL columns for label fields.
var Columns = []string{
	FieldID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}