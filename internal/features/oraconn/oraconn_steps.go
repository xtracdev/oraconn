package oraconn

import (
	. "github.com/lsegal/gucumber"
	log "github.com/Sirupsen/logrus"
	"github.com/xtraclabs/oraconn"
	"github.com/stretchr/testify/assert"
)

const connectStr = "system/oracle@//localhost:1521/xe.oracle.docker"
const bogusConnectStr = "system/oracle@//localhost:15121/xe.oracle.docker"

func init() {
	var db *oraconn.OracleDB
	var noConnectError error

	Given(`^a running oracle instance$`, func() {
		log.Infof("Oracle instance available via %s assumed", connectStr)
	})

	When(`^provide a connection string for the running instance$`, func() {
		//The connectStr constant
	})

	Then(`^a connection is returned$`, func() {
		var err error
		db, err = oraconn.OpenAndConnect(connectStr, 10)
		assert.Nil(T, err)
	})

	Given(`^a connection string with no listener$`, func() {
		log.Infof("No oracle instance available via %s assumed", connectStr)
	})

	When(`^I connect to no listener$`, func() {
		db, noConnectError = oraconn.OpenAndConnect(bogusConnectStr, 3)
	})

	Then(`^an error is returned$`, func() {
		assert.NotNil(T, noConnectError)
	})

}

