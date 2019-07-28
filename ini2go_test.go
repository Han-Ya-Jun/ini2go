package ini2go

import (
	"testing"
)

/*
* @Author:hanyajun
* @Date:2019/7/27 23:13
* @Name:app
* @Function:
 */

func TestIni2Go(t *testing.T) {
	_ = Ini2Go("example/app.ini", "example", "app.go", "cmd/", false, []string{})
}
