package assets_test

import (
	"bytes"
	"io"
	"log"
	S3Storage "profile-template/infra/storage/s3"
	"profile-template/internal/assets"
	"profile-template/internal/config"
	"profile-template/internal/db"
	"profile-template/internal/utils"
	"testing"
	"testing/iotest"
)

type UsecaseContract string

const (
	Upload         UsecaseContract = "Upload"
	Delete         UsecaseContract = "Delete"
	Read           UsecaseContract = "Read"
	LinkingAsset   UsecaseContract = "LinkingAsset"
	ReadByParentID UsecaseContract = "ReadByParentID"
)

type UploadTestPayload struct {
	Name        string
	ExpectedErr error
	Usecase     UsecaseContract
	Policy      assets.UploadPolicy
	Reader      io.Reader
}

type ReadTestPayload struct {
	Name        string
	ExpectedErr error
	FileKey     string
}

type LinkingAssetTestPayload struct {
	Name        string
	ExpectedErr error
	AssetID     []string
	ParentID    string
}

type DeleteTestPayload struct {
	Name        string
	ExpectedErr error
	AssetID     []string
}

type TestReport struct {
	Pass        bool
	Name        string
	Error       error
	ExpectedErr error
}

type Mime string

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

var TestBuffer = map[Mime][]byte{
	MimeImageJPEG: {0xFF, 0xD8, 0xFF, 0xE0},
	MimeImagePNG:  {0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A},
	MimeImageGIF:  []byte("GIF89a"),
	MimeImageWebP: append([]byte("RIFF"), append([]byte{0x00, 0x00, 0x00, 0x00}, []byte("WEBP")...)...),

	MimeDocDoc:  {0xD0, 0xCF, 0x11, 0xE0, 0xA1, 0xB1, 0x1A, 0xE1},
	MimeDocDocx: {0x50, 0x4B, 0x03, 0x04},
	MimeDocPptx: {0x50, 0x4B, 0x03, 0x04},
	MimeDocXlsx: {0x50, 0x4B, 0x03, 0x04},
	MimeDocPDF:  []byte("%PDF-1.4"),

	MimeVideoMP4:  {0x00, 0x00, 0x00, 0x18, 0x66, 0x74, 0x79, 0x70, 0x6D, 0x70, 0x34, 0x32},
	MimeVideoMOV:  {0x00, 0x00, 0x00, 0x14, 0x66, 0x74, 0x79, 0x70, 0x71, 0x74, 0x20, 0x20},
	MimeVideoAVI:  {0x52, 0x49, 0x46, 0x46, 0x00, 0x00, 0x00, 0x00, 0x41, 0x56, 0x49, 0x20},
	MimeVideoMKV:  {0x1A, 0x45, 0xDF, 0xA3},
	MimeVideoWebM: {0x1A, 0x45, 0xDF, 0xA3},
	MimeVideoFLV:  {0x46, 0x4C, 0x56, 0x01},
}

func NewUC(t *testing.T) assets.UsecaseContract {
	conf, err := config.LoadConfig()
	if err != nil {
		t.Fatal("Config Failed Init", err)
	}
	db, err := db.NewMySQLDatabase(&conf.DBConfig)
	if err != nil {
		t.Fatal("DB Failed Init", err, conf.DBConfig)
	}
	storage, err := S3Storage.NewS3Storage(t.Context(), conf.StorageConfig)
	if err != nil {
		t.Fatal("Storage Failed Init")
	}
	assetRepo := assets.NewMysqlRepo(db)
	assetManager := assets.NewUsecase(conf.StorageConfig, storage, assetRepo)
	if assetManager == nil {
		t.Fatal("Asset Manager Failed Init")
	}
	return assetManager
}

func CreateUploadUsecase() []UploadTestPayload {
	return []UploadTestPayload{
		{
			Name:        "VALID_CATEGORY_UPLOAD_TEST1",
			ExpectedErr: nil,
			Usecase:     Upload,
			Policy: assets.UploadPolicy{
				Category: assets.CategoryDocument,
				MaxSize:  1024 * 1024 * 10,
				UserID:   "VALID_CATEGORY_UPLOAD_TEST",
			},
			Reader: bytes.NewBuffer(TestBuffer[MimeDocPDF]),
		},
		{
			Name:        "INVALID_CATEGORY_UPLOAD_TEST2",
			ExpectedErr: assets.ErrAssetInvalidMime,
			Usecase:     Upload,
			Policy: assets.UploadPolicy{
				Category: assets.CategoryDocument,
				MaxSize:  1024 * 1024 * 10,
				UserID:   "VALID_CATEGORY_UPLOAD_TEST",
			},
			Reader: bytes.NewBuffer(TestBuffer[MimeVideoMP4]),
		},
		{
			Name:        "INVALID_READER_CATEGORY_UPLOAD_TEST",
			ExpectedErr: assets.ErrAssetInvalidMime,
			Usecase:     Upload,
			Policy: assets.UploadPolicy{
				Category: assets.AssetCategory("invalid"),
				MaxSize:  1024 * 1024 * 10,
				UserID:   "INVALID_READER_MIME_UPLOAD_TEST",
			},
			Reader: bytes.NewBuffer([]byte("%PDF-1.4")),
		},
		{
			Name:        "INVALID_READER_ERROR_UPLOAD_TEST",
			ExpectedErr: assets.ErrAssetFailedCreate,
			Usecase:     Upload,
			Policy: assets.UploadPolicy{
				Category: assets.CategoryDocument,
				MaxSize:  1024 * 1024 * 10,
				UserID:   "INVALID_READER_ERROR_UPLOAD_TEST",
			},
			Reader: iotest.ErrReader(assets.ErrAssetFailedCreate),
		},
		{
			Name:        "INVALID_READER_MIME_UPLOAD_TEST",
			ExpectedErr: assets.ErrAssetInvalidMime,
			Usecase:     Upload,
			Policy: assets.UploadPolicy{
				Category: assets.CategoryDocument,
				MaxSize:  1024 * 1024 * 10,
				UserID:   "INVALID_READER_MIME_UPLOAD_TEST",
			},
			Reader: bytes.NewBuffer([]byte("what is reader?")),
		},
	}
}

func TestUpload(t *testing.T) {
	assetsUC := NewUC(t)
	var validAssets []ReadTestPayload
	var linkingTestPayload []LinkingAssetTestPayload
	var validIDS []string
	var deleteTestPayload []DeleteTestPayload

	var Report []TestReport

	validAssets = append(validAssets, ReadTestPayload{
		Name:        "EMPTY_FILE_KEY_READ_TEST",
		ExpectedErr: assets.ErrAssetNotFound,
		FileKey:     "",
	}, ReadTestPayload{
		Name:        "NOT_EXISTED_CATEGORY_READ_TEST",
		ExpectedErr: assets.ErrAssetNotFound,
		FileKey:     "a1b2c3d4-e5f6-7890-1234-567890abcdef.pdf",
	}, ReadTestPayload{
		Name:        "NESTED_FILE_KEY_READ_TEST",
		ExpectedErr: assets.ErrAssetNotFound,
		FileKey:     "avatar/document/a1b2c3d4-e5f6-7890-1234-567890abcdef.pdf",
	})

	testUploads := CreateUploadUsecase()

	for _, tc := range testUploads {

		t.Run(tc.Name, func(t *testing.T) {
			res, err := assetsUC.Upload(t.Context(), tc.Reader, tc.Policy)
			if err == nil {
				validAssets = append(validAssets, ReadTestPayload{
					Name:        "VALID_READ_TEST",
					ExpectedErr: nil,
					FileKey:     res.FileKey,
				})
				log.Print(tc.Name, res.FileKey)
				validIDS = append(validIDS, res.ID)
			}
			if err != tc.ExpectedErr {
				Report = append(Report, TestReport{
					Pass:        false,
					Name:        tc.Name,
					Error:       err,
					ExpectedErr: tc.ExpectedErr,
				})
				t.Fatalf("Upload: expected =%v, got err=%v", tc.ExpectedErr, err)
			}
			Report = append(Report, TestReport{
				Pass:        true,
				Name:        tc.Name,
				Error:       nil,
				ExpectedErr: nil,
			})
		})
	}
	for _, tc := range validAssets {
		t.Run(tc.Name, func(t *testing.T) {
			f, err := assetsUC.Read(t.Context(), tc.FileKey)
			if err != tc.ExpectedErr {
				Report = append(Report, TestReport{
					Pass:        false,
					Name:        tc.Name,
					Error:       err,
					ExpectedErr: tc.ExpectedErr,
				})
				t.Fatalf("ReadByID: expected =%v, got err=%v,filePath=%s", tc.ExpectedErr, err, tc.FileKey)
			}
			Report = append(Report, TestReport{
				Pass:        true,
				Name:        tc.Name,
				Error:       nil,
				ExpectedErr: nil,
			})
			if f != nil {
				defer f.Close()
			}
		})
	}

	linkingTestPayload = append(linkingTestPayload, LinkingAssetTestPayload{
		Name:        "VALID_LINKING_TEST",
		ExpectedErr: nil,
		AssetID:     validIDS,
		ParentID:    "VALID_LINKING_TEST",
	}, LinkingAssetTestPayload{
		Name:        "INVALID_LINKING_TEST",
		ExpectedErr: assets.ErrAssetNotFound,
		AssetID:     []string{"123e4567-e89b-12d3-a456-426614174000"},
		ParentID:    "INVALID_LINKING_TEST",
	}, LinkingAssetTestPayload{
		Name:        "DUP_PARENT_ID_LINKING_TEST",
		ExpectedErr: assets.ErrAssetNotFound,
		AssetID:     validIDS,
		ParentID:    "DUP_LINKING_TEST",
	}, LinkingAssetTestPayload{
		Name:        "EMPTY_PARENT_ID_LINKING_TEST",
		ExpectedErr: assets.ErrAssetNotFound,
		AssetID:     validIDS,
		ParentID:    "",
	})

	for _, tc := range linkingTestPayload {
		t.Run(tc.Name, func(t *testing.T) {
			err := assetsUC.LinkingAsset(t.Context(), tc.AssetID, tc.ParentID)
			if err != tc.ExpectedErr {
				Report = append(Report, TestReport{
					Pass:        false,
					Name:        tc.Name,
					Error:       err,
					ExpectedErr: tc.ExpectedErr,
				})
				t.Fatalf("LinkingAsset: expected =%v, got err=%v", tc.ExpectedErr, err)
			}
			Report = append(Report, TestReport{
				Pass:        true,
				Name:        tc.Name,
				Error:       nil,
				ExpectedErr: nil,
			})
		})
		t.Run(tc.Name+"_READ_BY_PARENT_ID", func(t *testing.T) {
			_, err := assetsUC.ReadByParentID(t.Context(), []string{tc.ParentID})
			if err != tc.ExpectedErr {
				Report = append(Report, TestReport{
					Pass:        false,
					Name:        tc.Name + "_READ_BY_PARENT_ID",
					Error:       err,
					ExpectedErr: tc.ExpectedErr,
				})
				t.Fatalf("ReadByParentID: expected =%v, got err=%v", tc.ExpectedErr, err)
			}
			Report = append(Report, TestReport{
				Pass:        true,
				Name:        tc.Name + "_READ_BY_PARENT_ID",
				Error:       nil,
				ExpectedErr: nil,
			})
		})

	}
	for _, tc := range linkingTestPayload {
		t.Run(tc.Name+"_RELEASE_BY_PARENT_ID", func(t *testing.T) {
			err := assetsUC.ReleaseByParentID(t.Context(), []string{tc.ParentID})
			if err != tc.ExpectedErr {
				Report = append(Report, TestReport{
					Pass:        false,
					Name:        tc.Name + "_RELEASE_BY_PARENT_ID",
					Error:       err,
					ExpectedErr: tc.ExpectedErr,
				})
				t.Fatalf("ReleaseByParentID: expected =%v, got err=%v", tc.ExpectedErr, err)
			}
			Report = append(Report, TestReport{
				Pass:        true,
				Name:        tc.Name + "_RELEASE_BY_PARENT_ID",
				Error:       nil,
				ExpectedErr: nil,
			})
		})
	}

	deleteTestPayload = append(deleteTestPayload, DeleteTestPayload{
		Name:        "VALID_DELETE_TEST",
		ExpectedErr: nil,
		AssetID:     validIDS,
	}, DeleteTestPayload{
		Name:        "INVALID_DELETE_TEST",
		ExpectedErr: assets.ErrAssetNotFound,
		AssetID:     []string{"123e4567-e89b-12d3-a456-426614174000"},
	}, DeleteTestPayload{
		Name:        "DUP_PARENT_ID_DELETE_TEST",
		ExpectedErr: assets.ErrAssetNotFound,
		AssetID:     validIDS,
	}, DeleteTestPayload{
		Name:        "EMPTY_PARENT_ID_DELETE_TEST",
		ExpectedErr: assets.ErrAssetNotFound,
		AssetID:     validIDS,
	})

	for _, tc := range deleteTestPayload {
		t.Run(tc.Name, func(t *testing.T) {
			err := assetsUC.Delete(t.Context(), assets.DeletePayload{IDs: tc.AssetID})
			if err != tc.ExpectedErr {
				Report = append(Report, TestReport{
					Pass:        false,
					Name:        tc.Name,
					Error:       err,
					ExpectedErr: tc.ExpectedErr,
				})
				t.Fatalf("Delete: expected =%v, got err=%v", tc.ExpectedErr, err)
			}
			Report = append(Report, TestReport{
				Pass:        true,
				Name:        tc.Name,
				Error:       nil,
				ExpectedErr: nil,
			})
		})
	}

	PassReaport := utils.MapField(Report, func(tr TestReport) (TestReport, bool) {
		if tr.Pass {
			return tr, true
		}
		return tr, false
	})
	FailedReport := utils.MapField(Report, func(tr TestReport) (TestReport, bool) {
		if !tr.Pass {
			return tr, true
		}
		return tr, false
	})
	t.Log("Total Test: ", len(Report))
	t.Log("Pass Test: ", len(PassReaport))
	t.Log("Failed Test: ", len(FailedReport))

}

func TestReleaseByFileKeys(t *testing.T) {
	assetsUC := NewUC(t)
	var validFileKeys string
	var validAssets string
	uc := UploadTestPayload{
		Name:        "VALID_CATEGORY_UPLOAD_TEST1",
		ExpectedErr: nil,
		Usecase:     Upload,
		Policy: assets.UploadPolicy{
			Category: assets.CategoryDocument,
			MaxSize:  1024 * 1024 * 10,
			UserID:   "VALID_CATEGORY_UPLOAD_TEST",
		},
		Reader: bytes.NewBuffer(TestBuffer[MimeDocPDF]),
	}
	parentID := "asdasdwasd"

	t.Run("UPLOAD_VALID_ASSETS", func(t *testing.T) {
		res, err := assetsUC.Upload(t.Context(), uc.Reader, uc.Policy)
		if err != nil {
			t.Fatal("Should Be Success")
		}
		validFileKeys = res.FileKey
		validAssets = res.ID
	})

	t.Run("LINKING_ASSET_BY_ASSETS_ID", func(t *testing.T) {
		err := assetsUC.LinkingAsset(t.Context(), []string{validAssets}, parentID)
		if err != nil {
			t.Fatal("Should Be Success")
		}
	})
	t.Run("EMPTY_FILE_KEY_RELEASE", func(t *testing.T) {
		err := assetsUC.ReleaseByFileKeys(t.Context(), []string{""})
		if err == nil {
			t.Fatal("Should Be Error")
		}
	})
	t.Run("VALID_FILE_KEY_RELEASE", func(t *testing.T) {
		err := assetsUC.ReleaseByFileKeys(t.Context(), []string{validFileKeys})
		if err != nil {
			t.Fatal("Should Be Success")
		}
	})
	t.Run("DUP_FILE_KEY_RELEASE", func(t *testing.T) {
		err := assetsUC.ReleaseByFileKeys(t.Context(), []string{validFileKeys})
		if err == nil {
			t.Fatal("Should Be Error")
		}
	})

	t.Run("CLEANING_TEMP", func(t *testing.T) {
		err := assetsUC.Delete(t.Context(), assets.DeletePayload{IDs: []string{validAssets}})
		if err != nil {
			t.Fatal("Should Be Error")
		}
	})

}
func TestReleaseByFileKeysAndLinkingByFileKeys(t *testing.T) {
	assetsUC := NewUC(t)
	var validFileKeys string
	var validAssets string
	uc := UploadTestPayload{
		Name:        "VALID_CATEGORY_UPLOAD_TEST1",
		ExpectedErr: nil,
		Usecase:     Upload,
		Policy: assets.UploadPolicy{
			Category: assets.CategoryDocument,
			MaxSize:  1024 * 1024 * 10,
			UserID:   "VALID_CATEGORY_UPLOAD_TEST",
		},
		Reader: bytes.NewBuffer(TestBuffer[MimeDocPDF]),
	}
	parentID := "asdasdwasd"

	t.Run("UPLOAD_VALID_ASSETS", func(t *testing.T) {
		res, err := assetsUC.Upload(t.Context(), uc.Reader, uc.Policy)
		if err != nil {
			t.Fatal("Should Be Success")
		}
		validFileKeys = res.FileKey
		validAssets = res.ID
		log.Print("UPLOAD_VALID_ASSETS", res)
	})

	t.Run("LINKING_ASSET_BY_FILE_KEY", func(t *testing.T) {
		err := assetsUC.LinkingAssetByFileKeys(t.Context(), []string{validFileKeys}, parentID)
		if err != nil {
			log.Print("LINKING_ASSET_BY_FILE_KEY", err)
			t.Fatal("Should Be Success")
		}
	})
	t.Run("EMPTY_FILE_KEY_RELEASE", func(t *testing.T) {
		err := assetsUC.ReleaseByFileKeys(t.Context(), []string{""})
		if err == nil {
			t.Fatal("Should Be Error")
		}
	})
	t.Run("VALID_FILE_KEY_RELEASE", func(t *testing.T) {
		err := assetsUC.ReleaseByFileKeys(t.Context(), []string{validFileKeys})
		if err != nil {
			log.Print(err)
			t.Fatal("Should Be Success")
		}
	})
	t.Run("DUP_FILE_KEY_RELEASE", func(t *testing.T) {
		err := assetsUC.ReleaseByFileKeys(t.Context(), []string{validFileKeys})
		if err == nil {
			t.Fatal("Should Be Error")
		}
	})

	t.Run("CLEANING_TEMP", func(t *testing.T) {
		err := assetsUC.Delete(t.Context(), assets.DeletePayload{IDs: []string{validAssets}})
		if err != nil {
			log.Print(err)
			t.Fatal("Should Be Error")
		}
	})

}
