package db

import (
	"path"
	"runtime"
)

func Aaaa() {
	//dbPath := getCurrentAbPathByCaller() + "/rocksdb"
	//bbto := gorocksdb.NewDefaultBlockBasedTableOptions()
	//bbto.SetBlockCache(gorocksdb.NewLRUCache(3 << 30))
	//opts := gorocksdb.NewDefaultOptions()
	//opts.SetBlockBasedTableFactory(bbto)
	//opts.SetCreateIfMissing(true)
	//db, err := gorocksdb.OpenDb(opts, dbPath)
	//fmt.Println(db, err)
}

func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
