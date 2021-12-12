package controllers

import (
	"encoding/json"
	"net/http"
)

func (c *Controller) getBlockChainHandler(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	return json.Marshal(c.blockchain)

}

func (c *Controller) postMineBlockHandler(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	newBlock := c.blockchain.AddBlock("test transaction")
	return json.Marshal(newBlock)
}
