package sqlserver

import (
	"database/sql"

	log "github.com/sirupsen/logrus"
)

type SqlServerStatistics struct {
	db     *sql.DB
	logger *log.Logger
}

func NewSqlServerStatistics(connection string, logger *log.Logger) (*SqlServerStatistics, error) {
	db, err := sql.Open("sqlserver", connection)
	if err != nil {
		return nil, err
	}
	return &SqlServerStatistics{
		db:     db,
		logger: logger,
	}, nil
}

// Available checks if the hangfire database is reachable
func (s *SqlServerStatistics) Available() bool {
	err := s.db.Ping()

	if err != nil {
		s.logger.Error("error pinging sql server", err)
		return false
	}

	return true
}

// DeletedJobs gets the total number of deleted jobs on the hangfire database
func (s *SqlServerStatistics) DeletedJobs() float64 {
	return 0
}

// EnqueuedJobs gets the current number of enqueued jobs
func (s *SqlServerStatistics) EnqueuedJobs() float64 {
	return 0
}

// FetchedJobs gets the current number of fetched jobs
func (s *SqlServerStatistics) FetchedJobs() float64 {
	return 0
}

// FailedJobs gets the total number of failed jobs
func (s *SqlServerStatistics) FailedJobs() float64 {
	return 0
}

// ProcessingJobs gets the current number of processing jobs
func (s *SqlServerStatistics) ProcessingJobs() float64 {
	return 0
}

// Queues gets the current number of queues
func (s *SqlServerStatistics) Queues() float64 {
	return 0
}

// RecurringJobs gets the current number of recurring jobs
func (s *SqlServerStatistics) RecurringJobs() float64 {
	return 0
}

// ScheduledJobs gets the current nubmer of scheduled jobs
func (s *SqlServerStatistics) ScheduledJobs() float64 {
	return 0
}

// Servers gets the number of registered servers on hangfire database
func (s *SqlServerStatistics) Servers() float64 {
	var count int

	row := s.db.QueryRow("select count(Id) from Hangfire.Server with (nolock);")
	err := row.Scan(&count)
	if err != nil {
		s.logger.Error("error getting server", err)
		return 0
	}

	return float64(count)
}

// SucceededJobs gets the total number of succeeded jobs
func (s *SqlServerStatistics) SucceededJobs() float64 {
	return 0
}

// string sql = String.Format(@"
// set transaction isolation level read committed;
// select count(Id) from [{0}].Job with (nolock, forceseek) where StateName = N'Enqueued';
// select count(Id) from [{0}].Job with (nolock, forceseek) where StateName = N'Failed';
// select count(Id) from [{0}].Job with (nolock, forceseek) where StateName = N'Processing';
// select count(Id) from [{0}].Job with (nolock, forceseek) where StateName = N'Scheduled';
// select count(Id) from [{0}].Server with (nolock);
// select sum(s.[Value]) from (
//     select sum([Value]) as [Value] from [{0}].Counter with (nolock, forceseek) where [Key] = N'stats:succeeded'
//     union all
//     select [Value] from [{0}].AggregatedCounter with (nolock, forceseek) where [Key] = N'stats:succeeded'
// ) as s;
// select sum(s.[Value]) from (
//     select sum([Value]) as [Value] from [{0}].Counter with (nolock, forceseek) where [Key] = N'stats:deleted'
//     union all
//     select [Value] from [{0}].AggregatedCounter with (nolock, forceseek) where [Key] = N'stats:deleted'
// ) as s;
// select count(*) from [{0}].[Set] with (nolock, forceseek) where [Key] = N'recurring-jobs';
//                 ", _storage.SchemaName);
