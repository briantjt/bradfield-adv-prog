package rocksdb

/*
#cgo CFLAGS: -I/usr/include
// #cgo LDFLAGS: -L/usr/lib ${SRCDIR}/librocksdb.a -lstdc++ -lpthread -lpthread -lsnappy  -lz -lbz2 -lzstd -llz4 -ldl -luring -lm
#cgo LDFLAGS: -L/usr/lib -lrocksdb
#include <string.h>
#include "rocksdb/c.h"
#include "rockswrapper.h"
*/
import "C"

type rocksdb struct {
	db           *C.rocksdb_t
	options      *C.rocksdb_options_t
	writeOptions *C.rocksdb_writeoptions_t
	readOptions  *C.rocksdb_readoptions_t
}

func NewRocksDB(path string) *rocksdb {
	var rocksDB rocksdb
	C.initDB(C.CString(path), &rocksDB.db, &rocksDB.options, &rocksDB.writeOptions, &rocksDB.readOptions)
	return &rocksDB
}

func (rdb *rocksdb) Close() {
	C.closeDB(rdb.db, rdb.options, rdb.writeOptions, rdb.readOptions)
}

func (rdb *rocksdb) Put(key, value string) {
	c_key := C.CString(key)
	c_val := C.CString(value)
	var err *C.char
	C.rocksdb_put(rdb.db, rdb.writeOptions, c_key, C.strlen(c_key), c_val, C.strlen(c_val),
		&err)
}

func (rdb *rocksdb) Get(key string) string {
	c_key := C.CString(key)
	var err *C.char
	var len C.size_t
	returned_value := C.rocksdb_get(rdb.db, rdb.readOptions, c_key, C.strlen(c_key), &len, &err)
	return C.GoString(returned_value)
}
