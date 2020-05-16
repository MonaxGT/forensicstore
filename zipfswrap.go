package forensicstore

import (
	"archive/zip"
	"os"

	"github.com/spf13/afero"
	"github.com/spf13/afero/zipfs"
)

func newZIPFSWrapper(path string) (*zipFSWrap, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	info, err := f.Stat()
	if err != nil {
		return nil, err
	}
	zr, err := zip.NewReader(f, info.Size())
	if err != nil {
		return nil, err
	}
	return &zipFSWrap{f, zipfs.New(zr)}, nil
}

type zipFSWrap struct {
	base *os.File
	afero.Fs
}

func (fs *zipFSWrap) Close() error {
	return fs.base.Close()
}
