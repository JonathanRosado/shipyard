package manager

import (
//	"bytes"
//	"encoding/json"
//	"fmt"
//	"net/http"
//	"net/http/httptest"
	"testing"

//	"github.com/shipyard/shipyard/controller/mock_test"
	"github.com/stretchr/testify/assert"
	"github.com/jrosadohp/shipyard/controller/manager"
	"github.com/shipyard/shipyard/auth/builtin"
	"github.com/shipyard/shipyard/utils"
	"github.com/shipyard/shipyard/controller/api"
//	"github.com/shipyard/shipyard"
//	"bytes"
)

func newManager() manager.Manager {
	rethinkdbAddr := "rethinkdb:28015"
	rethinkdbDatabase := "shipyard"
	rethinkdbAuthKey := ""
	disableUsageInfo := ""

	dockerUrl := "tcp://127.0.0.1:2375"
	tlsCaCert := ""
	tlsCert := ""
	tlsKey := ""
	allowInsecure := false

	client, err := utils.GetClient(dockerUrl, tlsCaCert, tlsCert, tlsKey, allowInsecure)
	if err != nil {
		panic(err)
	}

	// default to builtin auth
	authenticator := builtin.NewAuthenticator("defaultshipyard")

	controllerManager, err := manager.NewManager(rethinkdbAddr, rethinkdbDatabase, rethinkdbAuthKey, client, disableUsageInfo, authenticator)
	if err != nil {
		panic(err)
	}

	return controllerManager
}

func newApi() api.Api {
	apiConfig := api.ApiConfig{
		ListenAddr:         listenAddr,
		Manager:            controllerManager,
		AuthWhiteListCIDRs: authWhitelist,
		EnableCORS:         enableCors,
		AllowInsecure:      allowInsecure,
		TLSCACertPath:      shipyardTlsCACert,
		TLSCertPath:        shipyardTlsCert,
		TLSKeyPath:         shipyardTlsKey,
	}

	shipyardApi, err := api.NewApi(apiConfig)
	if err != nil {
		log.Fatal(err)
	}

	if err := shipyardApi.Run(); err != nil {
		log.Fatal(err)
	}
}

func TestApiPostAccounts(t *testing.T) {
    manager := newManager()

	assert.Equal(t, false, true, "expected response code 204")
}