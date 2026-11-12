package localstorage_test

import (
	"io"
	"path"
	"profile-template/internal/assets"
	"profile-template/internal/config"
	"profile-template/internal/localstorage"
	"strings"
	"testing"
	"testing/iotest"

	"github.com/google/uuid"
)

type UsecaseContract string

const (
	Write  UsecaseContract = "WRITE"
	Delete UsecaseContract = "DELETE"
	Read   UsecaseContract = "READ"
)

type WriteTestCase struct {
	Name        string
	Usecase     UsecaseContract
	ExpectedErr error
	Destination string
	Reader      io.Reader
}

func createWriteTestUsecase(conf config.StorageConfig) []WriteTestCase {
	content := uuid.NewString() + "-_" + ".txt"
	return []WriteTestCase{
		{Name: "VALID_WRITE_TEST", Usecase: Write, ExpectedErr: nil, Destination: path.Join(conf.StoragePathTemp, "1_"+content), Reader: strings.NewReader("hello world")},
		{Name: "VALID_READ_TEST", Usecase: Read, ExpectedErr: nil, Destination: path.Join(conf.StoragePathTemp, "1_"+content), Reader: strings.NewReader("hello world")},
		{Name: "VALID_DELETE_TEST", Usecase: Delete, ExpectedErr: nil, Destination: path.Join(conf.StoragePathTemp, "1_"+content), Reader: strings.NewReader("hello world")},

		{Name: "EMPTY_CONTENT_WRITE_TEST", Usecase: Write, ExpectedErr: assets.ErrAssetFailedCreate, Destination: path.Join(conf.StoragePathTemp, "2_"+content), Reader: strings.NewReader("")},
		{Name: "EMPTY_CONTENT_READ_TEST", Usecase: Read, ExpectedErr: assets.ErrAssetNotFound, Destination: path.Join(conf.StoragePathTemp, "2_"+content), Reader: strings.NewReader("")},
		{Name: "EMPTY_CONTENT_DELETE_TEST", Usecase: Delete, ExpectedErr: assets.ErrAssetNotFound, Destination: path.Join(conf.StoragePathTemp, "2_"+content), Reader: strings.NewReader("")},

		{Name: "EMPTY_DESTINATION_WRITE_TEST", Usecase: Write, ExpectedErr: assets.ErrAssetInvalidPath, Destination: "", Reader: strings.NewReader("hello world")},
		{Name: "EMPTY_DESTINATION_READ_TEST", Usecase: Read, ExpectedErr: assets.ErrAssetInvalidPath, Destination: "", Reader: strings.NewReader("hello world")},
		{Name: "EMPTY_DESTINATION_DELETE_TEST", Usecase: Delete, ExpectedErr: assets.ErrAssetInvalidPath, Destination: "", Reader: strings.NewReader("hello world")},

		{Name: "INVALID_DESTINATION_WRITE_TEST", Usecase: Write, ExpectedErr: assets.ErrAssetInvalidPath, Destination: path.Join("adas", "4_"+content), Reader: strings.NewReader("hello world")},
		{Name: "INVALID_DESTINATION_READ_TEST", Usecase: Read, ExpectedErr: assets.ErrAssetInvalidPath, Destination: path.Join("adas", "4_"+content), Reader: strings.NewReader("hello world")},
		{Name: "INVALID_DESTINATION_DELETE_TEST", Usecase: Delete, ExpectedErr: assets.ErrAssetInvalidPath, Destination: path.Join("adas", "4_"+content), Reader: strings.NewReader("hello world")},

		{Name: "ERROR_READER_WRITE_TEST", Usecase: Write, ExpectedErr: assets.ErrAssetFailedCreate, Destination: path.Join(conf.StoragePathTemp, "5_"+content), Reader: iotest.ErrReader(assets.ErrAssetFailedCreate)},
		{Name: "ERROR_READER_READ_TEST", Usecase: Read, ExpectedErr: assets.ErrAssetNotFound, Destination: path.Join(conf.StoragePathTemp, "5_"+content), Reader: iotest.ErrReader(assets.ErrAssetFailedCreate)},
		{Name: "ERROR_READER_DELETE_TEST", Usecase: Delete, ExpectedErr: assets.ErrAssetNotFound, Destination: path.Join(conf.StoragePathTemp, "5_"+content), Reader: iotest.ErrReader(assets.ErrAssetFailedCreate)},

		{Name: "MULTIPLE_INVALID_DESTINATION_WRITE_TEST", Usecase: Write, ExpectedErr: assets.ErrAssetInvalidPath, Destination: path.Join("adas", "asdsad", "6_"+content), Reader: strings.NewReader("hello world")},
		{Name: "MULTIPLE_INVALID_DESTINATION_READ_TEST", Usecase: Read, ExpectedErr: assets.ErrAssetInvalidPath, Destination: path.Join("adas", "asdsad", "6_"+content), Reader: strings.NewReader("hello world")},
		{Name: "MULTIPLE_INVALID_DESTINATION_DELETE_TEST", Usecase: Delete, ExpectedErr: assets.ErrAssetInvalidPath, Destination: path.Join("adas", "asdsad", "6_"+content), Reader: strings.NewReader("hello world")},
	}
}

func newLocalStorage(t *testing.T, conf config.StorageConfig) assets.RepositoryContract {
	storge, err := localstorage.NewLocalStorage(conf)

	if err != nil {
		t.Fatal("Storage Failed Init")
	}
	return storge

}

func createExpectedDir(conf config.StorageConfig) []string {
	return []string{
		path.Join(conf.StorageRoot, conf.StoragePathTemp),
		path.Join(conf.StorageRoot, conf.StoragePathUpload),
		path.Join(conf.StorageRoot, conf.StoragePathDocument),
		path.Join(conf.StorageRoot, conf.StoragePathAvatar),
		path.Join(conf.StorageRoot, conf.StoragePathVideo),
	}
}

func TestWriteReadDelete(t *testing.T) {
	conf := config.StorageConfig{
		StorageRoot:         t.TempDir(),
		StoragePathTemp:     "temp",
		StoragePathUpload:   "upload",
		StoragePathDocument: "document",
		StoragePathAvatar:   "avatar",
		StoragePathVideo:    "video",
	}
	localStorage := newLocalStorage(t, conf)
	uc := createWriteTestUsecase(conf)

	for _, u := range uc {
		switch u.Usecase {
		case Write:
			t.Run(u.Name, func(t *testing.T) {
				_, writeErr := localStorage.Write(t.Context(), u.Reader, u.Destination)
				if writeErr != u.ExpectedErr {
					t.Fatalf("Write: expected =%v, got err=%v", u.ExpectedErr, writeErr)
				}
			})
		case Read:
			t.Run(u.Name, func(t *testing.T) {
				f, readErr := localStorage.Read(t.Context(), u.Destination)

				if readErr != u.ExpectedErr {
					t.Fatalf("Read: expected =%v, got err=%v", u.ExpectedErr, readErr)
				}
				if f != nil {
					defer f.Close()
				}
			})
		case Delete:
			t.Run(u.Name, func(t *testing.T) {
				delErr := localStorage.Delete(t.Context(), u.Destination)

				if delErr != u.ExpectedErr {
					t.Fatalf("Delete: expected =%v, got err=%v", u.ExpectedErr, delErr)
				}
			})
		}
	}

}
