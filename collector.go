package main

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

var (
	namespace = "hangfire"
)

type hangfireCollector struct {
	sync.Mutex
	statistics           Statistics
	logger               *log.Logger
	up                   *prometheus.Desc
	serverCounter        *prometheus.Desc
	deletedJobsCounter   *prometheus.Desc
	enqueuedJobsGauge    *prometheus.Desc
	failedJobsCounter    *prometheus.Desc
	fetchedJobsGauge     *prometheus.Desc
	processingJobsGauge  *prometheus.Desc
	queuesGauge          *prometheus.Desc
	recurringJobsGauge   *prometheus.Desc
	scheduledJobsGauge   *prometheus.Desc
	succeededJobsCounter *prometheus.Desc
}

func newHangfireCollector(logger *log.Logger, sts Statistics) *hangfireCollector {
	return &hangfireCollector{
		statistics: sts,
		logger:     logger,
		up: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "up"),
			"Can hangfire database be reached",
			nil,
			nil,
		),
		serverCounter: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "servers_total"),
			"Number of registered hangfire servers",
			nil,
			nil,
		),
		deletedJobsCounter: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "deleted_jobs_total"),
			"Total number of deleted jobs",
			nil,
			nil,
		),
		enqueuedJobsGauge: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "enqueued_jobs_total"),
			"Current number of enqueued jobs",
			nil,
			nil,
		),
		failedJobsCounter: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "failed_jobs_total"),
			"Total number of failed jobs",
			nil,
			nil,
		),
		fetchedJobsGauge: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "fetched_jobs_total"),
			"Current number of fetched jobs",
			nil,
			nil,
		),
		processingJobsGauge: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "processing_jobs_total"),
			"Current number of processing jobs",
			nil,
			nil,
		),
		queuesGauge: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "queues_total"),
			"Number of queues",
			nil,
			nil,
		),
		recurringJobsGauge: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "recurring_jobs_total"),
			"Current number of recurring jobs",
			nil,
			nil,
		),
		scheduledJobsGauge: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "scheduled_jobs_total"),
			"Current number of scheduled jobs",
			nil,
			nil,
		),
		succeededJobsCounter: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "succeeded_jobs_total"),
			"Total number of succeeded jobs",
			nil,
			nil,
		),
	}
}

func (h *hangfireCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- h.serverCounter
	ch <- h.deletedJobsCounter
	ch <- h.enqueuedJobsGauge
	ch <- h.enqueuedJobsGauge
	ch <- h.failedJobsCounter
	ch <- h.processingJobsGauge
	ch <- h.queuesGauge
	ch <- h.recurringJobsGauge
	ch <- h.scheduledJobsGauge
	ch <- h.succeededJobsCounter
}

func (h *hangfireCollector) Collect(ch chan<- prometheus.Metric) {
	h.Lock()
	defer h.Unlock()

	if !h.statistics.Available() {
		ch <- prometheus.MustNewConstMetric(
			h.up,
			prometheus.GaugeValue,
			0,
		)

		return
	}

	ch <- prometheus.MustNewConstMetric(
		h.up,
		prometheus.GaugeValue,
		1,
	)

	ch <- prometheus.MustNewConstMetric(
		h.serverCounter,
		prometheus.CounterValue,
		h.statistics.Servers())

	ch <- prometheus.MustNewConstMetric(
		h.deletedJobsCounter,
		prometheus.CounterValue,
		h.statistics.DeletedJobs())

	ch <- prometheus.MustNewConstMetric(
		h.enqueuedJobsGauge,
		prometheus.GaugeValue,
		h.statistics.EnqueuedJobs())

	ch <- prometheus.MustNewConstMetric(
		h.failedJobsCounter,
		prometheus.CounterValue,
		h.statistics.FailedJobs())

	ch <- prometheus.MustNewConstMetric(
		h.fetchedJobsGauge,
		prometheus.CounterValue,
		h.statistics.FetchedJobs())

	ch <- prometheus.MustNewConstMetric(
		h.processingJobsGauge,
		prometheus.GaugeValue,
		h.statistics.ProcessingJobs())

	ch <- prometheus.MustNewConstMetric(
		h.queuesGauge,
		prometheus.GaugeValue,
		h.statistics.Queues())

	ch <- prometheus.MustNewConstMetric(
		h.recurringJobsGauge,
		prometheus.GaugeValue,
		h.statistics.RecurringJobs())

	ch <- prometheus.MustNewConstMetric(
		h.scheduledJobsGauge,
		prometheus.GaugeValue,
		h.statistics.ScheduledJobs())

	ch <- prometheus.MustNewConstMetric(
		h.succeededJobsCounter,
		prometheus.CounterValue,
		h.statistics.SucceededJobs())
}
