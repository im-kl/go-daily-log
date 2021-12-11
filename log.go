/*
A thin wrapper around the standard logger. In addition to Writer of the
standard logger, go-daily-log also writes to log files timestamped in the
format of 2006-01-02, and rotates them daily.

	package main

	import (
		"net/http"

		"github.com/kok-leong-chan/go-daily-log"
	)

	func main() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.Method, r.URL.Path)

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		})

		http.ListenAndServe(":8080", nil)
	}
*/
package log

import (
	"io"
	"log"
	"os"
	"path"
	"sync"
	"time"
)

var (
	file   *os.File
	lock   = &sync.Mutex{}
	writer = log.Writer()

	dir    = "log"
	flag   = os.O_WRONLY | os.O_APPEND | os.O_CREATE
	layout = "2006-01-02"
	name   string
	perm   = os.ModePerm
	suffix = ".log"
)

func init() {
	if err := initDir(); err != nil {
		panic(err)
	}

	if err := initFile(); err != nil {
		panic(err)
	}
}

func initDir() error {
	return os.MkdirAll(dir, perm)
}

func initFile() error {
	newName := time.Now().Format(layout) + suffix

	if newName == name {
		return nil
	}

	if file != nil {
		if err := file.Close(); err != nil {
			return err
		}
	}

	name = newName

	var err error

	file, err = os.OpenFile(path.Join(dir, name), flag, perm)

	if err != nil {
		return err
	}

	log.SetOutput(io.MultiWriter(writer, file))

	return nil
}

// Wraps around log.Fatal.
func Fatal(v ...interface{}) {
	lock.Lock()

	defer lock.Unlock()

	if err := initFile(); err != nil {
		panic(err)
	}

	log.Fatal(v...)
}

// Wraps around log.Fatalf.
func Fatalf(format string, v ...interface{}) {
	lock.Lock()

	defer lock.Unlock()

	if err := initFile(); err != nil {
		panic(err)
	}

	log.Fatalf(format, v...)
}

// Wraps around log.Fatalln.
func Fatalln(v ...interface{}) {
	lock.Lock()

	defer lock.Unlock()

	if err := initFile(); err != nil {
		panic(err)
	}

	log.Fatalln(v...)
}

// Wraps around log.Panic.
func Panic(v ...interface{}) {
	lock.Lock()

	defer lock.Unlock()

	if err := initFile(); err != nil {
		panic(err)
	}

	log.Panic(v...)
}

// Wraps around log.Panicf.
func Panicf(format string, v ...interface{}) {
	lock.Lock()

	defer lock.Unlock()

	if err := initFile(); err != nil {
		panic(err)
	}

	log.Panicf(format, v...)
}

// Wraps around log.Panicln.
func Panicln(v ...interface{}) {
	lock.Lock()

	defer lock.Unlock()

	if err := initFile(); err != nil {
		panic(err)
	}

	log.Panicln(v...)
}

// Wraps around log.Print.
func Print(v ...interface{}) {
	lock.Lock()

	defer lock.Unlock()

	if err := initFile(); err != nil {
		panic(err)
	}

	log.Print(v...)
}

// Wraps around log.Printf.
func Printf(format string, v ...interface{}) {
	lock.Lock()

	defer lock.Unlock()

	if err := initFile(); err != nil {
		panic(err)
	}

	log.Printf(format, v...)
}

// Wraps around log.Println.
func Println(v ...interface{}) {
	lock.Lock()

	defer lock.Unlock()

	if err := initFile(); err != nil {
		panic(err)
	}

	log.Println(v...)
}

// Writes log files to the directory specified by d, by default ./log
func SetDir(d string) {
	lock.Lock()

	defer lock.Unlock()

	dir = d
	name = ""

	if err := initDir(); err != nil {
		panic(err)
	}

	if err := initFile(); err != nil {
		panic(err)
	}
}

// Wraps around log.SetPrefix.
func SetPrefix(p string) {
	lock.Lock()

	defer lock.Unlock()

	log.SetPrefix(p)
}
