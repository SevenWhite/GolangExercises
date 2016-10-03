package main

import (
	"bytes"
	"fmt"
	"io"
)

type WriterWrap struct {
	writer io.Writer
	count  int64
}

func (w *WriterWrap) Write(p []byte) (int, error) {
	count, err := w.writer.Write(p)

	if err != nil {
		return count, err
	}

	w.count += int64(count)

	return count, nil
}

func main() {
	var buf bytes.Buffer
	writer, count := CountingWriter(&buf)

	writer.Write([]byte("hello"))

	fmt.Println(*count)
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	newWriter := &WriterWrap{writer: w}
	return newWriter, &newWriter.count
}
