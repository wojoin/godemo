package conf

var AppConfig = struct {
	APPName         string `default:"userauth"`
	HTTPPORT        string `default:"9000"`
	LogLevel        string `default:"fatal"` //debug/info/warn/error/fatal
	CasbinModelPath string `default:""`
	MONGO           struct {
		URL      string `default:"10.10.108.145:27017"`
		USERNAME string `default:"root"`
		PASSWORD string `default:"zaq!wsx"`
		DBNAME   string `default:"userauth-dev"`
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
