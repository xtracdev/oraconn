package oraconn

import (
	. "github.com/lsegal/gucumber"
	log "github.com/Sirupsen/logrus"
	"github.com/xtraclabs/oraconn"
	"github.com/stretchr/testify/assert"
	"time"
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

}

