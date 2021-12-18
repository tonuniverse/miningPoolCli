/*
miningPoolCli – open-source tonuniverse mining pool client

Copyright (C) 2021 tonuniverse.com

This file is part of miningPoolCli.

miningPoolCli is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

miningPoolCli is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with miningPoolCli.  If not, see <https://www.gnu.org/licenses/>.
*/

package helpers

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"io"
	"log"
	"miningPoolCli/utils/mlog"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func ExecuteSimpleCommand(name string, arg ...string) []byte {
	stdout, err := exec.Command(name, arg...).Output()
	if err != nil {
		mlog.LogFatal("Error while executing sh: " + "\"" + name + "\"" + "; " + err.Error())
	}
	return stdout
}

func ExtractTarGz(gzipStream io.Reader, pathToExtarct string) {
	uncompressedStream, err := gzip.NewReader(gzipStream)
	if err != nil {
		log.Fatal("ExtractTarGz: NewReader failed")
	}

	tarReader := tar.NewReader(uncompressedStream)

	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("ExtractTarGz: Next() failed: %s", err.Error())
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.Mkdir(pathToExtarct+"/"+header.Name, 0755); err != nil {
				log.Fatalf("ExtractTarGz: Mkdir() failed: %s", err.Error())
			}
		case tar.TypeReg:
			outFile, err := os.Create(pathToExtarct + "/" + header.Name)
			if err != nil {
				log.Fatalf("ExtractTarGz: Create() failed: %s", err.Error())
			}
			if _, err := io.Copy(outFile, tarReader); err != nil {
				log.Fatalf("ExtractTarGz: Copy() failed: %s", err.Error())
			}
			outFile.Close()

		default:
			mlog.LogFatal("ExtractTarGz: uknown type: " + string(header.Typeflag) + " in " + header.Name)
		}

	}
}

// Extract ZIP files, but skip all directories in zip
func ExtractZip(filename string, dst string) {
	archive, err := zip.OpenReader(filename)
	if err != nil {
		panic(err)
	}
	defer archive.Close()

	for _, f := range archive.File {
		filePath := filepath.Join(dst, f.Name)

		if !strings.HasPrefix(filePath, filepath.Clean(dst)+string(os.PathSeparator)) {
			mlog.LogError("ExtractZip invalid file path: " + filePath)
			return
		}

		c := filepath.Join(dst, filepath.Base(filePath))
		dstFile, err := os.OpenFile(c, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			panic(err)
		}

		fileInArchive, err := f.Open()
		if err != nil {
			mlog.LogFatalStackError(err)
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			mlog.LogFatalStackError(err)
		}

		dstFile.Close()
		fileInArchive.Close()
	}
}
