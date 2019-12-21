package route

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pulsar-beam/src/pulsardriver"
	"github.com/pulsar-beam/src/util"
)

func init() {
	//where to initialize all DB connection
}

// StatusPage replies with basic status code
func StatusPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	return
}

// ReceiveHandler - the message receiver handler
func ReceiveHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader((http.StatusInternalServerError))
		return
	}

	vars := mux.Vars(r)
	tenant := vars["tenant"]
	token, topicFN, pulsarURL, err2 := util.ReceiverHeader(&r.Header)
	if err2 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	log.Printf("tenant %s token %s topicFN %s puslarURL %s", tenant, token, topicFN, pulsarURL)

	err = pulsardriver.SendToPulsar(pulsarURL, token, topicFN, b)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}

// UpdateTopicHandler creates or updates a topic
func UpdateTopicHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	return
}

// DeleteTopicHandler deletes a topic
func DeleteTopicHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	return
}