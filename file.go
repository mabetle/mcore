package mcore

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// PrepareFile
func PrepareFile(location string) error {
	if IsFileExist(location) {
		return nil
	}
	_, err := WriteFile(location, "")
	return err
}

// file may not exist.
// if file not exist, create one.
func NewFileWriter(location string) (io.Writer, error) {
	fs, err := os.Create(location)
	if err == os.ErrNotExist {
		fs.Write([]byte{})
		logger.Debugf("File:%s not exist, create one.")
	} else if err != nil {
		return nil, err
	}
	return bufio.NewWriter(fs), nil
}

// get file modified time
func GetFileModifyTime(file string) (int64, error) {
	f, e := os.Stat(file)
	if e != nil {
		return 0, e
	}
	return f.ModTime().Unix(), nil
}

// get file size as how many bytes
func GetFileSize(file string) (int64, error) {
	f, e := os.Stat(file)
	if e != nil {
		return 0, e
	}
	return f.Size(), nil
}

// delete file
func RemoveFile(file string) error {
	return os.Remove(file)
}

// rename file name
func RenameFile(file string, to string) error {
	return os.Rename(file, to)
}

// move file
// if the newPath not exist, create one
// make sure move can success.
func MoveFile(oldPath, newPath string) error {
	//TODO not finished yet.
	println("MoveFile not implement")
	return nil
}

// WriteFileLines
func WriteFileLines(file string, lines []string) (int, error) {
	content := StringArray(lines).MergeWithOsNewLine()
	return WriteFile(file, content)
}

// WriteFileBytes
func WriteFileBytes(file string, v []byte) error {
	r := bytes.NewBuffer(v)
	return WriteFileReader(file, r)
}

// WriteFileReader
func WriteFileReader(file string, r io.Reader) error {
	if err := PrepareFileDir(file); err != nil {
		return err
	}
	w, errW := NewFileWriter(file)
	if errW != nil {
		return errW
	}

	if _, err := io.Copy(w, r); err != nil {
		return err
	}
	return nil
}

func PrepareFileDir(file string) error {
	dir := GetFileDir(file)
	// exist, skip
	if IsFileExist(dir) {
		return nil
	}
	if err := MakeDir(dir); err != nil {
		logger.Debugf("Make directory error: Directory: %s Error: %v", dir, err)
		return err
	}
	return nil
}

func ProcessDir(dir string) string {
	if IsWindows() {
		return strings.Replace(dir, "/", "\\", -1)
	} else {
		return strings.Replace(dir, "\\", "/", -1)
	}
}

// put string to file
func WriteFile(file string, content string) (n int, err error) {
	defer func() {
		logger.Debugf("Write to file: %s, bytes: %d, ErrMsg: %v \n", file, n, err)
	}()
	file = ProcessDir(file)
	//prepare dir
	if err := PrepareFileDir(file); err != nil {
		return 0, err
	}
	fs, e := os.Create(file)
	if e != nil {
		logger.Debugf("Create file error: File: %s, Error: %v", file, e)
		return 0, e
	}
	defer fs.Close()

	n, err = fs.WriteString(content)
	return
}

// WriteFileGBK
func WriteFileGBK(file, text string) (int, error) {
	return WriteFile(file, EncodeGBK(text))
}

func AppendFile(file string, appendContent string) (int, error) {
	return AppendFileByLocation(file, appendContent, false)
}

func AppendFileBefore(file string, appendContent string) (int, error) {
	return AppendFileByLocation(file, appendContent, true)
}

func AppendFileByLocation(file string, appendContent string, before bool) (int, error) {
	if IsFileExist(file) {
		content, err := ReadFileAll(file)
		if err != nil {
			return 0, err
		}
		if before {
			content = appendContent + content
		} else {
			content = content + appendContent
		}
		return WriteFile(file, content)
	}
	//file not exist
	return WriteFile(file, appendContent)
}

// same to GetContent
func ReadFileAll(file string) (string, error) {
	return GetFileContent(file)
}

// get string from text file
func GetFileContent(file string) (string, error) {
	file = ProcessDir(file)
	if !IsFile(file) {
		return "", os.ErrNotExist
	}
	b, e := ioutil.ReadFile(file)
	if e != nil {
		return "", e
	}
	return string(b), nil
}

// it returns false when it's a directory or does not exist.
func IsFile(file string) bool {
	f, e := os.Stat(file)
	if e != nil {
		return false
	}
	return !f.IsDir()
}

// IsExist returns whether a file or directory exists.
func IsFileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// TouchFile
func TouchFile(file string) error {
	if IsFileExist(file) {
		// TODO update modify time
		println("not implement TouchFile when exists.")
		return nil
	}
	//not exist, create blank file
	_, err := WriteFile(file, "")
	return err
}

// GetFileDir
func GetFileDir(location string) string {
	vs := strings.Split(location, "/")
	sb := NewStringBuffer()
	sb.Append("/")
	for index, v := range vs {
		if v == "" {
			continue
		}
		if index < len(vs)-1 {
			sb.Append(v, "/")
		}
	}
	return sb.String()
}

// MakeDir make dir
func MakeDir(dir string) error {
	if IsFileExist(dir) {
		return fmt.Errorf("%s is exist", dir)
	}
	if err := os.MkdirAll(dir, 0777); err != nil {
		if os.IsPermission(err) {
			return fmt.Errorf("No enough permission to create dir")
		}
		return err
	}
	return nil
}

// GetPathFiles returns all file in dir
func GetDirSubFiles(dir string) (fs []string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}
	for _, f := range files {
		if f.IsDir() {
			// skip dir
			continue
		}
		fs = append(fs, dir+"/"+f.Name())
	}
	return
}

// GetSubFiles returns dir sub files
// exts: format .xx format.
// r: recursive
func GetSubFiles(
	dir string,
	r bool,
	exts string,
	skipDirs string,
	skipFiles string,
) []string {
	var result []string
	err := GetSubFilesImpl(&result, dir, r, exts, skipDirs, skipFiles)
	if err != nil {
		logger.Warn(err)
	}
	return result
}

// GetSubFilesImpl
func GetSubFilesImpl(
	result *[]string,
	dir string,
	r bool,
	exts string,
	skipDirs string,
	skipFiles string,
) error {
	dir = String(dir).TrimEnd("/").String()
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		return err
	}

	extsArray := strings.Split(exts, ",")

	for _, file := range files {
		// skip start with . dir or file
		if strings.HasPrefix(file.Name(), ".") {
			continue
		}
		// full match
		if skipDirs != "" && String(file.Name()).IsContainsInSepStringIgnoreCase(skipDirs, ",") {
			continue
		}
		if file.IsDir() {
			if r {
				GetSubFilesImpl(result, dir+"/"+file.Name(), r, exts, skipDirs, skipFiles)
			}
		} else {
			name := dir + "/" + file.Name()
			if skipFiles != "" && String(file.Name()).IsContainsInSepStringIgnoreCase(skipFiles, ",") {
				continue
			}
			// has exts
			if len(extsArray) > 0 {
				for _, ext := range extsArray {
					if String(name).IsEndWith(ext) {
						*result = append(*result, name)
					}
				}
			} else {
				// no exts
				*result = append(*result, name)
			}
		}
	}
	return nil
}

// WriteFileWithError get content from somewhere may be occur error
// when error is not nil, skip write
func WriteFileWithError(location, content string, err error) error {
	if err != nil {
		return err
	}
	_, err = WriteFile(location, content)
	return err
}
