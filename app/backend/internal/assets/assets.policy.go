package assets

type UploadPolicy struct {
	Category AssetCategory
	MaxSize  int64
	UserID   string
}

type AssetCategory string

const (
	KB = 1 << 10 // 1024
	MB = 1 << 20 // 1.048.576
)
const (
	CategoryProfile  AssetCategory = "avatars"
	CategoryDocument AssetCategory = "document"
	CategoryPost     AssetCategory = "post"
	CategoryTemp     AssetCategory = "temp"
	CategoryVideo    AssetCategory = "video"
)

type AssetPolicy struct {
	Folder      string
	AllowedMime []Mime
	MaxSize     int64
}

var PolicyAsset = map[AssetCategory]AssetPolicy{
	CategoryProfile: {
		Folder:      string(CategoryProfile),
		AllowedMime: []Mime{MimeImageJPEG, MimeImagePNG, MimeImageGIF, MimeImageWebP},
		MaxSize:     1024 * 1024 * 5,
	},
	CategoryDocument: {
		Folder:      string(CategoryDocument),
		AllowedMime: []Mime{MimeDocPDF, MimeDocDocx, MimeDocXlsx, MimeDocPptx, MimeDocDoc},
		MaxSize:     1024 * 1024 * 10,
	},
	CategoryPost: {
		Folder:      string(CategoryPost),
		AllowedMime: []Mime{MimeImageJPEG, MimeImagePNG, MimeImageGIF, MimeImageWebP},
		MaxSize:     1024 * 1024 * 10,
	},
	CategoryTemp: {
		Folder:      string(CategoryTemp),
		AllowedMime: []Mime{MimeImageJPEG, MimeImagePNG, MimeImageGIF, MimeImageWebP},
		MaxSize:     1024 * 1024 * 1,
	},
	CategoryVideo: {
		Folder:      string(CategoryVideo),
		AllowedMime: []Mime{MimeVideoMP4, MimeVideoMOV, MimeVideoAVI, MimeVideoMKV, MimeVideoWebM, MimeVideoFLV},
		MaxSize:     1024 * 1024 * 100,
	},
}
