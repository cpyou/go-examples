package file

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

var SyncFileLock sync.RWMutex

func Write(data []byte) error {
	SyncFileLock.Lock()
	defer SyncFileLock.Unlock()
	f, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0766)
	defer f.Close()
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if err != nil {
		return err
	}
	time.Sleep(1 * time.Second)
	return nil
}

func Read() (data string, err error) {
	SyncFileLock.RLock()
	defer SyncFileLock.RUnlock()
	f, err := os.Open("test.txt")
	defer f.Close()
	buf := make([]byte, 1)
	if err != nil {
		return
	}
	l, err := f.Read(buf)
	fmt.Println(l)
	if err != nil {
		return
	}
	time.Sleep(1 * time.Second)
	return string(buf), nil
}

func ReadWrite() (err error) {
	SyncFileLock.Lock()
	defer SyncFileLock.Unlock()
	f, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0766)
	defer f.Close()
	if err != nil {
		return err
	}
	buf := make([]byte, 1)
	if err != nil {
		return
	}
	l, err := f.Read(buf)
	fmt.Println(l)
	if err != nil {
		return
	}
	count, err := strconv.Atoi(string(buf))
	count += 1
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, err = f.Write([]byte{byte(count)})
	if err != nil {
		return err
	}
	time.Sleep(1 * time.Second)
	return nil
}
