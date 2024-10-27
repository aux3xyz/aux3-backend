package main

import (
    "os"
)

func init() {
    // TODO load env
}

func GetConfig(name string) string {
    // TODO access env var from map?
    return os.Getenv(name)
}
