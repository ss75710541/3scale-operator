package controller

import (
	"github.com/ss75710541/3scale-operator/pkg/controller/tenant"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, tenant.Add)
}
