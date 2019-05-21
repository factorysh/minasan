package metrics

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

var (
	MailReceivedCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "mail_received",
		Help: "Amount of mails receiveds",
	})
	MailSentCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "mail_sent",
		Help: "Amount of mails sents",
	})
	WrongProjectCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "wrong_project",
		Help: "Project that doesn't exist",
	})
)

func ListenAndServe(listen string) error {
	http.Handle("/metrics", promhttp.Handler())
	log.WithFields(
		log.Fields{
			"url": fmt.Sprintf("http://%s/metrics", listen),
		},
	).Info("metrics.ListenAndServe")
	return http.ListenAndServe(listen, nil)
}
