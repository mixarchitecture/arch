package entity

type MySQLExample struct {
	UUID    string `db:"uuid"`
	Field   string `db:"field"`
	Content string `db:"content"`
}

type fields struct {
	UUID    string
	Field   string
	Content string
	Table   string
}

var Fields = fields{
	UUID:    "uuid",
	Field:   "field",
	Content: "content",
	Table:   "example",
}
