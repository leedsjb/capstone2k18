package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	// go sql driver
	_ "github.com/go-sql-driver/mysql"

	// "github.com/leedsjb/capstone2k18/servers/gateway/handlers"

	"github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
)

//main is the main entry point for the server
func main() {

	// get the present working directory
	pwd := os.Getenv("PWD")

	// Read the ADDR environment variable to get the address
	// the server should listen on.
	addr := os.Getenv("ADDR")
	// If empty, default to ":443" for https
	if len(addr) == 0 {
		addr = ":443"
	}

	//TLSKEY and TLSCERT: paths to TLS key and cert
	// TLS Key and Cert facilitate secure communication between apiserver and the Google Cloud
	// Platform HTTPS Load Balancer. Certificate can be self-signed. Keys are managed by GCP KMS.
	// See https://cloud.google.com/kms/
	tlsKeyPath := os.Getenv("TLSKEY")
	tlsCertPath := os.Getenv("TLSCERT")
	if tlsKeyPath == "" || tlsCertPath == "" {
		fmt.Println("Attempting to use default TLSKEY and TLSCERT paths")
		tlsKeyPath = pwd + "/tls/privkey.pem"
		tlsCertPath = pwd + "/tls/fullchain.pem"
		fmt.Println(tlsKeyPath)
		fmt.Println(tlsCertPath)
	}

	sqlInstance := os.Getenv("SQLINSTANCE")
	sqlUser := os.Getenv("SQLUSER")
	sqlPass := os.Getenv("SQLPASS")
	sqlDbName := os.Getenv("SQLDBNAME")
	// sqlTblName := os.Getenv("SQLTABLENAME")
	first25Missions := os.Getenv("SQLMISSIONS")

	cfg := mysql.Cfg(sqlInstance, sqlUser, sqlPass)
	cfg.DBName = sqlDbName
	db, err := mysql.DialCfg(cfg)
	if err != nil {
		fmt.Printf("Error dialing MySQL: %v", err)
		os.Exit(1)
	}
	defer db.Close()

	rows, err := db.Query(first25Missions)
	// rows, err := db.Query("SELECT * FROM " + sqlTblName)
	if err != nil {
		fmt.Printf("Error querying MySQL: %v", err)
		os.Exit(1)
	}

	// TODO: DO SOMETHING WITH ROWS
	i := 0
	for rows.Next() { 
		var mission_id string
		var aircraft_id string
		var agency_id string
		var mission_date string
		var aircraft_id2 string
		var aircraft_callsign string
		var model_id string
		var aircraft_lat string
		var aircraft_long string
		var agency_id2 string
		var agency_name string
		var agency_area_code string
		var agency_phone string
		var address_id string
		err := rows.Scan(
			&mission_id,
			&aircraft_id,
			&agency_id,
			&mission_date,
			&aircraft_id2,
			&aircraft_callsign,
			&model_id,
			&aircraft_lat,
			&aircraft_long,
			&agency_id2,
			&agency_name,
			&agency_area_code,
			&agency_phone,
			&address_id,
		)
		if err != nil {
			fmt.Printf("Error parsing MySQL rows: %v", err)
			os.Exit(1)
		}
		fmt.Printf(
			"========================================================\nFLIGHT %d\naircraft_lat: %s\naircraft_long: %s\nmodel_title: %s\nmodel_desc: %s\naircraft_category: %s\naircraft_callsign: %s\nmission_date: %s\nagency_name: %s\nagency_area_code: %s\nagency_phone:%s\naddress_street: %s\naddress_city: %s\naddress_state: %s\naddress_zip: %s\n",
			i, 
			mission_id,
			aircraft_id,
			agency_id,
			mission_date,
			aircraft_id2,
			aircraft_callsign,
			model_id,
			aircraft_lat,
			aircraft_long,
			agency_id2,
			agency_name,
			agency_area_code,
			agency_phone,
			address_id,
		)
		i++
	}

	// Create a new mux for the web server.
	mux := http.NewServeMux()

	//Wrap this new mux with CORS middleware handler and add that
	//to the main server mux.
	// wrappedMux := handlers.NewCORSHandler(mux)

	// Tell the mux to call your handlers

	// Start a web server listening on the address you read from
	// the environment variable, using the mux you created as
	// the root handler. Use log.Fatal() to report any errors
	// that occur when trying to start the web server.
	log.Printf("server is listening at %s...", addr)
	log.Fatal(http.ListenAndServeTLS(addr, tlsCertPath, tlsKeyPath, mux))
}