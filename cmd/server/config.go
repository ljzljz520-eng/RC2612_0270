package main

import (
	"os"
	"strconv"
	"strings"
)

type config struct {
	Address, Database string
	Debug             bool
}

func loadConfig() config {
	c := config{Address: ":8080", Database: "profile.db"}
	if v := strings.TrimSpace(os.Getenv("PROFILE_ADDR")); v != "" {
		c.Address = v
	}
	if v := strings.TrimSpace(os.Getenv("PROFILE_DB")); v != "" {
		c.Database = v
	}
	c.Debug, _ = strconv.ParseBool(os.Getenv("PROFILE_DEBUG"))
	return c
}
func (c config) Valid() bool { return c.Address != "" && c.Database != "" }
