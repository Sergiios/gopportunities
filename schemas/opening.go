package schemas

import (
	"fmt"

	"github.com/Sergiios/gopportunities/utils"
	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int    `json:"salary"`
}

func (o *Opening) CreateValidate() error {
	if o.ID != 0 {
		return fmt.Errorf("field id cannot be defined at creation")
	} else if o.Role == "" {
		return utils.ParamRequired("role", "string")
	} else if o.Company == "" {
		return utils.ParamRequired("company", "string")
	} else if o.Location == "" {
		return utils.ParamRequired("location", "string")
	} else if o.Remote == nil {
		return utils.ParamRequired("remote", "bool")
	} else if o.Salary <= 0 {
		return utils.ParamRequired("salary", "int")
	}
	return nil
}
