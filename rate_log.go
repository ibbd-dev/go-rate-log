package rateLog

import (
	"fmt"
	"io"
	"sync"
	"time"
)

type RateLogger struct {
	mu     sync.Mutex // ensures atomic writes; protects the following fields
	prefix string     // prefix to write at beginning of each line
	flag   string     // properties，time.Time.Format的参数
	out    io.Writer  // destination for output
	buf    []byte     // for accumulating text to write

	timeMu   sync.Mutex
	duration time.Duration
	lastTime time.Time
}

// New creates a new Logger. The out variable sets the
// destination to which log data will be written.
// The prefix appears at the beginning of each generated log line.
// The flag argument defines the logging properties. 如time.RFC3339
/*
   ANSIC       = "Mon Jan _2 15:04:05 2006"
   UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
   RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
   RFC822      = "02 Jan 06 15:04 MST"
   RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
   RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
   RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
   RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
   RFC3339     = "2006-01-02T15:04:05Z07:00"
   RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
   Kitchen     = "3:04PM"
   // Handy time stamps.
   Stamp      = "Jan _2 15:04:05"
   StampMilli = "Jan _2 15:04:05.000"
   StampMicro = "Jan _2 15:04:05.000000"
   StampNano  = "Jan _2 15:04:05.000000000"
*/
func New(out io.Writer, prefix string, flag string) *RateLogger {
	return &RateLogger{out: out, prefix: prefix, flag: flag}
}

func (l *RateLogger) formatHeader(buf *[]byte, t time.Time) {
	if l.prefix != "" {
		*buf = append(*buf, l.prefix...)
		*buf = append(*buf, ' ')
	}

	if l.flag != "" {
		*buf = append(*buf, t.Format(l.flag)...)
		*buf = append(*buf, ' ')
	}
}

func (l *RateLogger) Output(s string) error {
	println("==")
	fmt.Println(l.lastTime)
	now := time.Now() // get this early.
	l.timeMu.Lock()
	if l.lastTime.Add(l.duration).After(now) {
		l.timeMu.Unlock()
		fmt.Println(l.lastTime)
		return nil
	}
	l.lastTime = now
	l.timeMu.Unlock()

	l.mu.Lock()
	defer l.mu.Unlock()

	l.buf = l.buf[:0]
	l.formatHeader(&l.buf, now)
	l.buf = append(l.buf, s...)
	if len(s) == 0 || s[len(s)-1] != '\n' {
		l.buf = append(l.buf, '\n')
	}
	fmt.Println(string(l.buf))
	_, err := l.out.Write(l.buf)
	return err
}

// Printf calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
func (l *RateLogger) Printf(format string, v ...interface{}) {
	l.Output(fmt.Sprintf(format, v...))
}

// Print calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Print.
func (l *RateLogger) Print(v ...interface{}) { l.Output(fmt.Sprint(v...)) }

// Println calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Println.
func (l *RateLogger) Println(v ...interface{}) { l.Output(fmt.Sprintln(v...)) }

// SetFlags sets the output flags for the logger.
func (l *RateLogger) SetFlags(flag string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.flag = flag
}

// SetPrefix sets the output prefix for the logger.
func (l *RateLogger) SetPrefix(prefix string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.prefix = prefix
}

// SetOutput sets the output destination for the logger.
func (l *RateLogger) SetOutput(w io.Writer) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.out = w
}
