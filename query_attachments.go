package wrike

import (
	"fmt"

	parameters "github.com/AkihikoITOH/wrike.go/parameters"
	types "github.com/AkihikoITOH/wrike.go/types"
	"github.com/google/go-querystring/query"
)

// QueryTaskAttachments fetches a list of attachments.
// For details refer to https://developers.wrike.com/api/v4/attachments/#get-attachments
func (api *API) QueryTaskAttachments(taskId types.TaskID, params parameters.QueryTaskAttachments) (*types.Attachments, error) {
	path := fmt.Sprintf("/tasks/%s/attachments", taskId)

	qp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	data, err := api.get(path, &qp)
	if err != nil {
		return nil, err
	}

	return types.NewAttachmentsFromJSON(data)
}
