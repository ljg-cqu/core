package utils

import (
	"github.com/spf13/afero"
	"os"
	"testing"
)

func TestExist(t *testing.T) {
	appFS := afero.NewMemMapFs()
	// create test files and directories
	appFS.MkdirAll("src/a", 0755)

	afero.WriteFile(appFS, "src/a/b", []byte("file b"), 0644)
	afero.WriteFile(appFS, "src/c", []byte("file c"), 0644)
	afero.WriteFile(appFS, "阿斯顿撒多撒.pdf", []byte("ddddd"), 0644)
	name := "src/c"
	_, err := appFS.Stat(name)
	if os.IsNotExist(err) {
		t.Errorf("file \"%s\" does not exist.\n", name)
	}

	name2 := "阿斯顿撒多撒.pdf"
	_, err = appFS.Stat(name2)
	if os.IsNotExist(err) {
		t.Errorf("file \"%s\" does not exist.\n", name2)
	}
}
