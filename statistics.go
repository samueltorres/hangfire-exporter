package main

// Statistics is an interface for getting hangfire database statistics
type Statistics interface {

	// Available checks if the hangfire database is reachable
	Available() bool

	// DeletedJobs gets the total number of deleted jobs on the hangfire database
	DeletedJobs() float64

	// EnqueuedJobs gets the current number of enqueued jobs
	EnqueuedJobs() float64

	// FetchedJobs gets the current number of fetched jobs
	FetchedJobs() float64

	// FailedJobs gets the total number of failed jobs
	FailedJobs() float64

	// ProcessingJobs gets the current number of processing jobs
	ProcessingJobs() float64

	// Queues gets the current number of queues
	Queues() float64

	// RecurringJobs gets the current number of recurring jobs
	RecurringJobs() float64

	// ScheduledJobs gets the current nubmer of scheduled jobs
	ScheduledJobs() float64

	// Servers gets the number of registered servers on hangfire database
	Servers() float64

	// SucceededJobs gets the total number of succeeded jobs
	SucceededJobs() float64
}
