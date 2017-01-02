package main

type Vel struct {
	Raw                 []byte
	Json                *JsonVel
	VirtualHostLocation *string
}

type JsonVel struct {
	About struct {
		Docs    string
		Credits struct {
			Links string
			Email string
		}
		Notes   []string
		License struct {
			Name string
			Link string
		}
	}
	Last_updated struct {
		Timestamp    string
		Human        string
		Current_hash string
	}
	Data []Data
}

type Data struct {
	Official_vel_link      string
	Title                  string
	Type                   string
	Published_date         string
	Component_name         string
	Com_whatever           string
	Version_effected       string
	Version_effected_trend string
	Developer_website      string
	Extension_webpage      string
	Developers_email       string
	Response_url           string
	Github_repo            string
}
