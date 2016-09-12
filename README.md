## Oraconn - Oracle Connection Utilities

This package provide a simple mechanism to extend the Oracle sql.DB
implementation with some additional capabilities:

* The ability to retry initial connection attempts using a simple backoff
 mechanism. This is useful in scenarios such as starting containerized
 applications without wanting to worry about start up order.
 * Detecting certain classes of connection related errors.
 * Reconnecting to the database using retrys and backoff.
 
## Usage
 
 * Use the `OpenAndConnect` method to instantiate a `*sql.DB` instance.
 * Use `IsConnectionError` to determine if the error indicates the connection
 to the database is lost or hopeless, in which `Reconnect` should be used
 to reconnect to the database.
 
## Implementation Notes
 
This package uses the `https://github.com/mattn/go-oci8` package for 
working with Oracle.