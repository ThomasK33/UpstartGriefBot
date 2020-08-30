package main

import (
	"bufio"
	"io"
	"os"

	"github.com/reiver/go-oi"
	"github.com/reiver/go-telnet"
)

type telnetCaller struct {
	w *telnet.Writer
}

func (c *telnetCaller) CallTELNET(ctx telnet.Context, w telnet.Writer, r telnet.Reader) {
	c.w = &w

	go func(writer io.Writer, reader io.Reader) {

		var buffer [1]byte // Seems like the length of the buffer needs to be small, otherwise will have to wait for buffer to fill up.
		p := buffer[:]

		for {
			// Read 1 byte.
			n, err := reader.Read(p)
			if n <= 0 && nil == err {
				continue
			} else if n <= 0 && nil != err {
				break
			}

			oi.LongWrite(writer, p)
		}
	}(os.Stdout, r)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		oi.LongWrite(w, scanner.Bytes())
		oi.LongWrite(w, []byte("\n"))
	}
}
