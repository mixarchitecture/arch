package example

type messages struct {
	Get_Failed        string
	Validation_Failed string
	Create_Failed     string
	Update_Failed     string
	List_Failed       string
	Count_Failed      string
	Unmarshal_Failed  string
	Not_Found         string
	Content_Too_long  string
	Field_Too_long    string
	Content_Too_Short string
	Field_Too_Short   string
}

var I18nMessages = messages{
	Get_Failed:        "error_example_get",
	Validation_Failed: "error_example_validation",
	Create_Failed:     "error_example_create",
	Update_Failed:     "error_example_update",
	List_Failed:       "error_example_list",
	Count_Failed:      "error_example_count",
	Unmarshal_Failed:  "error_example_unmarshal",
	Not_Found:         "error_example_not_found",
	Content_Too_long:  "error_example_content_too_long",
	Field_Too_long:    "error_example_field_too_long",
	Content_Too_Short: "error_example_content_too_short",
	Field_Too_Short:   "error_example_field_too_short",
}
