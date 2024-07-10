package config

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
)

type Config struct {
	ServerAddress ServerAddress
	BaseURL       BaseURL
	FileStorage   FileStorage
	DB            DB
}

func (o *Config) GetServer() string {
	return o.ServerAddress.String()
}

func (o *Config) GetBaseURL() string {
	return o.BaseURL.String()
}

func (o *Config) GetFileStorage() string { return o.FileStorage.String() }

type ServerAddress struct {
	Host string
	Port int
}

func (s *ServerAddress) String() string {
	return s.Host + ":" + strconv.Itoa(s.Port)
}

func (s *ServerAddress) Set(flag string) error {

	ss := strings.Split(flag, ":")

	if len(ss) != 2 {
		return errors.New("invalid server argument")
	}
	sp, err := strconv.Atoi(ss[1])
	if err != nil {
		return err
	}

	s.Host = ss[0]
	s.Port = sp

	return nil
}

type BaseURL struct {
	Host string
	Port int
}

func (b *BaseURL) String() string {
	return "http://" + b.Host + ":" + strconv.Itoa(b.Port)
}

func (b *BaseURL) Set(flag string) error {

	bu, err := url.Parse(flag)

	if err != nil {
		return err
	}
	bup, err := strconv.Atoi(bu.Port())
	if err != nil {
		return err
	}

	b.Host = bu.Hostname()
	b.Port = bup

	return nil
}

type FileStorage struct {
	file string
}

func (f *FileStorage) String() string {
	return f.file
}

func (f *FileStorage) Set(flag string) error {
	f.file = flag
	return nil
}

type DB struct {
	db string
}

func (d *DB) String() string {
	return d.db
}

func (d *DB) Set(flag string) error {
	d.db = flag
	return nil
}

func NewOptions() Config {
	return Config{
		ServerAddress: ServerAddress{Host: "localhost", Port: 8080},
		BaseURL:       BaseURL{Host: "localhost", Port: 8080},
		FileStorage:   FileStorage{},
		DB:            DB{},
	}
}
