package ioutils

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// NewAtomicFileWriter returns WriteCloser so that writing to it writes to a
// temporary file and closing it atomically changes the temporary file to
// destination path. Writing and closing concurrently is not allowed.
func NewAtomicFileWriter(filename string, perm os.FileMode) (io.WriteCloser, error) {
	f, err := ioutil.TempFile(filepath.Dir(filename), ".tmp-"+filepath.Base(filename))
	if err != nil {
		return nil, err
	}
<<<<<<< HEAD
=======

>>>>>>> 12a5469... start on swarm services; move to glade
	abspath, err := filepath.Abs(filename)
	if err != nil {
		return nil, err
	}
	return &atomicFileWriter{
<<<<<<< HEAD
		f:  f,
		fn: abspath,
=======
		f:    f,
		fn:   abspath,
		perm: perm,
>>>>>>> 12a5469... start on swarm services; move to glade
	}, nil
}

// AtomicWriteFile atomically writes data to a file named by filename.
func AtomicWriteFile(filename string, data []byte, perm os.FileMode) error {
	f, err := NewAtomicFileWriter(filename, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
<<<<<<< HEAD
=======
		f.(*atomicFileWriter).writeErr = err
>>>>>>> 12a5469... start on swarm services; move to glade
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

type atomicFileWriter struct {
	f        *os.File
	fn       string
	writeErr error
<<<<<<< HEAD
=======
	perm     os.FileMode
>>>>>>> 12a5469... start on swarm services; move to glade
}

func (w *atomicFileWriter) Write(dt []byte) (int, error) {
	n, err := w.f.Write(dt)
	if err != nil {
		w.writeErr = err
	}
	return n, err
}

func (w *atomicFileWriter) Close() (retErr error) {
	defer func() {
<<<<<<< HEAD
		if retErr != nil {
=======
		if retErr != nil || w.writeErr != nil {
>>>>>>> 12a5469... start on swarm services; move to glade
			os.Remove(w.f.Name())
		}
	}()
	if err := w.f.Sync(); err != nil {
		w.f.Close()
		return err
	}
	if err := w.f.Close(); err != nil {
		return err
	}
<<<<<<< HEAD
=======
	if err := os.Chmod(w.f.Name(), w.perm); err != nil {
		return err
	}
>>>>>>> 12a5469... start on swarm services; move to glade
	if w.writeErr == nil {
		return os.Rename(w.f.Name(), w.fn)
	}
	return nil
}
