package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Sandy247/blockchain-db/internal/blockchain"
)

type Controller struct {
	blockchain *blockchain.BlockChain
}

type RootHandler func(http.ResponseWriter, *http.Request) ([]byte, error)

func NewController() *Controller {
	return &Controller{blockchain: blockchain.NewBlockChain()}
}

func (c *Controller) welcomeHandler(w http.ResponseWriter, req *http.Request) ([]byte, error) {
	return json.Marshal(map[string]string{"message": "Welcome to the blockchain"})
}

func (c *Controller) healthCheckHandler(w http.ResponseWriter, req *http.Request) ([]byte, error) {
	return json.Marshal(map[string]string{"message": "Healthy"})
}

func (c *Controller) RequestsHandler(w http.ResponseWriter, req *http.Request) ([]byte, error) {
	w.Header().Set("Content-Type", "application/json")
	p := req.URL.Path
	m := req.Method
	var res []byte
	var err error
	switch {
	case p == "/":
		res, err = c.welcomeHandler(w, req)
	case p == "/health":
		res, err = c.healthCheckHandler(w, req)
	case p == "/blockchain" && m == http.MethodGet:
		res, err = c.getBlockChainHandler(w, req)
	case p == "/blockchain/mine" && m == http.MethodGet:
		res, err = c.postMineBlockHandler(w, req)
	default:
		res, err = nil, errors.New("not found")
	}
	return res, err
}

func (fn RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res, err := fn(w, r)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		w.Write(res)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}
