# Hangfire Exporter

Export Hangfire statistics to Prometheus. This exporter currently supports the MongoDB database.

[![CircleCI](https://circleci.com/gh/samueltorres/hangfire-exporter.svg?style=svg)](https://circleci.com/gh/samueltorres/hangfire-exporter)

## Exported Metrics

| Metric | Meaning |
| ------ | ------- |
| hangfire_up | Can hangfire database be reached
| hangfire_servers_total | Number of registered hangfire servers
| hangfire_deleted_jobs_total | Total number of deleted jobs
| enqueued_jobs_total | Current number of enqueued jobs
| failed_jobs_total | Total number of failed jobs
| fetched_jobs_total | Current number of fetched jobs
| processing_jobs_total | Current number of processing jobs
| queues_total | Number of queues
| recurring_jobs_total | Current number of recurring jobs
| scheduled_jobs_total | Current number of scheduled jobs
| succeeded_jobs_total | Total number of succeeded jobs


### Flags

* __`listenaddress`:__ The address where the exporter will be listening to (defaults to :8888)
* __`metricspath`:__ The metrics path where the metrics will be available on (defaults to /metrics)
* __`mongoconnection`:__ The mongo connection where the exporter will get the metrics from (defaults to mongo://localhost:27017)
* __`mongodatabase`:__ The mongo database name where the hangfire data will be on (defaults to hangfire)
* __`sqlserverConnection`:__ The SQL Server connection  where the exporter will get the metrics from (defaults to server=localhost;port=1433;Database=Hangfire.Sample;Trusted_Connection=False;User ID=SA;Password=yourStrong(!)Password;")

## Using Docker

You can deploy this exporter using the [samuel/hangfire-exporter](https://hub.docker.com/r/samueltorres/hangfire-exporter/) Docker image.

For example:

```bash
docker pull samueltorres/hangfire-exporter

docker run -d -p 8888:8888 samueltorres/hangfire-exporter
```

[circleci]: https://circleci.com/gh/samueltorres/hangfire-exporter
[hub]: https://hub.docker.com/r/samueltorres/hangfire-exporter/
