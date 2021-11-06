#include "rocksdb/c.h"

void initDB(char *path, rocksdb_t **db, rocksdb_options_t **options, rocksdb_writeoptions_t **writeoptions, rocksdb_readoptions_t **readoptions);
void closeDB(rocksdb_t *db, rocksdb_options_t *options, rocksdb_writeoptions_t *writeoptions, rocksdb_readoptions_t *readoptions);
