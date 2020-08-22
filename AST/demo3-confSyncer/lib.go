package main

import "strings"

var GoModPackageName string

func CheckImportIsHimSelf(importPackageName string) bool {

	return !strings.HasPrefix(importPackageName, GoModPackageName)

}
