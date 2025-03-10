package wrike

import (
	types "github.com/golyakov/wrike.go/types"
)

// QueryWorkflows fetches a Workflows
// For details refer to https://developers.wrike.com/documentation/api/methods/query-workflows
func (api *API) QueryWorkflows() (*types.Workflows, error) {
	path := "/workflows"

	data, err := api.get(path, nil)
	if err != nil {
		return nil, err
	}

	return types.NewWorkflowsFromJSON(data)
}
