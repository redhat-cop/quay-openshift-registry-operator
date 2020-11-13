package controller

import (
	"github.com/redhat-cop/quay-openshift-registry-operator/controller/quayintegration"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, quayintegration.Add)
}
