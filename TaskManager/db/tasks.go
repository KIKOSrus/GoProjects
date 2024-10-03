package db

import (
	"encoding/binary"
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

var tasks = []byte("Tasks")
var db *bolt.DB

func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(tasks)
		return err
	})
}

func CreateTask(task string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(tasks)
		id64, _ := b.NextSequence()
		key := itob(int(id64))
		return b.Put([]byte(key), []byte(task))
	})
}

func itob(i int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(i))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
func ViewTasks() error {
	return db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(tasks)
		c := b.Cursor()
		fmt.Println("ID             TASK")
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("%d              %s\n", btoi(k), v)
		}
		return nil
	})
}

func DeleteTask(id int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(tasks)
		return b.Delete(itob(id))
	})
}
