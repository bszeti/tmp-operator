package controller

import (
	"github.com/redhat/azure-agent-operator/pkg/controller/azureagentpool"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, azureagentpool.Add)
}
