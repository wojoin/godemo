package conf

var AppConfig = struct {
	MONGO           struct {
		URL      string `default:"127.0.0.1"`
		USERNAME string `default:""`
		PASSWORD string `default:""`
		DBNAME   string `default:"test"`
	}
}{}
