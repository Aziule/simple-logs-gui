package listener

import (
	"fmt"
	"os"

	"github.com/aziule/simple-logs-gui/backend/log"
	"github.com/hpcloud/tail"
)

type locallyListenedLogFile struct {
	*logFileInfo
	TailChan *tail.Tail `json:"-"`
}

// createLocallyListenedFile creates an instance of locallyListenedLogFile and initialises it
// but without starting to listen to it
func createLocallyListenedFile(path string) ListenedLogFile {
	hash := generateHash(fmt.Sprintf("local-%s", path))

	return &locallyListenedLogFile{
		logFileInfo: &logFileInfo{
			Hash: hash,
			Path: path,
		},
	}
}

// Listen starts listening for incoming logs and stores them
func (f *locallyListenedLogFile) Listen() error {
	tailChan, err := tail.TailFile(f.Path, tail.Config{
		Follow:   true,
		Location: &tail.SeekInfo{Whence: os.SEEK_END},
	})

	if err != nil {
		return ErrCannotOpen
	}

	f.TailChan = tailChan

	for line := range f.TailChan.Lines {
		parsedLog := log.ParseString(line.Text)
		parsedLog.Id = len(f.Logs) + 1

		f.Logs = append(f.Logs, parsedLog)
	}

	return nil
}

func (f *locallyListenedLogFile) StopListening() {
	f.TailChan.StopAtEOF()
}
