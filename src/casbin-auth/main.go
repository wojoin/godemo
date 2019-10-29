package main

import (
	"fmt"
	"github.com/casbin/casbin"
	"github.com/casbin/casbin/rbac"
	"github.com/casbin/casbin/rbac/default-role-manager"
	"github.com/casbin/mongodb-adapter"
	"log"
	"sync"
)


type CasbinRbac struct {
	RM       rbac.RoleManager
	Enforcer *casbin.SyncedEnforcer
}

var cr *CasbinRbac
var lock = &sync.Mutex{}


const mongoUrl = "192.168.232.161:27017"
const model = "src/casbin-auth/conf/rbac_model.conf"

func Casbin() *casbin.SyncedEnforcer {
	a := mongodbadapter.NewAdapter(mongoUrl)
	e,_ := casbin.NewSyncedEnforcer("src/casbin-auth/conf/rbac_model.conf", a)
	//e.LoadPolicy()
	return e
}

// init casbin.SyncedEnforcer
func NewEnforcerUseMongo(mongodbURL string) *casbin.SyncedEnforcer {
	adapter := mongodbadapter.NewAdapter(mongodbURL)
	casbinModelPath := model
	e,_ := casbin.NewSyncedEnforcer(casbinModelPath, adapter)
	return e
}

func InitCasbin(e *casbin.SyncedEnforcer) *CasbinRbac {
	lock.Lock()
	defer lock.Unlock()
	if cr == nil {
		cr = &CasbinRbac{}
		cr.Enforcer = e
		cr.RM = defaultrolemanager.NewRoleManager(3)
		cr.Enforcer.SetRoleManager(cr.RM)
		err := cr.Enforcer.LoadPolicy()
		if err != nil {
			log.Println("init casbin load policy error", err)
		}
	}
	return cr
}

func AddRoleForUser(user,role string) bool {
	if cr.Enforcer.HasNamedGroupingPolicy("g", user, role) {
		return true
	}

	if ok ,_ := cr.Enforcer.AddNamedGroupingPolicy("g", user, role); !ok{
		return false
	}

	return true
}

func main()  {
	//mongoUrl := "mongodb://root:123@192.168.232.161:27017/casbin-auth"

	cr := InitCasbin(NewEnforcerUseMongo(mongoUrl))
	//e := cr.Enforcer

	//e := Casbin()

	cr.Enforcer.AddPolicy("alice","data1","read")
	AddRoleForUser("alice", "data2_admin")


	if err := cr.Enforcer.SavePolicy(); err != nil{
		fmt.Println("fail")
	}


}
