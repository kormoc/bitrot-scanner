package main

import "os"

func enqueuePath(path string, info os.FileInfo, err error) error {
	Trace.Printf("Enqueueing job %v\n", path)
	jobs <- job{path, info, err}
	return nil
}
