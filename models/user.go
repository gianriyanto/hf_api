package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int64    `orm:"pk;auto;column(id)" json:"id"`
	Username string   `orm:"column(username)" json:"username"`
	Password string   `orm:"column(password)" json:"password"`
	Profile  *Profile `orm:"rel(one)" json:"profile"` // OneToOne relation
}

type Profile struct {
	Id      int64  `orm:"pk;auto;column(id)" json:"id"`
	Gender  string `orm:"column(gender)" json:"gender"`
	Age     int    `orm:"column(age)" json:"age"`
	Address string `orm:"column(address)" json:"address"`
	Email   string `orm:"column(email)" json:"email"`
	User    *User  `orm:"reverse(one)" json:"-"` // Reverse relationship (optional)
}

func AddUser(u User) int64 {
	o := orm.NewOrm()

	_, err := o.Insert(u.Profile)
	if err != nil {
		return 0
	}

	id, err := o.Insert(&u)
	if err == nil {
		return id
	}

	return u.Id
}

func GetUser(uid int64) (u *User, err error) {
	o := orm.NewOrm()
	var user User
	err = o.QueryTable("user").Filter("id", uid).RelatedSel().One(&user)
	if err == nil {
		u = &user
		return
	}

	return
}

func GetAllUsers() (user []User) {
	o := orm.NewOrm()
	_, err := o.QueryTable("user").RelatedSel().All(&user)
	if err == nil {
		return
	}

	user = []User{}

	return
}

func UpdateUser(uid int64, uu *User) (u *User, err error) {
	o := orm.NewOrm()

	u, err = GetUser(uid)
	if err != nil {
		err = errors.New("User Does Not Exist")
		return
	}

	if uu.Username != "" {
		u.Username = uu.Username
	}
	if uu.Password != "" {
		u.Password = uu.Password
	}
	if uu.Profile.Age != 0 {
		u.Profile.Age = uu.Profile.Age
	}
	if uu.Profile.Address != "" {
		u.Profile.Address = uu.Profile.Address
	}
	if uu.Profile.Gender != "" {
		u.Profile.Gender = uu.Profile.Gender
	}
	if uu.Profile.Email != "" {
		u.Profile.Email = uu.Profile.Email
	}

	if _, err = o.Update(u.Profile); err == nil {
		return
	}
	if _, err = o.Update(u); err == nil {
		return
	}

	return nil, errors.New("User Does Not Exist")
}

func DeleteUser(uid int64) {
	o := orm.NewOrm()
	if num, err := o.Delete(&User{Id: uid}); err == nil {
		fmt.Println(num)
	}
}
