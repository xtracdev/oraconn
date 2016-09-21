build:
	export PKG_CONFIG_PATH=$GOPATH/src/github.com/xtraclabs/oraconn/pkgconfig/
	go get github.com/rjeczalik/pkgconfig/cmd/pkg-config
	go get -u github.com/mattn/go-oci8
	go get github.com/Sirupsen/logrus
	go build
