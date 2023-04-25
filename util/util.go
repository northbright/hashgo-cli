package util

import (
	"context"
	"os"

	"github.com/cheggaaa/pb/v3"
	"github.com/northbright/hasher"
)

func Compute(file string, algs []string) (map[string][]byte, int64, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, 0, err
	}
	defer f.Close()

	// Get file info.
	fi, err := f.Stat()
	if err != nil {
		return nil, 0, err
	}

	// Get file size.
	total := fi.Size()

	// Start a new bar.
	bar := pb.Full.Start64(total)
	defer bar.Finish()

	// Create a proxy reader.
	barReader := bar.NewProxyReader(f)

	// Create a hasher.
	h, err := hasher.New(barReader, algs)
	if err != nil {
		return nil, 0, err
	}

	return h.Compute(context.Background())
}
