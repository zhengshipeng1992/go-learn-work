// Code generated by entc, DO NOT EDIT.

package house

const (
	// Label holds the string label denoting the house type in the database.
	Label = "house"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldCommunity holds the string denoting the community field in the database.
	FieldCommunity = "community"
	// FieldImage holds the string denoting the image field in the database.
	FieldImage = "image"
	// FieldToiletCount holds the string denoting the toilet_count field in the database.
	FieldToiletCount = "toilet_count"
	// FieldKitchenCount holds the string denoting the kitchen_count field in the database.
	FieldKitchenCount = "kitchen_count"
	// FieldFloorCount holds the string denoting the floor_count field in the database.
	FieldFloorCount = "floor_count"
	// FieldHallCount holds the string denoting the hall_count field in the database.
	FieldHallCount = "hall_count"
	// FieldRoomCount holds the string denoting the room_count field in the database.
	FieldRoomCount = "room_count"
	// Table holds the table name of the house in the database.
	Table = "houses"
)

// Columns holds all SQL columns for house fields.
var Columns = []string{
	FieldID,
	FieldPrice,
	FieldTitle,
	FieldCommunity,
	FieldImage,
	FieldToiletCount,
	FieldKitchenCount,
	FieldFloorCount,
	FieldHallCount,
	FieldRoomCount,
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

var (
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// DefaultImage holds the default value on creation for the "image" field.
	DefaultImage string
	// DefaultToiletCount holds the default value on creation for the "toilet_count" field.
	DefaultToiletCount int32
	// DefaultKitchenCount holds the default value on creation for the "kitchen_count" field.
	DefaultKitchenCount int32
	// DefaultFloorCount holds the default value on creation for the "floor_count" field.
	DefaultFloorCount int32
	// DefaultHallCount holds the default value on creation for the "hall_count" field.
	DefaultHallCount int32
	// DefaultRoomCount holds the default value on creation for the "room_count" field.
	DefaultRoomCount int32
)
