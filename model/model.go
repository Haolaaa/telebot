package model

func GetModelStruct(tableName string) interface{} {
	var modelStruct interface{}

	switch tableName {
	// TODO: make it dynamic and not hard-coded
	case "video_changes":
		modelStruct = VideoChanges{}
	}

	return modelStruct
}
