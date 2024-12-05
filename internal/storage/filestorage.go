package storage

import (
	"bufio"
	"encoding/json"
	"github.com/sheinsviatoslav/shortener/internal/config"
	"io"
	"os"
)

type FileData map[string]string

type FileStorage struct {
	Producer Producer
	Consumer Consumer
}

func NewFileStorage() *FileStorage {
	return &FileStorage{
		Producer: Producer{},
		Consumer: Consumer{},
	}
}

func (fs *FileStorage) GetOriginalURLByShortURL(inputShortURL string) (string, error) {
	urlItems, err := fs.ReadURLItems()
	if err != nil {
		return "", err
	}

	for originalURL, shortURL := range *urlItems {
		if shortURL == inputShortURL {
			return originalURL, nil
		}
	}

	return "", nil
}

func (fs *FileStorage) GetShortURLByOriginalURL(originalURL string) (string, bool, error) {
	urlItems, err := fs.ReadURLItems()
	if err != nil {
		return "", false, err
	}

	if shortURL, ok := (*urlItems)[originalURL]; ok {
		return shortURL, true, nil
	}

	return "", false, nil
}

func (fs *FileStorage) AddNewURL(originalURL string, shortURL string) error {
	urlItems, err := fs.ReadURLItems()
	if err != nil {
		return err
	}

	if _, ok := (*urlItems)[originalURL]; !ok {
		(*urlItems)[originalURL] = shortURL
		if fsError := fs.WriteURLItem(*urlItems); fsError != nil {
			return err
		}
	}

	return nil
}

func (fs *FileStorage) WriteURLItem(urlMap FileData) error {
	var fileWriter, err = NewProducer(*config.FileStoragePath)
	if err != nil {
		return err
	}
	defer fileWriter.Close()

	if err = fileWriter.WriteURLItems(&urlMap); err != nil {
		return err
	}

	return nil
}

func (fs *FileStorage) ReadURLItems() (*FileData, error) {
	if _, err := os.Stat(*config.FileStoragePath); err == nil {
		fileReader, fileErr := NewConsumer(*config.FileStoragePath)
		if fileErr != nil {
			return nil, fileErr
		}

		defer fileReader.Close()

		byteValue, readErr := io.ReadAll(fileReader.reader)
		if readErr != nil {
			return nil, readErr
		}

		urlItems := FileData{}
		if err = json.Unmarshal(byteValue, &urlItems); err != nil {
			return nil, err
		}

		return &urlItems, nil
	}

	return &FileData{}, nil
}

type Producer struct {
	file   *os.File
	writer *bufio.Writer
}

func NewProducer(filename string) (*Producer, error) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	return &Producer{
		file:   file,
		writer: bufio.NewWriter(file),
	}, nil
}

func (p *Producer) WriteURLItems(urlItems *FileData) error {
	data, err := json.MarshalIndent(&urlItems, "", "   ")
	if err != nil {
		return err
	}

	if _, err = p.writer.Write(data); err != nil {
		return err
	}

	if err = p.writer.WriteByte('\n'); err != nil {
		return err
	}

	return p.writer.Flush()
}

func (p *Producer) Close() error {
	return p.file.Close()
}

type Consumer struct {
	file   *os.File
	reader *bufio.Reader
}

func NewConsumer(filename string) (*Consumer, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	return &Consumer{
		file:   file,
		reader: bufio.NewReader(file),
	}, nil
}

func (c *Consumer) Close() error {
	return c.file.Close()
}