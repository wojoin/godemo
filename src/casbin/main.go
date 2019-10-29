package main

import (
	"github.com/casbin/casbin"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	e := casbin.NewEnforcer("src/casbin/conf/rbac_model.conf", "src/casbin/conf/rbac_policy.csv")

	//e := casbin.NewEnforcer("./conf/rbac_model_with_resource_roles.conf",
	//	"./conf/rbac_policy_with_resource_role.csv")

	u := "alice"
	log.Printf("user [%s] has role %s",u, e.GetRolesForUser(u))
	log.Printf("user [%s] has permission %s",u, e.GetPermissionsForUser(u))
	log.Printf("user [%s] has permission %s",u, e.GetImplicitPermissionsForUser(u))

	log.Println("-----------add user jack, role: data2_admin")
	if ok := e.AddRoleForUser("jack","data2_admin"); !ok {
		return
	}
	u = "jack"
	log.Printf("user [%s] has role %s",u, e.GetRolesForUser(u))
	//log.Printf("user [%s] has permission %s",u, e.GetPermissionsForUser(u))
	log.Printf("user [%s] has permission %s",u, e.GetImplicitPermissionsForUser(u))

	if ok := e.DeleteUser(u); !ok {
		return
	}
	log.Println("delete user ", u, "success.")

	log.Printf("user [%s] has role %s",u, e.GetRolesForUser(u))
	//log.Printf("user [%s] has permission %s",u, e.GetPermissionsForUser(u))
	log.Printf("user [%s] has permission %s",u, e.GetImplicitPermissionsForUser(u))



	log.Println("---delete role:data2_admin before, alice can access data2 via role:data2_admin")
	u = "alice"
	log.Printf("user [%s] has role %s",u, e.GetRolesForUser(u))
	log.Printf("user [%s] has permission %s",u, e.GetPermissionsForUser(u))
	log.Printf("user [%s] has permission %s",u, e.GetImplicitPermissionsForUser(u))

	log.Println("---delete role:data2_admin after, alice can not access data2 via role:data2_admin")
	e.DeleteRole("data2_admin") // delete role , just only role, not related with user permission

	// alice has role:data1_admin only
	log.Printf("user [%s] has role %s",u, e.GetRolesForUser(u))
	log.Printf("user [%s] has permission %s",u, e.GetPermissionsForUser(u))
	log.Printf("user [%s] has permission %s",u, e.GetImplicitPermissionsForUser(u))

	// add user for resource role
	log.Println("-----------add role for resource")
	u = "jack"
	role := "developer"
	// g, alice, developer
	// so jack can operation resource by role:developer
	if ok := e.AddRoleForUser(u, role); !ok {
		return
	}
	log.Printf("(resource)user [%s] has role %s",u, e.GetRolesForUser(u))
	log.Printf("(resource)user [%s] has role %s",u, e.GetImplicitPermissionsForUser(u))

	// get all role of user:alice
	log.Printf("user [%s] has role %s","alice", e.GetImplicitRolesForUser("alice"))

	// add role for resource
	log.Println("-------------add role for resource")
	res := "data1"
	role = "owner"
	// g, data1, owner
	// so the user adding role:owner can operation resource data1 by role:owner
	if ok := e.AddRoleForUser(res, role); !ok {
		return
	}

	log.Printf("(resource)user [%s] has role %s",res, e.GetImplicitRolesForUser(res))
	// role:developer role of resource can read permission on data1 and data2
	// new user adding role:developer can also read permission on data1 and data2
	// eg. jack

	u = "join"
	role = "developer"
	if ok := e.AddRoleForUser(u, role); !ok {
		return
	}
	log.Printf("(resource)user [%s] has role %s",u, e.GetRolesForUser(u))
	log.Printf("(resource)user [%s] has role %s",u, e.GetImplicitPermissionsForUser(u))

	users := e.GetRolesForUser("data1")
	//log.Printf("user [%s] can access to resource",users, "data1")
	log.Printf("user [%s] can access to resource %s",users, "data1")

}
