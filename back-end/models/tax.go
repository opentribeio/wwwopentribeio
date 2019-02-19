package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	u "go-contacts/utils"
)

type Tax struct {
	gorm.Model
	Name   string `json:"name"`
	Value  uint   `json:"value"` //procent value
	Avg    string `json:"avg"`
	UserId uint   `json:"user_id"` //The user that this contact belongs to
}

/*
 This struct function validate the required parameters sent through the http request body

returns message and true if the requirement is met
*/
func (tax *Tax) Validate() (map[string]interface{}, bool) {

	if tax.Name == "" {
		return u.Message(false, "Contact name should be on the payload"), false
	}

	if tax.Value == 0 {
		return u.Message(false, "tax number should be on the payload"), false
	}
	if tax.Value > 100 {
		return u.Message(false, "tax value should be percent between 0 and 100"), false
	}

	if tax.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (tax *Tax) Create() map[string]interface{} {

	if resp, ok := tax.Validate(); !ok {
		return resp
	}

	GetDB().Create(tax)

	resp := u.Message(true, "success")
	resp["tax"] = tax
	return resp
}

func GetTax(name string) *Tax {

	tax := &Tax{}
	err := GetDB().Table("taxes").Where("Name = ?", name).First(tax).Error
	if err != nil {
		return nil
	}
	return tax
}

func (tax *Tax) Update(name string, value uint) map[string]interface{} {

	GetDB().Model(&tax).Where("Name = ?", tax.Name).Update("Value", tax.Value)
	resp := u.Message(true, "success")
	resp["tax"] = tax
	return resp
}

func GetTaxes(user uint) []*Tax {

	taxes := make([]*Tax, 0)
	err := GetDB().Table("taxes").Where("user_id = ?", user).Find(&taxes).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return taxes
}
