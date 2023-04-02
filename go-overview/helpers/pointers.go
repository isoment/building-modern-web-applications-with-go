package helpers

func ChangeUsingPointer(s *string) {
	newValue := "Red"
	*s = newValue
}
