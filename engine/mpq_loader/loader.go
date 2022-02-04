package mpq_loader

import (
	"fmt"
	"io"
	"strings"

	mpq "github.com/OpenDiablo2/mpq/pkg"
)

type MpqLoader struct {
	mpq *mpq.MPQ
}

func (m *MpqLoader) Name() string {
	return "MPQ Loader"
}

func (m *MpqLoader) Exists(path string) bool {
	if len(path) == 0 {
		return false
	}

	path = strings.ReplaceAll(path, "/", "\\")

	return m.mpq.Contains(path)
}

func (m *MpqLoader) Load(path string) (io.ReadSeeker, error) {
	path = strings.ReplaceAll(path, "/", "\\")

	if !m.Exists(path) {
		return nil, fmt.Errorf("could not locate file %q in %q", path, m.mpq.Path())
	}

	return m.mpq.ReadFileStream(path)
}

func New(fileName string) (*MpqLoader, error) {
	result := &MpqLoader{}

	mpq, err := mpq.FromFile(fileName)

	if err != nil {
		return nil, err
	}

	result.mpq = mpq

	return result, nil
}
