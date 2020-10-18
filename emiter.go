package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/awrenn/hal/parser"

	"github.com/joomcode/errorx"
)

func processFile(conf HalConfig, src, dst string) error {
	// TODO acwrenn pass in formatter from conf
	out, err := openFile(filepath.Join(conf.Dst, dst), "js-beautify")
	if err != nil {
		return err
	}
	defer out.Close()
	indentCount := 0
	refHistory := make([]string, 0)
	return processBytesIntoWriter(conf, src, out, indentCount, refHistory)
}

func processBytesIntoWriter(conf HalConfig, src string, out io.Writer, indentCount int, refHistory []string) error {
	fname := filepath.Join(conf.Src, src)
	raw, err := ioutil.ReadFile(fname)
	if err != nil {
		return errorx.Decorate(err, fmt.Sprintf("Failed to read file %s", fname))
	}

	for {
		var tag parser.Tag
		var err error
		tag, raw, err = parser.NextTag(raw)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		if tag.Kind == parser.TagClosing {
			indentCount -= 1
		}

		if tag.Kind == parser.TagReference {
			refed := tag.Inner.Reference()
			if Contains(refHistory, refed) {
				return parser.ReferenceCycle.New(fmt.Sprintf("Reference cycle detected in file %s - make sure nothing it depends, on depends on it", src))
			}
			refHistory = append(refHistory, tag.Inner.Reference())
			err := processBytesIntoWriter(conf, refed, out, indentCount, refHistory)
			if err != nil {
				return err
			}
		} else {
			_, err = io.Copy(out, tag.Reader(indentCount))
			if err != nil {
				return err
			}
		}

		if tag.Kind == parser.TagOpening {
			indentCount += 1
		}
	}
}

func Contains(his []string, src string) bool {
	for _, h := range his {
		if h == src {
			return true
		}
	}
	return false
}

func openFile(fname, formaterCommand string) (io.WriteCloser, error) {
	f, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return nil, err
	}
	if formaterCommand != "" {
		var args []string
		if formaterCommand == "js-beautify" {
			args = []string{"-f", "-", "--type", "html"}
		}
		cmd := exec.Command(formaterCommand, args...)
		stdin, err := cmd.StdinPipe()
		if err != nil {
			return nil, err
		}
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			return nil, err
		}
		err = cmd.Start()
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		e := newEmitterAdapater(stdin, stdout, f)
		go e.emit()
		return e, nil
	}
	return f, nil
}

type emitterAdapater struct {
	in  io.WriteCloser
	mid io.ReadCloser
	out io.WriteCloser
	ec  chan error
}

func newEmitterAdapater(in io.WriteCloser, mid io.ReadCloser, out io.WriteCloser) *emitterAdapater {
	return &emitterAdapater{
		in:  in,
		mid: mid,
		out: out,
		ec:  make(chan error),
	}

}

func (e *emitterAdapater) emit() {
	_, err := io.Copy(e.out, e.mid)
	e.ec <- err
}

func (e *emitterAdapater) Close() error {
	defer e.out.Close()
	defer e.mid.Close()
	e.in.Close()
	return <-e.ec
}

func (e *emitterAdapater) Write(b []byte) (int, error) {
	return e.in.Write(b)
}
