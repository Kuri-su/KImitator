/*
 * 	ConfSyncer - a little sync config files tool in the Linux.
 *     Copyright (C) 2020  Amatist_kurisu<misaki.zhcy@gmail.com>
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package unit

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/mitchellh/go-homedir"
)

// RunCommandInShell call bash shell && return stdout and stderr
func RunCommandInShell(command string) {
	commandWords := strings.Split(command, " ")
	cmd := exec.Command(commandWords[0], commandWords[1:]...)
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	stdinPipe, _ := cmd.StdinPipe()

	var stdoutBuf, stderrBuf bytes.Buffer
	var errStdout, errStderr error
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)
	err := cmd.Start()
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		_, _ = io.Copy(stdout, stdoutIn)
	}()
	go func() {
		_, _ = io.Copy(stderr, stderrIn)
	}()
	go func() {
		_, _ = io.Copy(stdinPipe, os.Stdin)
	}()
	err = cmd.Wait()
	if err != nil {
		log.Fatalln(err)
	}
	if errStdout != nil || errStderr != nil {
		log.Fatal("failed to capture stdout or stderr\n")
	}
	outStr, errStr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())
	fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)
}

// RunCommandInShellGetOutput return output in once time
func RunCommandInShellGetOutput(command string) (string, error) {
	out, err := exec.Command("sh", "-c", command).CombinedOutput()
	if err != nil {
		return string(out), err
	}
	return string(out), nil
}

func MakeDirWithFilePath(path string) error {
	return os.MkdirAll(filepath.Dir(path), 0744)
}

// CreateFile will create file
func CreateFile(filename string) error {
	newFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	newFile.Close()
	return nil
}

// WriteFile is write content into file
func WriteFile(dir string, content []byte) error {
	return ioutil.WriteFile(dir, content, 0744)
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func IsFile(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}

func RealPath(path string) (string, error) {
	// clean '~' char
	realPath, err := homedir.Expand(path)
	if err != nil {
		return "", err
	}

	return realPath, nil
}

// Copy can copy file or directory to dist path
func Copy(src, dist string) error {
	var err error
	// get realPath
	src, err = RealPath(src)
	if err != nil {
		return err
	}
	dist, err = RealPath(dist)
	if err != nil {
		return err
	}

	stat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if stat.IsDir() {
		return CopyDirectory(src, dist)
	} else {
		return CopyFile(src, dist)
	}
}

// CopyFile can copy file to dist path
func CopyFile(src, dist string) error {
	var err error

	// TODO copy dir
	// open source file
	originalFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer originalFile.Close()

	err = MakeDirWithFilePath(dist)
	if err != nil {
		return err
	}

	// create new file
	newFile, err := os.Create(dist)
	if err != nil {
		return err
	}
	defer newFile.Close()
	// copy content
	_, err = io.Copy(newFile, originalFile)
	if err != nil {
		return err
	}
	// flush file content to disk
	err = newFile.Sync()
	if err != nil {
		return err
	}
	return nil
}

// CopyDirectory can copy directory to dist path
func CopyDirectory(srcDir, dist string) error {
	entries, err := ioutil.ReadDir(srcDir)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		sourcePath := filepath.Join(srcDir, entry.Name())
		destPath := filepath.Join(dist, entry.Name())

		fileInfo, err := os.Stat(sourcePath)
		if err != nil {
			return err
		}

		stat, ok := fileInfo.Sys().(*syscall.Stat_t)
		if !ok {
			return fmt.Errorf("failed to get raw syscall.Stat_t data for '%s'", sourcePath)
		}

		switch fileInfo.Mode() & os.ModeType {
		case os.ModeDir:
			if err := CreateIfNotExists(destPath, 0755); err != nil {
				return err
			}
			if err := CopyDirectory(sourcePath, destPath); err != nil {
				return err
			}
		case os.ModeSymlink:
			if err := CopySymLink(sourcePath, destPath); err != nil {
				return err
			}
		default:
			if err := CopyFile(sourcePath, destPath); err != nil {
				return err
			}
		}

		if err := os.Lchown(destPath, int(stat.Uid), int(stat.Gid)); err != nil {
			return err
		}

		isSymlink := entry.Mode()&os.ModeSymlink != 0
		if !isSymlink {
			if err := os.Chmod(destPath, entry.Mode()); err != nil {
				return err
			}
		}
	}
	return nil
}

// Exists to check filePath is exists
func Exists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
}

// CreateIfNotExists create dir if folder is exists
func CreateIfNotExists(dir string, perm os.FileMode) error {
	if Exists(dir) {
		return nil
	}

	if err := os.MkdirAll(dir, perm); err != nil {
		return fmt.Errorf("failed to create directory: '%s', error: '%s'", dir, err.Error())
	}

	return nil
}

// CopySymLink is copy symlink to dist
func CopySymLink(src, dist string) error {
	link, err := os.Readlink(src)
	if err != nil {
		return err
	}
	return os.Symlink(link, dist)
}

// Move is move src to dist
func Move(src, dist string) error {
	return os.Rename(src, dist)
}

// RemoveFiles is remove files
func RemoveFiles(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		return err
	}
	return nil
}
