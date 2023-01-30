package mongo

type UriParams struct {
	Host string
	Port string
	User string
	Pass string
	Db   string
}

func CalcMongoUri(params UriParams) string {
	uri := "mongodb://"
	if params.User != "" && params.Pass != "" {
		uri += params.User + ":" + params.Pass + "@"
	}
	uri += params.Host + ":" + params.Port + "/" + params.Db
	return uri
}
