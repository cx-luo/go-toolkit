// Package file provides file operation utilities
package file

import (
	"bufio"
	"io"
	"os"
)

// ReadLinesStream reads a file line by line and calls the callback for each line
// This is memory-efficient for large files as it doesn't load the entire file into memory
func ReadLinesStream(filePath string, callback func(line string, lineNum int) error) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		if err := callback(scanner.Text(), lineNum); err != nil {
			return err
		}
	}
	return scanner.Err()
}

// ReadLinesChannel reads a file line by line and sends each line to a channel
// The channel will be closed when the file is fully read or an error occurs
func ReadLinesChannel(filePath string, bufferSize int) (<-chan string, <-chan error) {
	lines := make(chan string, bufferSize)
	errChan := make(chan error, 1)

	go func() {
		defer close(lines)
		defer close(errChan)

		file, err := os.Open(filePath)
		if err != nil {
			errChan <- err
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines <- scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			errChan <- err
		}
	}()

	return lines, errChan
}

// ReadChunk reads a file in chunks of specified size
// Returns the chunk data, number of bytes read, and any error
func ReadChunk(filePath string, offset int64, chunkSize int) ([]byte, int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, 0, err
	}
	defer file.Close()

	// Seek to the specified offset
	if offset > 0 {
		_, err = file.Seek(offset, io.SeekStart)
		if err != nil {
			return nil, 0, err
		}
	}

	buffer := make([]byte, chunkSize)
	n, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		return nil, 0, err
	}

	return buffer[:n], n, nil
}

// ReadChunksStream reads a file in chunks and calls the callback for each chunk
// This is memory-efficient for large files
func ReadChunksStream(filePath string, chunkSize int, callback func(chunk []byte, offset int64) error) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	buffer := make([]byte, chunkSize)
	offset := int64(0)

	for {
		n, err := file.Read(buffer)
		if n > 0 {
			if err := callback(buffer[:n], offset); err != nil {
				return err
			}
			offset += int64(n)
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}

	return nil
}

// ReadChunksChannel reads a file in chunks and sends each chunk to a channel
// The channel will be closed when the file is fully read or an error occurs
func ReadChunksChannel(filePath string, chunkSize int, bufferSize int) (<-chan Chunk, <-chan error) {
	chunks := make(chan Chunk, bufferSize)
	errChan := make(chan error, 1)

	go func() {
		defer close(chunks)
		defer close(errChan)

		file, err := os.Open(filePath)
		if err != nil {
			errChan <- err
			return
		}
		defer file.Close()

		buffer := make([]byte, chunkSize)
		offset := int64(0)

		for {
			n, err := file.Read(buffer)
			if n > 0 {
				chunk := make([]byte, n)
				copy(chunk, buffer[:n])
				chunks <- Chunk{
					Data:   chunk,
					Offset: offset,
					Size:   n,
				}
				offset += int64(n)
			}

			if err == io.EOF {
				break
			}
			if err != nil {
				errChan <- err
				return
			}
		}
	}()

	return chunks, errChan
}

// Chunk represents a chunk of file data
type Chunk struct {
	Data   []byte // The chunk data
	Offset int64  // The offset in the file where this chunk starts
	Size   int    // The size of the chunk
}

// ReadLinesWithLimit reads a file line by line with a maximum line limit
// Useful when you only need to process the first N lines of a large file
func ReadLinesWithLimit(filePath string, maxLines int) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	lineCount := 0

	for scanner.Scan() && lineCount < maxLines {
		lines = append(lines, scanner.Text())
		lineCount++
	}

	return lines, scanner.Err()
}

// ReadChunkWithOffset reads a specific chunk of a file starting at the given offset
func ReadChunkWithOffset(filePath string, offset int64, size int) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Seek to the specified offset
	_, err = file.Seek(offset, io.SeekStart)
	if err != nil {
		return nil, err
	}

	buffer := make([]byte, size)
	n, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		return nil, err
	}

	return buffer[:n], nil
}
