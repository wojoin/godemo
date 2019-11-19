package wrapper

import (
	"github.com/casbin/casbin"
	"github.com/casbin/casbin/rbac"
	"github.com/casbin/casbin/rbac/default-role-manager"
	"github.com/casbin/mongodb-adapter"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
	"user_auth/conf"
	"user_auth/utils"
	// "user_auth/utils"
)

type CasbinRbac struct {
	RM       rbac.RoleManager
	Enforcer *casbin.SyncedEnforcer
}

var cr *CasbinRbac
var lock = &sync.RWMutex{}

// init casbin.SyncedEnforcer
func NewEnforcerUseMongo(mongodbURL string) *casbin.SyncedEnforcer {
	// w := utils.NewWatcher(conf.AppConfig.ZKHOST.HOST)
	adapter := mongodbadapter.NewAdapter(mongodbURL) // Your MongoDB URL.
	casbinModelPath := conf.AppConfig.CasbinModelPath
	e, _ := casbin.NewSyncedEnforcer(casbinModelPath, adapter)
	// e.SetWatcher(w)
	// w.SetUpdateCallback(func(rev string) {
	// 	log.WithFields(log.Fields{"module": "zkwatcher"}).Info(
	// 		"zk callback %s\n", rev)
	// 	e.LoadPolicy()
	// })
	return e
}

// 使用自定义的函数初始化casbin
func InitCasbin(e *casbin.SyncedEnforcer) *CasbinRbac {
	lock.Lock()
	defer lock.Unlock()
	if cr == nil {
		cr = &CasbinRbac{}
		cr.Enforcer = e
		// 组织结构定义为3层，继承关系最多3层，超过3层需要修改这个参数
		cr.RM = defaultrolemanager.NewRoleManager(3)
		cr.Enforcer.SetRoleManager(cr.RM)
		// 添加自定义函数，model中的matcher中可以使用了
		//cr.Enforcer.AddFunction("customerfunc", roleInheritFunc)
		//e.AddFunction("keyMatch", KeyMatchFunc)
		err := cr.Enforcer.LoadPolicy()
		if err != nil {
			log.Println("init casbin load policy error", err)
		}
	}
	return cr
}

//func KeyMatch(key1 string, key2 string) bool {
//	i := strings.Index(key2, "*")
//	if i == -1 {
//		return key1 == key2
//	}
//
//	if len(key1) > i {
//		return key1[:i] == key2[:i]
//	}
//	return key1 == key2[:i]
//}
//
//func KeyMatchFunc(args ...interface{}) (interface{}, error) {
//	name1 := args[0].(string)
//	name2 := args[1].(string)
//
//	return (bool)(KeyMatch(name1, name2)), nil
//}

// 自定义函数，判断role1是否继承了role2的权限
func roleInherit(role1 string, role2 string, rm rbac.RoleManager) bool {
	if role1 == role2 {
		return true
	}
	// 判断role1是否继承了role2，最多做3层判断。
	ok, err := rm.HasLink(role1, role2)
	if err != nil {
		log.WithFields(log.Fields{"module": "casbin"}).Error(err.Error())
	}
	return ok
}

// 封装一下自定义函数
func roleInheritFunc(args ...interface{}) (interface{}, error) {
	user := args[0].(string)
	role := args[1].(string)
	return (bool)(roleInherit(user, role, cr.RM)), nil
}

//add role for user,you can also treat group as a user
// Returns true if the user already has the role (aka not affected).
func AddRoleForUser(role, user string) bool {
	t := time.Now()
	if ok, _ := cr.Enforcer.HasRoleForUser(user, role); !ok {
		log.Info("HasRoleForUser--UserDefine-medium latency:", time.Since(t))
		return true
	}

	if ok, _ := cr.Enforcer.AddRoleForUser(user, role); !ok {
		log.Info("AddRoleForUser--UserDefine-medium latency:", time.Since(t))
		return true
	}
	//if cr.Enforcer.HasNamedGroupingPolicy("g", user, role) {
	//	log.Info("HasRoleForUser method latency:", time.Since(t))
	//	return true
	//}
	//
	//if cr.Enforcer.AddNamedGroupingPolicy("g", user, role) {
	//	log.Info("AddRoleForUser method latency:", time.Since(t))
	//	return true
	//}

	log.Info("AddRoleForUser(sum) method latency:", time.Since(t))

	return false
}

// concurrency safe
func GetRolesForUser2(user string) []string {
	startime := time.Now()
	result := make([]string, 0)
	roles := cr.Enforcer.GetFilteredGroupingPolicy(0, user)
	//log.Infof("user [%s] [%d] roles [%v]: ", user, len(roles), roles)
	for _, r := range roles {
		result = append(result, r[1])
		rolesInh := cr.Enforcer.GetFilteredGroupingPolicy(0, r[1])
		for _, inhRole := range rolesInh {
			result = append(result, inhRole[1])
		}

	}

	result = utils.RemoveDuplicate(result)
	log.Info("GetRolesForUser latency ", time.Since(startime), result)
	return result
}

// 获取用户的所有角色
//this function retrieves indirect roles besides direct roles.
func GetRolesForUser(user string) []string {
	lock.RLock()
	defer lock.RUnlock()
	retval, _ := cr.Enforcer.GetImplicitRolesForUser(user)
	return retval
}

// delete role for user
// Returns false if the user does not have the role (aka not affected).
func DeleteRoleForUser(role, user string) bool {
	ok, _ := cr.Enforcer.DeleteRoleForUser(user, role)
	return ok
}

// add  permission by who,what,how
// If the rule already exists, the function returns true and the rule will not be added.
func AddPermissionForRole(subject, domain, object, action string) bool {
	if ok, _ := cr.Enforcer.AddPolicy(subject, domain, object, action); ok {
		return true
	} else {
		if cr.Enforcer.HasNamedPolicy("p", subject, domain, object, action) {
			return true
		}
	}
	return false
}

//// delete permisson by who,what,how
//func DeletePermissionForRole(subject, object, action string) bool {
//	ok, _ := cr.Enforcer.DeletePermissionForUser(subject, object, action)
//	return ok
//}
//
//// verify by who,what,how
//func Verify(subject, object, action string) (bool, error) {
//	return cr.Enforcer.EnforceSafe(subject, object, action)
//}
//
//// 获取用户/角色的所有权限
//// this function retrieves permissions for inherited roles.
//func GetPermissionForUser(user string) [][]string {
//	return cr.Enforcer.GetImplicitPermissionsForUser(user)
//}
//
//// Get all resources that a user or role can access
//func GetResourcesForUser(user string) []string {
//	var resources []string
//	userpermissions := cr.Enforcer.GetImplicitPermissionsForUser(user)
//	for _, userpermisson := range userpermissions {
//		if !Contains(resources, userpermisson[1]) {
//			resources = append(resources, userpermisson[1])
//		}
//	}
//	return resources
//}
//
//func GetResourceByID(userid string) []string {
//	return cr.Enforcer.GetImplicitRolesForUser(userid)
//}
//
//func GetResourceByRole(role string) [][]string {
//	return cr.Enforcer.GetPermissionsForUser(role)
//}
//
//// get users for role,you can define prefix of user,such as "user_" "platform_" "auto_"
//func GetUsersForRole(role string, prefixstring string) []string {
//	roles := cr.Enforcer.GetUsersForRole(role)
//	users := make([]string, 0)
//	for _, role := range roles {
//		if strings.HasPrefix(role, prefixstring) {
//			users = append(users, role)
//		}
//	}
//	return users
//}
//
//func GetUsersForGroupRole(role string) []string {
//	roles := cr.Enforcer.GetUsersForRole(role)
//	users := make([]string, 0, len(roles))
//	for _, role := range roles {
//		users = append(users, role)
//
//	}
//	return users
//}
//
//// 获取可以操作资源的所有角色/用户
//func GetAllRolesForDataPermission(data ...string) map[string]struct{} {
//	permissions := cr.Enforcer.GetFilteredNamedPolicy("p", 1, data...)
//	//log.Println("permissions in GetAllRolesForDataPermission", permissions)
//	roles := make(map[string]struct{}, 0)
//	for _, permission := range permissions {
//		roles[permission[0]] = struct{}{}
//		users := cr.Enforcer.GetUsersForRole(permission[0])
//		for _, user := range users {
//			roles[user] = struct{}{}
//		}
//	}
//	//log.Println("GetAllRolesForDataPermission", roles)
//	return roles
//}
//
//// 给一个资源加入到资源组里
//// If the rule already exists, the function returns true and the rule will not be added.
//func AddGroupForObject(data, group string) bool {
//	if ok, _ := cr.Enforcer.AddNamedGroupingPolicySafe("g2", []string{data, group}); ok {
//		return true
//	} else {
//		if cr.Enforcer.HasNamedGroupingPolicy("g2", []string{data, group}) {
//			log.Println("data already added to this group")
//			return true
//		}
//	}
//	return false
//}
//
////todo
////get all resources that a group has
//
////delete resources including related resource roles and policy
//func DeleteResource(data string) bool {
//	return cr.Enforcer.RemoveFilteredNamedPolicy("p", 1, data)
//}
//
//// 获取资源加入的分组
//func GetGroupsForObject(data string) []string {
//	return cr.Enforcer.GetImplicitRolesForUser(data)
//}
