package main

import (
	"log"
	"strings"
)

func (v *Vel) check(name string, version string) (bool, string, error) {
	for i := range v.Json.Data {
		if v.Json.Data[i].Com_whatever != "" {
			if strings.EqualFold(v.Json.Data[i].Com_whatever, name) {
				// 9999 means all versions are effected
				if v.Json.Data[i].Version_effected == "" {
					continue
				}
				// Check for multiple versions in the feed
				if strings.Contains(v.Json.Data[i].Version_effected, ",") {
					split := strings.Split(v.Json.Data[i].Version_effected, ",")
					for _, foo := range split {
						lessThan, err := lessThan(version, strings.TrimSpace(foo))
						if err != nil {
							log.Printf("Error was: %s", err.Error())
							continue
						}
						if lessThan {
							return true, v.Json.Data[i].Version_effected + " was published " + v.Json.Data[i].Published_date + " " + v.Json.Data[i].Official_vel_link, nil
						}
					}
				} else {
					// This is a single version in the feed
					lessThan2, err := lessThan(version, v.Json.Data[i].Version_effected)
					if err != nil {
						log.Printf("Error was: %s", err.Error())
					}
					if lessThan2 {
						return true, v.Json.Data[i].Version_effected + " published " + v.Json.Data[i].Published_date + " " + v.Json.Data[i].Official_vel_link, nil
					}
				}
			}
		}
	}
	return false, "", nil
}
