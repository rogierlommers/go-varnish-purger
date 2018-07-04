package varnish

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
)

// setHeaders adds HTTP headers to a request.
func (i *Invalidator) setHeaders(req *http.Request) {
	if i.beforeRequestFunc != nil {
		i.beforeRequestFunc(req)
	}
}

func (i *Invalidator) readError(r *http.Response) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return fmt.Errorf(string(data))
}

func (i *Invalidator) buildFullPath(targetPath string) string {
	hostPort := fmt.Sprintf("%s:%d", i.hostname, i.port)
	pathToPurge := path.Join(hostPort, targetPath)
	fullPath := fmt.Sprintf("%s://%s", i.protocol, pathToPurge)

	return fullPath
}
