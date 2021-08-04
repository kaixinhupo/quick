package core

import (
	"archive/zip"
	"io"
	"os"
)

// Compress 压缩文件
// files 文件数组，可以是不同dir下的文件或者文件夹
// dest 压缩文件存放地址
func Compress(dir string, dest string) error {
	d, _ := os.Create(dest)
	defer func(d *os.File) {
		_ = d.Close()
	}(d)
	w := zip.NewWriter(d)
	defer func(w *zip.Writer) {
		_ = w.Close()
	}(w)
	file, _ := os.OpenFile(dir, os.O_RDONLY, os.ModeDir)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	return compress(file, "", w)
}

func compress(file *os.File, prefix string, zw *zip.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		if prefix == "" {
			prefix = "/"
		} else {
			prefix = prefix + "/" + info.Name()
		}
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = compress(f, prefix, zw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := zip.FileInfoHeader(info)
		header.Name = prefix + "/" + header.Name
		header.Method = zip.Deflate
		if err != nil {
			return err
		}
		writer, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, file)
		err = file.Close()
		if err != nil {
			return err
		}
		if err != nil {
			return err
		}
	}
	return nil
}
