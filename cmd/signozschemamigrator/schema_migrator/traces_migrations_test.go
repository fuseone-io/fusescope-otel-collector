package schemamigrator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTracesMigrations(t *testing.T) {
	manager := newTestMigrationManager(t)
	checkSchemaMigratorMigrationRecords(t, manager, TracesMigrations)
}

func TestTracesMigrationsExactNature(t *testing.T) {
	manager := newTestMigrationManager(t)
	checkSchemaMigrationRecordExactNature(
		t,
		manager,
		[]SchemaMigrationRecord{
			TracesMigrations[0],
			TracesMigrations[2],
			TracesMigrations[4],
			TracesMigrations[5],
			TracesMigrations[6],
			TracesMigrations[7],
			TracesMigrations[8],
			TracesMigrations[11],
			TracesMigrations[12],
			TracesMigrations[13],
			TracesMigrations[14],
			TracesMigrations[15],
			TracesMigrations[16],
			TracesMigrations[17],
			TracesMigrations[18],
			TracesMigrations[19],
		},
		[]SchemaMigrationRecord{
			TracesMigrations[1], // 1001 (async)
			TracesMigrations[3], // 1003 (async)
		},
	)
}

func TestTracesMigrationResourceCreatedAt(t *testing.T) {
	manager := newTestMigrationManager(t)
	migration := migrationByID(t, TracesMigrations, 1019)

	assert.True(t, manager.IsSync(migration))
	assert.Equal(t, []string{
		"ALTER TABLE signoz_traces.traces_v3_resource ON CLUSTER cluster ADD COLUMN IF NOT EXISTS created_at DateTime64(3) MATERIALIZED now64(3) CODEC(Delta(8), ZSTD(1))",
		"ALTER TABLE signoz_traces.distributed_traces_v3_resource ON CLUSTER cluster ADD COLUMN IF NOT EXISTS created_at DateTime64(3) MATERIALIZED now64(3) CODEC(Delta(8), ZSTD(1))",
		"ALTER TABLE signoz_traces.signoz_error_index_v2 ON CLUSTER cluster ADD COLUMN IF NOT EXISTS created_at DateTime64(3) MATERIALIZED now64(3) CODEC(Delta(8), ZSTD(1))",
		"ALTER TABLE signoz_traces.distributed_signoz_error_index_v2 ON CLUSTER cluster ADD COLUMN IF NOT EXISTS created_at DateTime64(3) MATERIALIZED now64(3) CODEC(Delta(8), ZSTD(1))",
	}, operationsSQL(migration.UpItems))
	assert.Equal(t, []string{
		"ALTER TABLE signoz_traces.distributed_traces_v3_resource ON CLUSTER cluster DROP COLUMN IF EXISTS created_at",
		"ALTER TABLE signoz_traces.traces_v3_resource ON CLUSTER cluster DROP COLUMN IF EXISTS created_at",
		"ALTER TABLE signoz_traces.distributed_signoz_error_index_v2 ON CLUSTER cluster DROP COLUMN IF EXISTS created_at",
		"ALTER TABLE signoz_traces.signoz_error_index_v2 ON CLUSTER cluster DROP COLUMN IF EXISTS created_at",
	}, operationsSQL(migration.DownItems))
}
