package main

import (
	"fmt"
	"time"

	bolt "go.etcd.io/bbolt"
)

var (
	db         *bolt.DB
	bucketName = []byte("recorded_times")
)

const (
	fileName = "store.db"
)

func connect() {
	defer logTime()()

	var err error
	db, err = bolt.Open(fileName, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		t := trace(2)
		errOut(err, t)
	}

	var t *Stack
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		if b == nil {
			_, err := tx.CreateBucket(bucketName)
			if err != nil {
				t = trace(2)
				return err
			}
		}

		return nil
	})

	if err != nil {
		errOut(err, t)
	}
}

func checkBucket(tx *bolt.Tx) (*bolt.Bucket, error) {
	b := tx.Bucket(bucketName)
	if b == nil {
		return nil, fmt.Errorf("No bucket found: %s", bucketName)
	} else {
		return b, nil
	}
}

func addEntry(key string, value string) {
	defer logTime()()

	var s *Stack
	err := db.Update(func(tx *bolt.Tx) error {
		b, err := checkBucket(tx)
		if err != nil {
			s = trace(2)
			return err
		}
		err = b.Put([]byte(key), []byte(value))
		if err != nil {
			s = trace(2)
			return err
		}

		return nil
	})

	if err != nil {
		errOut(err, s)
	}
}

func getEntry(key string) string {
	defer logTime()()
	var s *Stack
	var value []byte

	err := db.View(func(tx *bolt.Tx) error {
		b, err := checkBucket(tx)
		if err != nil {
			s = trace(2)
			return err
		}

		value = b.Get([]byte(key))
		return nil
	})
	if err != nil {
		errOut(err, s)
	}

	return string(value)
}
