package conf

var AppConfig = struct {
	APPName         string `default:"userauth"`
	HTTPPORT        string `default:"9000"`
	LogLevel        string `default:"fatal"`
	CasbinModelPath string `default:""`
	MONGO           struct {
		URL      string `default:"127.0.0.1:27017"`
		USERNAME string `default:""`
		PASSWORD string `default:""`
		DBNAME   string `default:"test"`
	}
	ZKHOST struct {
		HOST string `default:"127.0.0.1:2181"`
	}
	LDAP struct {
		HOST string `default:"127.0.0.1:1236"`
	}
	REDIS struct {
		HOST string `default:"127.0.0.1:6379"`
	}
	KEYPATH struct {
		PRIKEYPATH string `default:"./keys/private.pem"`
		PUBKEYPATH string `default:"./keys/public.pub"`
	}
}{}
