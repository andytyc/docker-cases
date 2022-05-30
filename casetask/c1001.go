package casetask

import (
	"fmt"
	"os"
)

func C1001OSENV() {
	DARWIN_ETCD_ENDPOINTS := os.Getenv("DARWIN_ETCD_ENDPOINTS")
	DARWIN_ETCD_AUTH := os.Getenv("DARWIN_ETCD_AUTH")
	fmt.Println("DARWIN_ETCD_ENDPOINTS:", DARWIN_ETCD_ENDPOINTS, "DARWIN_ETCD_AUTH:", DARWIN_ETCD_AUTH)

	ALG_DIR := os.Getenv("ALG_DIR")
	fmt.Println("ALG_DIR:", ALG_DIR)

	PATH := os.Getenv("PATH")
	fmt.Println("PATH:", PATH)
}
