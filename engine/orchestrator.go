package engine

import (
	"fmt"
	"minotor/config"
	"minotor/utils"
	"net/http"
)

func HarvestComsosWallet() {

	url := fmt.Sprintf("%s:%d/cosmos/wrapper/", config.Cfg.APIAdress, config.Cfg.APIPort)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	defer resp.Body.Close()
}

func FluxNodeRentability() {
	url := fmt.Sprintf("%s:%d/flux/calcul_nodes_rentability", config.Cfg.APIAdress, config.Cfg.APIPort)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	defer resp.Body.Close()

	fmt.Println(resp)
}

func HealthCheck() {
	url := fmt.Sprintf("%s:%d/health", config.Cfg.APIAdress, config.Cfg.APIPort)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	defer resp.Body.Close()

	fmt.Println(resp)
}

func GetStreamR() {
	url := fmt.Sprintf("%s:%d/streamr/status/%s", config.Cfg.APIAdress, config.Cfg.APIPort, config.Cfg.StreamRAddr)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	defer resp.Body.Close()

	fmt.Println(resp)
}

func WebhookOsmoPullRequest() {
	url := fmt.Sprintf("%s:%d/github/osmosis-labs/osmosis-frontend", config.Cfg.APIAdress, config.Cfg.APIPort)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	defer resp.Body.Close()

	fmt.Println(resp)
}
