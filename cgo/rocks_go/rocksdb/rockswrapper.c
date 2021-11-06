#include <assert.h>

#include "rocksdb/c.h"
#if defined(OS_WIN)
#include <Windows.h>
#else
#include <unistd.h> // sysconf() - get CPU count
#endif

void initDB(char *path, rocksdb_t **db, rocksdb_options_t **options, rocksdb_writeoptions_t **writeoptions, rocksdb_readoptions_t **readoptions)
{
  rocksdb_backup_engine_t *be;
  *options = rocksdb_options_create();
#if defined(OS_WIN)
  SYSTEM_INFO system_info;
  GetSystemInfo(&system_info);
  long cpus = system_info.dwNumberOfProcessors;
#else
  long cpus = sysconf(_SC_NPROCESSORS_ONLN);
#endif
  // Set # of online cores
  rocksdb_options_increase_parallelism(*options, (int)(cpus));
  rocksdb_options_optimize_level_style_compaction(*options, 0);
  // create the DB if it's not already present
  rocksdb_options_set_create_if_missing(*options, 1);

  // open DB
  char *err = NULL;
  *db = rocksdb_open(*options, path, &err);
  assert(!err);
  *writeoptions = rocksdb_writeoptions_create();
  *readoptions = rocksdb_readoptions_create();
}

void closeDB(rocksdb_t *db, rocksdb_options_t *options, rocksdb_writeoptions_t *writeoptions, rocksdb_readoptions_t *readoptions)
{
  rocksdb_writeoptions_destroy(writeoptions);
  rocksdb_readoptions_destroy(readoptions);
  rocksdb_options_destroy(options);
  rocksdb_close(db);
}
