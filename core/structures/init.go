package Structures

func initiator() {
	models := GetModels()

	for _, model := range models {
		NewController(model)
	}
}
