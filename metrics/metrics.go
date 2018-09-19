package metrics

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

var (
	MailReceivedCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "mail_received",
		Help: "Amount of mails receiveds",
	})
	MailSentCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "mail_sent",
		Help: "Amount of mails sents",
	})
	WrongProjectCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "Wrong project",
		Help: "Project that doesn't exist",
	})
)

func init() {
	prometheus.MustRegister(MailReceivedCounter)
}

func ListenAndServe(listen string) error {
	http.Handle("/metrics", promhttp.Handler())
	log.WithFields(
		log.Fields{
			"url": fmt.Sprintf("http://%s/metrics", listen),
		},
	).Info("metrics.ListenAndServe")
	return http.ListenAndServe(listen, nil)
}
