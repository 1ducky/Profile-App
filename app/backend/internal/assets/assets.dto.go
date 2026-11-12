package assets

type RecordPayload struct {
	ID       string
	FileKey  string
	Status   AssetStatus
	Filename string
	Mime     Mime
	AuthorID string
	Category AssetCategory
}

type UpdatePayload struct {
	Id        string
	Size      int64
	Status    AssetStatus
	OldStatus AssetStatus
	AuthorID  string
}

type DeletePayload struct {
	IDs []string
}

type DeleteReaport struct {
	id  string
	err error
}

type StoreResult struct {
	FileKey string
	Size    int64
}
