package logging

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
)

/*
type logger interface {
	New(nameDir, locationDir string, filesType []string) (loggingData, error)
	Write(data, typeLogFile string) (int, error)
	ClosingFiles()
}
*/

type LoggingData struct {
	locationDirectory string
	logDescription    map[string]*log.Logger
	filesDescription  map[string]*os.File
}

func NewLoggingData(nameDir, locationDir string, typeFiles []string) (LoggingData, error) {
	dir := path.Join(locationDir, nameDir)
	ld := LoggingData{
		locationDirectory: dir,
	}
	ftmp := make(map[string]*os.File, len(typeFiles))
	ltmp := make(map[string]*log.Logger, len(typeFiles))

	if _, err := os.ReadDir(dir); err != nil {
		if err := os.Mkdir(dir, 0777); err != nil {
			return ld, errors.New("error: it is not possible to create a directory for log files")
		}
	}

	for _, fn := range typeFiles {
		fullFN := path.Join(dir, fn+".log")
		_, err := os.Stat(fullFN)
		if err == nil {
			_ = os.Remove(fullFN)
		}

		f, err := os.OpenFile(fullFN, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return ld, fmt.Errorf("error: it is impossible to create a log file %s", fullFN)
		}

		l := log.New(f, "", log.LstdFlags)
		if fn == "error" {
			l.SetFlags(log.Lshortfile | log.LstdFlags)
		}

		ltmp[fn] = l
		ftmp[fn] = f
	}

	ld.filesDescription = ftmp
	ld.logDescription = ltmp

	return ld, nil
}

func (ld *LoggingData) ClosingFiles() {
	for _, fd := range ld.filesDescription {
		fd.Close()
	}
}

func (ld *LoggingData) GetCountFileDescription() int {
	return len(ld.filesDescription)
}

func (ld LoggingData) GetListTypeFile() []string {
	list := make([]string, 0, len(ld.filesDescription))

	for k, _ := range ld.filesDescription {
		list = append(list, k)
	}

	return list
}

func (ld LoggingData) WriteLoggingData(str, typeLogFile string) bool {
	if ld, ok := ld.logDescription[typeLogFile]; ok {
		ld.Println(str)

		return true
	}

	return false
}
