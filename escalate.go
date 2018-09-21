package main

import (
	"flag"
	"log"

	"github.com/ClinicalSystemsEngineering/nats/natssub"
	"github.com/nats-io/go-nats-streaming"
)

func main() {
	var clusterID string
	var clientID string
	var showTime bool
	var startSeq uint64
	var startDelta string
	var deliverAll bool
	var deliverLast bool
	var durable string
	var qgroup string
	var unsubscribe bool
	var URL string
	var subj string
	var cancelsubj string
	var readmsgs = make(chan string, 10000) //message processing channel for nats reading

	flag.StringVar(&URL, "s", stan.DefaultNatsURL, "The nats server URLs (separated by comma)")
	flag.StringVar(&URL, "server", stan.DefaultNatsURL, "The nats server URLs (separated by comma)")
	flag.StringVar(&clusterID, "c", "test-cluster", "The NATS Streaming cluster ID")
	flag.StringVar(&clusterID, "cluster", "test-cluster", "The NATS Streaming cluster ID")
	flag.StringVar(&clientID, "id", "", "The NATS Streaming client ID to connect with")
	flag.StringVar(&clientID, "clientid", "", "The NATS Streaming client ID to connect with")
	flag.BoolVar(&showTime, "t", false, "Display timestamps")
	// Subscription options
	flag.Uint64Var(&startSeq, "seq", 0, "Start at sequence no.")
	flag.BoolVar(&deliverAll, "all", false, "Deliver all")
	flag.BoolVar(&deliverLast, "last", false, "Start with last value")
	flag.StringVar(&startDelta, "since", "", "Deliver messages since specified time offset")
	flag.StringVar(&durable, "durable", "", "Durable subscriber name")
	flag.StringVar(&qgroup, "qgroup", "", "Queue group name")
	flag.BoolVar(&unsubscribe, "unsubscribe", false, "Unsubscribe the durable on exit")
	flag.StringVar(&subj, "subj", "Hospital.System", "Name of subj to read ingest messages from")
	flag.StringVar(&cancelsubj, "cancelsubj", "Hospital.System.Cancel", "Name of subject to read cancel messages from")

	log.SetFlags(0)
	//flag.Usage = usage
	flag.Parse()

	/* args := flag.Args()

	if clientID == "" {
		log.Printf("Error: A unique client ID must be specified.")
		usage()
	}
	if len(args) < 1 {
		log.Printf("Error: A subject must be specified.")
		usage()
	} */

	go natssub.Subber(clusterID, clientID, showTime, startSeq, startDelta, deliverAll, deliverLast, durable, qgroup, unsubscribe, URL, subj, readmsgs)

	//implement channel listener and exec spawner
}
