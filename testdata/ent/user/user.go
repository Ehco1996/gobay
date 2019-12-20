// Code generated by entc, DO NOT EDIT.

package user

import (
	"github.com/shanbay/gobay/testdata/ent/schema"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNickname holds the string denoting the nickname vertex property in the database.
	FieldNickname = "nickname"
	// FieldUsername holds the string denoting the username vertex property in the database.
	FieldUsername = "username"

	// Table holds the table name of the user in the database.
	Table = "users"
)

// Columns holds all SQL columns are user fields.
var Columns = []string{
	FieldID,
	FieldNickname,
	FieldUsername,
}

var (
	fields = schema.User{}.Fields()

	// descNickname is the schema descriptor for nickname field.
	descNickname = fields[0].Descriptor()
	// DefaultNickname holds the default value on creation for the nickname field.
	DefaultNickname = descNickname.Default.(string)
)
