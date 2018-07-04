package varnish

import (
	"fmt"
	"net/http"
)

// Invalidatepath sends purge request to the path and return the status code
func (i Invalidator) InvalidatePath(targetPath string) error {
	req, err := http.NewRequest("PURGE", i.buildFullPath(targetPath), nil)
	if err != nil {
		return err
	}

	i.setHeaders(req)

	resp, err := i.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf(resp.Status)
	}

	return nil
}
