package oraconn

import (
	. "github.com/gucumber/gucumber"
	log "github.com/Sirupsen/logrus"
	"github.com/xtracdev/oraconn"
	"github.com/stretchr/testify/assert"
	"time"
	"errors"
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

	And(`^I can select the system timestamp from dual$`, func() {

		rows, err := db.Query("select systimestamp from dual")
		if assert.Nil(T, err) {
			defer rows.Close()

			for rows.Next() {
				var ts time.Time
				rows.Scan(&ts)
				log.Infof("systimestamp from dual is %s", ts.Format(time.RFC3339))
			}

			assert.Nil(T, rows.Err())
		}
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

	Given(`^a loss of database connectivity$`, func() {
		var err error
		db, err = oraconn.OpenAndConnect(connectStr, 10)
		if assert.Nil(T, err) {
			err = db.Close()
			assert.Nil(T, err)
		}
	})

	When(`^I detect I've lost connectivity$`, func() {
		//If I close the database I get a sql: database is closed error. I don't want to bake that
		//message into the error function as I want to detect loss of connection, not sloppy programming.
		//For this test I'll simulate the errors I've seen while running sample code with bring the
		//db up and down by hand.

		//TODO: see if we can automate this using the Docker API to bring Oracle up and down
		assert.True(T, oraconn.IsConnectionError(errors.New("ORA-03114: Not Connected to Oracle")), "Expected a connection error")
	})

	Then(`^I can reconnect$`, func() {
		err := db.Reconnect(3)
		assert.Nil(T, err)
	})

	And(`^I can select data after reconnecting$`, func() {
		rows, err := db.Query("select systimestamp from dual")
		if assert.Nil(T, err) {
			defer rows.Close()

			for rows.Next() {
				var ts time.Time
				rows.Scan(&ts)
				log.Infof("systimestamp from dual is %s", ts.Format(time.RFC3339))
			}

			assert.Nil(T, rows.Err())
		}
	})
}

