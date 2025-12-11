package base

// DeletedFlag represents whether an entity is deleted.
type DeletedFlag bool

// bool returns the boolean value of the DeletedFlag.
func (d DeletedFlag) getBool() bool {
	return bool(d)
}

// DeletedFlagFromBool creates a DeletedFlag from a boolean value.
func DeletedFlagFromBool(b bool) DeletedFlag {
	return DeletedFlag(b)
}

// NewDeletedFlag creates a new DeletedFlag set to false.
func NewDeletedFlag() DeletedFlag {
	return DeletedFlag(false)
}
