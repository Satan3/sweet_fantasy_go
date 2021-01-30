package validation

type Error struct {
	FailedField string
	Tag         string
	Value       string
}
