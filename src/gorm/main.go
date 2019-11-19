package main

import (
	"demo/src/gorm/model"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const DBDRIVER = "mysql"
const CONNSTR = "root:zaq!wsx@tcp(10.10.108.105:3306)/auth-aidi-dev2?charset=utf8&parseTime=True&loc=Local"

//type Product struct {
//	//gorm.Model
//	ID    int `gorm:"primary_key"`
//	Code  string
//	Price uint
//}

type User struct {
	ID      uint   `gorm:"primarty_key"`
	Name    string `gorm:"type:varchar(255)"`
	Address string `gorm:"type:varchar(255)"`
}

//type Profile struct {
//	gorm.Model
//	Name string
//	User   User
//	UserID int64
//}

func main() {
	db, err := gorm.Open(DBDRIVER, CONNSTR)
	if err != nil {
		fmt.Errorf("connect mysql error %s", err.Error())
		panic(err.Error())
	}
	defer db.Close()

	fmt.Println(CONNSTR + " Connection Established")

	db.Debug().DropTableIfExists(&User{})
	db.Debug().AutoMigrate(&User{})

	var users []User = []User{
		User{Name: "Ricky", Address: "Sydey"},
		User{Name: "Adam", Address: "Brisbane"},
		User{Name: "Justin", Address: "California"},
	}

	for _, user := range users {
		db.Debug().Create(&user)
	}

	user := &User{Name: "John", Address: "New York"}
	fmt.Println(user)
	user.Address = "California"

	// Updates
	db.Debug().Save(&user)
	fmt.Println(user)

	// update with cloumn name
	//db.Model(&User{}).Update("Name","Jack")
	//db.Model(&User{}).Update(map[string]interface{}{
	//	"Name":"Amy",
	//	"Address":"boston",
	//})

	// delete one record
	db.Debug().Where("address = ?", "Brisbane").Delete(&User{})

	// Queries
	usersView := []User{}
	db.Debug().Where("address=?", "California").Find(&User{})
	db.Debug().Where("address in (?)", []string{"Sydey", "California"}).Find(&User{}).Scan(&usersView)
	fmt.Println(len(usersView), usersView)

	// Associations

	// one to one relationship
	db.Debug().DropTableIfExists(&model.Place{}, &model.Town{})
	db.Debug().AutoMigrate(&model.Place{}, &model.Town{})

	db.Model(&model.Place{}).AddForeignKey("town_id", "towns(id)", "CASCADE", "CASCADE")

	t1 := model.Town{
		Name: "Pune",
	}

	t2 := model.Town{
		Name: "Mumbai",
	}

	t3 := model.Town{
		Name: "Hyderabad",
	}

	p1 := model.Place{
		Name: "Katraj",
		Town: t1,
	}

	p2 := model.Place{
		Name: "Thane",
		Town: t2,
	}
	p3 := model.Place{
		Name: "Secundarabad",
		Town: t3,
	}
	db.Save(&p1)
	db.Save(&p2)
	db.Save(&p3)
	fmt.Println("t1==>", t1, "p1==>", p1)
	fmt.Println("t2==>", t2, "p2==>", p2)
	fmt.Println("t3==>", t3, "p3==>", p3)

	//Delete
	db.Where("name=?", "Hyderabad").Delete(&model.Town{})

	//Update
	db.Model(&model.Place{}).Where("id=?", 1).Update("name", "Shivaji Nagar")

	// Select
	places := model.Place{}
	fmt.Println("before association", places)
	db.Debug().Where("name=?", "Shivaji Nagar").Find(&places)
	fmt.Println("after association", places)

	towns := model.Town{}
	err = db.Model(&places).Association("town").Find(&places.Town).Error
	fmt.Println("after association", towns, places)
	fmt.Println("after association", towns, places, err)

	// Belongs To

	db.Debug().DropTableIfExists(&model.Profile{})
	//db.Debug().DropTableIfExists(&model.User{})
	db.Debug().AutoMigrate(&model.Profile{}, &model.User{})

	db.Debug().Model(&model.Profile{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	//db.Debug().Model(&model.Profile{}).RemoveForeignKey("user_id","users(id)")
	//db.Debug().Model(&model.Profile{}).RemoveForeignKey("user_refer","users(id)")

	//db.Debug().Model(&model.Profile{}).AddForeignKey("user_refer","users(id)", "CASCADE","CASCADE")

	u1 := model.User{
		Name: "Jack",
	}

	profile1 := model.Profile{
		Name: ".bashrc",
		User: u1,
	}
	db.Debug().Save(&profile1)

	u2 := model.User{}
	profile2 := model.User{}
	//db.Model(&u2).Related(&profile2)
	fmt.Println("user====>", u2, "profile====>", profile2)

	// one to many
	db.Debug().DropTableIfExists(&model.Contact{}, &model.Customer{})
	db.Debug().AutoMigrate(&model.Customer{}, &model.Contact{})

	Custs1 := model.Customer{CustomerName: "Martin", Contacts: []model.Contact{
		{CountryCode: 90, MobileNo: 808988},
		{CountryCode: 90, MobileNo: 909699}}}

	db.Debug().Create(&Custs1)

	customers := &model.Customer{}
	contacts := &model.Contact{}

	db.Debug().Where("customer_name=?", "Martin").
		Preload("Contacts").
		Find(&customers) //db.Debug().Where(“customer_name=?”,”John”).Preload(“Contacts”).Find(&customers)

	fmt.Println("Customers", customers)
	fmt.Println("Contacts", contacts)

	// many to many
	db.Debug().DropTableIfExists(&model.UserLanguage{}, &model.Language{}, &model.UserLan{})
	db.Debug().AutoMigrate(&model.UserLan{}, &model.Language{}, &model.UserLanguage{})

	// All foreign keys need to define here
	db.Debug().Model(model.UserLanguage{}).AddForeignKey("user_lan_id", "user_lans(id)", "CASCADE", "CASCADE")
	db.Debug().Model(model.UserLanguage{}).AddForeignKey("language_id", "languages(id)", "CASCADE", "CASCADE")

	// //not define foreign key
	//db.Debug().DropTableIfExists(&model.Language{}, &model.UserLan{})
	//db.Debug().AutoMigrate(&model.UserLan{}, &model.Language{})


	//lang1 := []model.Language{{Name: "English"}, {Name: "French"}}
	lang2 := []model.Language{{Name: "English"}, {Name: "French"}, {Name: "Germany"}}

	//db.Model(&model.Language{}).AddUniqueIndex("idx_name", "name")

	user1 := model.UserLan{Uname: "Martin", Languages: lang2}
	user2 := model.UserLan{Uname: "Ray", Languages: lang2}



	db.Debug().Save(&user1)
	db.Debug().Save(&user2)

	fmt.Println("User1", user1)
	fmt.Println("User2", user2)

	// Has Many Default -- credit card belongs to usermodel
	db.Debug().DropTableIfExists(&model.CreditCardModelDefault{}, &model.UserModelDefault{})
	db.Debug().AutoMigrate(&model.CreditCardModelDefault{}, &model.UserModelDefault{})

	creditCardDefault := []*model.CreditCardModelDefault{
		&model.CreditCardModelDefault{Number:"123"},
		&model.CreditCardModelDefault{Number:"456"},
	}
	userModelDefault1 := model.UserModelDefault{
		CreditCards:creditCardDefault,
	}

	db.Debug().Save(&userModelDefault1)

	fmt.Println("userModelDefault1", userModelDefault1)

	// Has Many -- credit card belongs to usermodel
	db.Debug().DropTableIfExists(&model.CreditCardModel{}, &model.UserModel{})
	db.Debug().AutoMigrate(&model.CreditCardModel{}, &model.UserModel{})

	creditCard := []*model.CreditCardModel{
		&model.CreditCardModel{Number:"123", UserRefer:1},
		&model.CreditCardModel{Number:"456", UserRefer:2},
	}
	userModel1 := model.UserModel{
		Name:"jim",
		CreditCards:creditCard,
	}

	db.Debug().Save(&userModel1)

	fmt.Println("userModel1", userModel1)


}
