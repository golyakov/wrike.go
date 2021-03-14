package types

import "encoding/json"

type Attachments struct {
	Kind string `json:"kind"`
	Data []Task `json:"data"`
}

type AttachmentID string

type Attachment struct {
	ID                  AttachmentID `json:"id"`
	AuthorId            string       `json:"authorId"`
	Name                string       `json:"name"`
	CreatedDate         string       `json:"createdDate"`
	Version             int          `json:"version"`
	Type                string       `json:"type"`
	ContentType         string       `json:"contentType"`
	Size                int          `json:"size"`
	TaskId              string       `json:"taskId"`
	FolderId            string       `json:"folderId"`
	CommentId           string       `json:"commentId,omitempty"`
	CurrentAttachmentId string       `json:"currentAttachmentId,omitempty"`
	PreviewUrl          string       `json:"previewUrl,omitempty"`
	Url                 string       `json:"url,omitempty"`
	PlaylistUrl         string       `json:"playlistUrl,omitempty"`
	ReviewIds           []string     `json:"reviewIds,omitempty"`
	Width               int          `json:"width,omitempty"`
	Height              int          `json:"height,omitempty"`
}

// NewTasksFromJSON parses the given JSON (as byte sequence) and returns a new Tasks.
func NewAttachmentsFromJSON(data []byte) (*Attachments, error) {
	var attachments Attachments
	err := json.Unmarshal(data, &attachments)
	return &attachments, err
}
