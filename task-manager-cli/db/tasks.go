package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

// bucket name
// bucket is equivalent to table in sql databases
var taskBucket = []byte("tasks")
var db *bolt.DB // a package level var

type Task struct {
	Key   int
	Value string
}

// Init with capital I is different from init()
// it wont be called before the package is loaded, its just normal func
func Init(dbPath string) error  {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

func CreateTask(task string) (int, error) {
	var  id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		id = int(id64) // casting because we return int not int64
		key := itob(id)
		return b.Put(key, []byte(task))
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

func AllTasks() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor() // its like linked list
		for k, v := c.First(); k != nil; k, v = c.Next() {
			// Note that every thing in boldDB is []byte, both key and value
			// so we must convert them to int, string ...etc before using them in our app
			tasks = append(tasks, Task{
				Key:   btoi(k),
				Value: string(v),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func DeleteTask(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(itob(key))
	})
}

// itob is integer to byte slice converter
func itob(v int) []byte {
	b := make([]byte, 8)
	// BigEndian = most significant bits
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// btoi is byte slice to integer converter
func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}

