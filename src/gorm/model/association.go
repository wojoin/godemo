package model

// `Profile` belongs to `User`, `UserID` is the forgien key
type Profile struct {
	ID   int    `gorm:"primary_key"`
	Name string `gorm:"name"`
	//User   User   `gorm:"ForgienKey:UserID"`
	//UserID int    `gorm:"ForeignKey:id"`

	User   User `gorm:"ForgienKey:UserID"`
	UserID int
}

type User struct {
	ID   int    `gorm:"primary_key"`
	Name string `gorm:"name"`
}

// one to many
type Customer struct {
	CustomerID   int       `gorm:"primary_key"`
	CustomerName string    `gorm:"customer_name"`
	Contacts     []Contact `gorm:"ForeignKey:CustId"` //you need to do like this
}

type Contact struct {
	ContactID   int `gorm:"primary_key"`
	CountryCode int
	MobileNo    uint
	CustId      int
}

// Many To Many Relationship
type UserLan struct {
	ID    int `gorm:"primary_key"`
	Uname string
	//Languages []Language `gorm:"many2many:user_languages;ForeignKey:UserId"`
	Languages []Language `gorm:"many2many:user_languages;ForeignKey:id"`
	//Based on this 3rd table user_languages will be created
}

type Language struct {
	ID   int    `gorm:"primary_key"`
	Name string `gorm:"unique_index"`
}

type UserLanguage struct {
	UserLanId  int
	LanguageId int
}

// Has Many Relationship
// https://gorm.io/docs/has_many.html

// Has Many default foreign key

// User has many CreditCards, UserID is the foreign key
type UserModelDefault struct {
	ID          int
	CreditCards []*CreditCardModelDefault
}

type CreditCardModelDefault struct {
	ID                 int
	Number             string
	UserModelDefaultID uint
}

// To use another field as foreign key, you can customize it with a foreignkey tag, e.g:
type UserModel struct {
	//gorm.Model
	ID          int `gorm:"primary_key"`
	Name        string
	CreditCards []*CreditCardModel `gorm:"foreignkey:UserRefer"`
}

type CreditCardModel struct {
	//gorm.Model
	ID        int `gorm:"primary_key"`
	Number    string
	UserRefer uint
}
