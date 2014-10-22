package main

import (
	"archive/zip"
	"bytes"
	"github.com/cheggaaa/pb"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func main() {
	url := "https://nodeload.github.com/goagent/goagent/legacy.zip/3.0"
	location := "/home/wangcl"
	if len(os.Args) == 3 {
		location = os.Args[1]
		url = os.Args[2]
	}
	log.Println("begin downlaod goagent")
	log.Printf("Url:%s,Location:%s\n", url, location)
	resp, err := http.Get(url)
	checkerr(err)
	defer resp.Body.Close()
	err = unzip(resp, location)
	checkerr(err)
	log.Println("download end!")
}
func unzip(resp *http.Response, dest string) error {
	i, _ := strconv.Atoi(resp.Header.Get("Content-Length"))
	sourceSize := int64(i)
	source := resp.Body

	bar := pb.New(int(sourceSize)).SetUnits(pb.U_BYTES).SetRefreshRate(time.Millisecond * 10)
	bar.ShowSpeed = true
	bar.Start()
	reader := bar.NewProxyReader(source)

	b, err := ioutil.ReadAll(reader)
	checkerr(err)

	bar.Finish()

	readerAt := bytes.NewReader(b)
	r, err := zip.NewReader(readerAt, int64(len(b)))
	if err != nil {
		return err
	}
	log.Println("download sucessful, now extracting...")

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()
		fn := strings.Join(strings.Split(f.Name, "/")[1:], "/")
		path := filepath.Join(dest, "goagent", fn)
		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			f, err := os.OpenFile(
				path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
func Pipeline(cmds ...*exec.Cmd) (pipeLineOutput, collectedStandardError []byte, pipeLineError error) {
	// Require at least one command
	if len(cmds) < 1 {
		return nil, nil, nil
	}

	// Collect the output from the command(s)
	var output bytes.Buffer
	var stderr bytes.Buffer

	last := len(cmds) - 1
	for i, cmd := range cmds[:last] {
		var err error
		// Connect each command's stdin to the previous command's stdout
		if cmds[i+1].Stdin, err = cmd.StdoutPipe(); err != nil {
			return nil, nil, err
		}
		// Connect each command's stderr to a buffer
		cmd.Stderr = &stderr
	}

	// Connect the output and error for the last command
	cmds[last].Stdout, cmds[last].Stderr = &output, &stderr

	// Start each command
	for _, cmd := range cmds {
		if err := cmd.Start(); err != nil {
			return output.Bytes(), stderr.Bytes(), err
		}
	}

	// Wait for each command to complete
	for _, cmd := range cmds {
		if err := cmd.Wait(); err != nil {
			return output.Bytes(), stderr.Bytes(), err
		}
	}

	// Return the pipeline output and the collected standard error
	return output.Bytes(), stderr.Bytes(), nil
}

func checkerr(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
