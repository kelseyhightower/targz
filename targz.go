package targz

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
)

func Create(root, out string) error {
	o, err := os.Create(out)
	if err != nil {
		return err
	}
	defer o.Close()
	gw := gzip.NewWriter(o)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()

	err = filepath.Walk(root, walkTarFunc(tw))
	if err != nil {
		return err
	}
	return nil
}

func walkTarFunc(tw *tar.Writer) filepath.WalkFunc {
	fn := func(path string, info os.FileInfo, err error) error {
		if info.Mode().IsDir() {
			return nil
		}
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()
		hdr, err := tar.FileInfoHeader(info, "")
		hdr.Name = path
		if err != nil {
			return err
		}
		err = tw.WriteHeader(hdr)
		if err != nil {
			return err
		}
		_, err = io.Copy(tw, f)
		if err != nil {
			return err
		}
		return nil
	}
	return fn
}
