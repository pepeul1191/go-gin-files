// config/models.go

package config

type SignResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type JWTAccess struct {
	Token string `json:"token,omitempty"`
}

type FileData struct {
	Filename         string `json:"filename"`
	Path             string `json:"path"`
	OriginalFilename string `json:"original_filename"`
	Size             int64  `json:"size"`
	MimeType         string `json:"mime_type"`
}
