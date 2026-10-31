package schemamigrator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogsMigrations(t *testing.T) {
	manager := newTestMigrationManager(t)
	checkSchemaMigratorMigrationRecords(t, manager, LogsMigrations)
}

func TestLogsMigrationsExactNature(t *testing.T) {
	manager := newTestMigrationManager(t)
	checkSchemaMigrationRecordExactNature(
		t,
		manager,
		[]SchemaMigrationRecord{
			LogsMigrations[1],
			LogsMigrations[2],
			LogsMigrations[4],
			LogsMigrations[5],
			LogsMigrations[6], // 2001 (sync)
			LogsMigrations[8], // 2003 (sync)
		},
		[]SchemaMigrationRecord{
			LogsMigrations[0], // 1000 (async)
			LogsMigrations[3], // 1003 (async)
		},
	)
}

func TestLogsMigrationResourceCreatedAt(t *testing.T) {
	manager := newTestMigrationManager(t)
	migration := migrationByID(t, LogsMigrations, 2003)

	assert.True(t, manager.IsSync(migration))
	assert.Equal(t, []string{
		"ALTER TABLE signoz_logs.logs_v2_resource ON CLUSTER cluster ADD COLUMN IF NOT EXISTS created_at DateTime64(3) MATERIALIZED now64(3) CODEC(Delta(8), ZSTD(1))",
		"ALTER TABLE signoz_logs.distributed_logs_v2_resource ON CLUSTER cluster ADD COLUMN IF NOT EXISTS created_at DateTime64(3) MATERIALIZED now64(3) CODEC(Delta(8), ZSTD(1))",
	}, operationsSQL(migration.UpItems))
	assert.Equal(t, []string{
		"ALTER TABLE signoz_logs.distributed_logs_v2_resource ON CLUSTER cluster DROP COLUMN IF EXISTS created_at",
		"ALTER TABLE signoz_logs.logs_v2_resource ON CLUSTER cluster DROP COLUMN IF EXISTS created_at",
	}, operationsSQL(migration.DownItems))
}
