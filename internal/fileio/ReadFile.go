package fileio

import "os"

func ReadWholeFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Get the file size
	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	// Read the file
	bytes := make([]byte, stat.Size())
	_, err = file.Read(bytes)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
