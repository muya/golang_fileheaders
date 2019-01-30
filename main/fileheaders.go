package main

import (
	"fmt"
	"os"
)

func main() {
	/*
	   - create some sample files - zip, gzip, tar
	   - print out first characters of different files to see the types
	*/
	// expand this to use array
	baseDir := "./sample_files/"
	filesToParse := []string{
		fmt.Sprintf("%s%s", baseDir, "dummy.txt"),
		fmt.Sprintf("%s%s", baseDir, "dummy.txt.gz"),
		fmt.Sprintf("%s%s", baseDir, "dummy.txt.zip"),
		fmt.Sprintf("%s%s", baseDir, "dummy.txt.tar"),
		fmt.Sprintf("%s%s", baseDir, "dummy_dir.tar"),
	}

	fmt.Println(fmt.Sprintf("Apparent zip file sequence: %v", []byte("\x50\x4B\x03\x04")))
	fmt.Println(fmt.Sprintf("Apparent gzip file sequence: %v", []byte("\x1F\x8B\x08")))
	fmt.Println(fmt.Sprintf("Apparent tar file sequence: %v", []byte("\x75\x73\x74\x61\x72\x00\x30\x30")))

	for _, fileName := range filesToParse {
		if fh, err := os.Open(fileName); err != nil {
			fmt.Println(fmt.Sprintf("Error while opening file %s - error: %v", fileName, err))

			return
		} else {
			fmt.Println(fmt.Sprintf("file handler: %v", fh))

			tmpBuff := make([]byte, 32)

			if _, err := fh.Read(tmpBuff); err != nil {
				fmt.Println(fmt.Sprintf("error while reading file %v", err))
				fh.Close()
				return
			}

			// rewind the file
			_, _ = fh.Seek(0, 0)
			fmt.Println(fileName)
			fmt.Println(fmt.Sprintf("first 4 (zip): %v", (tmpBuff[:4])))
			fmt.Println(fmt.Sprintf("first 3 (gzip): %v", (tmpBuff[:3])))
			fmt.Println(fmt.Sprintf("first 8 (tar): %v", (tmpBuff[:8])))

			_ = fh.Close()
		}
	}

}
