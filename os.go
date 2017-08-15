/**
 * Created by 93201 on 2017/7/21.
 */
package utils

import (
	"errors"
	"io/ioutil"
	"os"
	"strconv"
)

func SavePID(p string) error {
	if p == "" {
		return errors.New("invalid file name")
	}
	pid := os.Getpid()
	return ioutil.WriteFile(p, []byte(strconv.Itoa(pid)), 0666)
}
