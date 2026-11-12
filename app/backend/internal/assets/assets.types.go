package assets

type Mime string
type Ext string

type TypeResolver struct {
	Mime Mime
	Ext  Ext
}

const (
	//  Image Mime Allowed
	MimeImageJPEG Mime = "image/jpeg"
	MimeImagePNG  Mime = "image/png"
	MimeImageGIF  Mime = "image/gif"
	MimeImageWebP Mime = "image/webp"

	// Video
	MimeVideoMP4  Mime = "video/mp4"
	MimeVideoMOV  Mime = "video/quicktime"
	MimeVideoAVI  Mime = "video/x-msvideo"
	MimeVideoMKV  Mime = "video/x-matroska"
	MimeVideoWebM Mime = "video/webm"
	MimeVideoFLV  Mime = "video/x-flv"

	// Document
	MimeDocPDF  Mime = "application/pdf"
	MimeDocDocx Mime = "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	MimeDocXlsx Mime = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	MimeDocPptx Mime = "application/vnd.openxmlformats-officedocument.presentationml.presentation"
	MimeDocDoc  Mime = "application/msword"
)

const (
	//  Image Mime Allowed
	ExtJpg  Ext = ".jpg"
	ExtPng  Ext = ".png"
	ExtGif  Ext = ".gif"
	ExtWebp Ext = ".webp"

	// Video Ext
	ExtMp4  Ext = ".mp4"
	ExtMov  Ext = ".mov"
	ExtAvi  Ext = ".avi"
	ExtMkv  Ext = ".mkv"
	ExtWebm Ext = ".webm"
	ExtFlv  Ext = ".flv"

	// Document Ext
	ExtPdf  Ext = ".pdf"
	ExtDocx Ext = ".docx"
	ExtXlsx Ext = ".xlsx"
	ExtPptx Ext = ".pptx"
	ExtDoc  Ext = ".doc"
)

var AllowedImageType = map[Mime]TypeResolver{
	MimeImageJPEG: {Mime: MimeImageJPEG, Ext: ExtJpg},
	MimeImagePNG:  {Mime: MimeImagePNG, Ext: ExtPng},
	MimeImageGIF:  {Mime: MimeImageGIF, Ext: ExtGif},
	MimeImageWebP: {Mime: MimeImageWebP, Ext: ExtWebp},
}

var AllowedTextType = map[Mime]TypeResolver{
	MimeDocDoc:  {Mime: MimeDocDoc, Ext: ExtDoc},
	MimeDocDocx: {Mime: MimeDocDocx, Ext: ExtDocx},
	MimeDocPptx: {Mime: MimeDocPptx, Ext: ExtPptx},
	MimeDocXlsx: {Mime: MimeDocXlsx, Ext: ExtXlsx},
	MimeDocPDF:  {Mime: MimeDocPDF, Ext: ExtPdf},
}

var AllowedVideoType = map[Mime]TypeResolver{
	MimeVideoMP4:  {Mime: MimeVideoMP4, Ext: ExtMp4},
	MimeVideoMOV:  {Mime: MimeVideoMOV, Ext: ExtMov},
	MimeVideoAVI:  {Mime: MimeVideoAVI, Ext: ExtAvi},
	MimeVideoMKV:  {Mime: MimeVideoMKV, Ext: ExtMkv},
	MimeVideoWebM: {Mime: MimeVideoWebM, Ext: ExtWebm},
	MimeVideoFLV:  {Mime: MimeVideoFLV, Ext: ExtFlv},
}

func isAllowedExt(mime Mime) (TypeResolver, bool) {

	if mime, ok := AllowedImageType[mime]; ok {
		return mime, true
	}
	if mime, ok := AllowedTextType[mime]; ok {
		return mime, true
	}
	if mime, ok := AllowedVideoType[mime]; ok {
		return mime, true
	}
	return TypeResolver{}, false
}

func isAllowedMimeByPolicy(payload AssetPolicy, mime Mime) bool {
	for _, m := range payload.AllowedMime {
		if m == mime {
			return true
		}
	}
	return false
}
