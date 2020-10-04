package archiver_test

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/jsnjack/archiver/v3"
)

func TestExtract_file(t *testing.T) {
	source := "testdata/test.zip"

	destination, err := ioutil.TempDir("", "archiver_tar_test")
	if err != nil {
		t.Fatalf("creating temp dir: %v", err)
	}
	defer os.RemoveAll(destination)

	err = archiver.Extract(source, "test.txt", destination)
	if err != nil {
		t.Fatalf("unarchiving '%s' to '%s': %v", source, destination, err)
	}

	requireRegularFile(t, path.Join(destination, "test.txt"))
}
