package dao

import (
	"lawencon.com/bootest/model"
)

type EmployeeProfileDao interface {
	CreateEmployee(data *model.EmployeeProfiles) error
}
