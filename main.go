package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/peterbourgon/ff"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/samueltorres/hangfire-exporter/mongo"
	"github.com/samueltorres/hangfire-exporter/sqlserver"
	log "github.com/sirupsen/logrus"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	fs := flag.NewFlagSet("hangfire-exporter", flag.ExitOnError)
	var (
		listenAddress       = fs.String("listenaddress", ":8888", "listen address")
		dbType              = fs.String("dbtype", "sqlserver", "the database type (mongo/sqlserver)")
		metricsPath         = fs.String("metricspath", "/metrics", "metrics path")
		mongoConnection     = fs.String("mongoconnection", "mongodb://localhost:27017", "mongo connection")
		mongoDatabase       = fs.String("mongodatabase", "default", "hangfire database")
		sqlserverConnection = fs.String("sqlserverconnection", "server=localhost;port=1433;Database=Hangfire.Sample;Trusted_Connection=False;User ID=SA;Password=yourStrong(!)Password;", "sqlserver connection")
	)
	ff.Parse(fs, os.Args[1:], ff.WithEnvVarNoPrefix())
	logger := log.New()

	var statistics Statistics
	if *dbType == "mongo" {
		s, err := mongo.NewMongoStatistics(*mongoConnection, *mongoDatabase, logger)
		if err != nil {
			logger.Fatal("Could not establish a connection with mongo", err)
		}
		statistics = s
	} else if *dbType == "sqlserver" {
		s, err := sqlserver.NewSqlServerStatistics(*sqlserverConnection, logger)
		if err != nil {
			logger.Fatal("Could not establish a connection with sqlserver", err)
		}
		statistics = s
	} else {
		logger.Fatal("Invalid database type")
	}

	collector := newHangfireCollector(logger, statistics)
	prometheus.MustRegister(collector)

	http.Handle(*metricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
             <head><title>Hangfire Exporter</title></head>
             <body>
             <h1>Hangfire Exporter</h1>
             <p><a href='` + *metricsPath + `'>Metrics</a></p>
             </body>
             </html>`))
	})

	logger.Infof("Providing metrics at %s%s", *listenAddress, *metricsPath)
	logger.Fatal(http.ListenAndServe(*listenAddress, nil))
}
