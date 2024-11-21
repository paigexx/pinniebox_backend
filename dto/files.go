package dto

type FileUploadRequest struct {
	File string `json:"file"`
	TelegramID string `json:"tg_id"`
}

//responses from Pinata Files API
type FileUploadResponse struct {  
	Data struct {
	ID            string            `json:"id"`
	Name          string            `json:"name"`
	CID           string            `json:"cid"`
	CreatedAt     string            `json:"created_at"`
	Size          int               `json:"size"`
	NumberOfFiles int               `json:"number_of_files"`
	MimeType      string            `json:"mime_type"`
	UserID        string            `json:"user_id"`
	KeyValues     map[string]string `json:"keyvalues"`
	IsDuplicate   *bool             `json:"is_duplicate"` 
	} `json:"data"`
}

type File struct {
	ID 		  		string            `json:"id"`
	Name      		string            `json:"name"`
	CID       		string            `json:"cid"`
	Size      		int               `json:"size"`
	NumberOfFiles 	int           	  `json:"number_of_files"`
	MimeType  		string            `json:"mime_type"`
	GroupID  		string            `json:"group_id"`
	KeyValues 		map[string]string `json:"keyvalues"`
	CreatedAt 		string            `json:"created_at"`
}

type ListFilesResponse struct {
    Data struct {
        Files          []File `json:"files"`
        NextPageToken  string `json:"next_page_token"`
    } `json:"data"`
}

type UpdateFileResponse struct {
	Data File `json:"data"`
}

type SignedURLResponse struct {
	Data string `json:"data"`
}