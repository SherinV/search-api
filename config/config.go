/*
IBM Confidential
OCO Source Materials
(C) Copyright IBM Corporation 2019 All Rights Reserved
The source code for this program is not published or otherwise divested of its trade secrets,
irrespective of what has been deposited with the U.S. Copyright Office.

Copyright (c) 2020 Red Hat, Inc.
*/
// Copyright Contributors to the Open Cluster Management project

package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/golang/glog"
)

const (
	API_SERVER_URL           = "https://kubernetes.default.svc"
	DEFAULT_CONTEXT_PATH     = "/searchapi"
	DEFAULT_QUERY_LIMIT      = 10000
	DEFAULT_QUERY_LOOP_LIMIT = 5000
	HTTP_PORT                = 4010
	RBAC_POLL_INTERVAL       = 60000
	RBAC_INACTIVITY_TIMEOUT  = 600000
	SERVICEACCT_TOKEN        = ""
)

// Define a config type to hold our config properties.
type Config struct {
	API_SERVER_URL          string // address for API_SERVER
	ContextPath             string
	defaultQueryLimit       int // number of queries handled at a time
	defaultQueryLoopLimit   int
	HttpPort                int
	RBAC_POLL_INTERVAL      int
	RBAC_INACTIVITY_TIMEOUT int
	SERVICEACCT_TOKEN       string
}

var Cfg = Config{}

func init() {
	// If environment variables are set, use those values constants
	// Simply put, the order of preference is env -> default constants (from left to right)
	setDefault(&Cfg.API_SERVER_URL, "API_SERVER_URL", API_SERVER_URL)
	setDefault(&Cfg.ContextPath, "CONTEXT_PATH", DEFAULT_CONTEXT_PATH)
	setDefaultInt(&Cfg.defaultQueryLimit, "QUERY_LIMIT", DEFAULT_QUERY_LIMIT)
	setDefaultInt(&Cfg.defaultQueryLoopLimit, "QUERY_LOOP_LIMIT", DEFAULT_QUERY_LOOP_LIMIT)
	setDefaultInt(&Cfg.HttpPort, "HTTP_PORT", HTTP_PORT)
	setDefaultInt(&Cfg.RBAC_POLL_INTERVAL, "RBAC_POLL_INTERVAL", RBAC_POLL_INTERVAL)
	setDefaultInt(&Cfg.RBAC_INACTIVITY_TIMEOUT, "RBAC_INACTIVITY_TIMEOUT", RBAC_INACTIVITY_TIMEOUT)
	setDefault(&Cfg.SERVICEACCT_TOKEN, "SERVICEACCT_TOKEN", SERVICEACCT_TOKEN)
}

func setDefault(field *string, env, defaultVal string) {
	if val := os.Getenv(env); val != "" {
		if env == "REDIS_PASSWORD" {
			glog.Infof("Using %s from environment", env)
		} else {
			glog.Infof("Using %s from environment: %s", env, val)
		}
		*field = val
	} else if *field == "" && defaultVal != "" {
		// Skip logging when running tests to reduce confusing output.
		if !strings.HasSuffix(os.Args[0], ".test") {
			glog.Infof("%s not set, using default value: %s", env, defaultVal)
		}
		*field = defaultVal
	}
}

func setDefaultInt(field *int, env string, defaultVal int) {
	if val := os.Getenv(env); val != "" {
		glog.Infof("Using %s from environment: %s", env, val)
		var err error
		*field, err = strconv.Atoi(val)
		if err != nil {
			glog.Error("Error parsing env [", env, "].  Expected an integer.  Original error: ", err)
		}
	} else if *field == 0 && defaultVal != 0 {
		// Skip logging when running tests to reduce confusing output.
		if !strings.HasSuffix(os.Args[0], ".test") {
			glog.Infof("No %s from file or environment, using default value: %d", env, defaultVal)
		}
		*field = defaultVal
	}
}
