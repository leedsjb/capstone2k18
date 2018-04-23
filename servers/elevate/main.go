package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	// go sql driver
	_ "github.com/go-sql-driver/mysql"

	"github.com/leedsjb/capstone2k18/servers/gateway/handlers"

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
	}

	//DBADDR: the address of your database server
	dbAddr := os.Getenv("DBADDR")
	if dbAddr == "" {
		fmt.Println("Please provide the address of your database server")
		os.Exit(1)
	}

	sqlInstance := os.Getenv("SQLINSTANCE")
	sqlUser := os.Getenv("SQLUSER")
	sqlPass := os.Getenv("SQLPASS")
	sqlDbName := os.Getenv("SQLDBNAME")
	sqlTblName := os.Getenv("SQLTABLENAME")

	cfg := mysql.Cfg(sqlInstance, sqlUser, sqlPass)
	cfg.DBName = sqlDbName
	db, err := mysql.DialCfg(cfg)
	if err != nil {
		fmt.Printf("Error dialing MySQL: %v", err)
		os.Exit(1)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM " + sqlTblName)
	if err != nil {
		fmt.Printf("Error querying MySQL: %v", err)
		os.Exit(1)
	}

	// TODO: DO SOMETHING WITH ROWS
	i := 0
	for rows.Next() {
		i = i + 1
		fmt.Println(i)
		// var aircraft_lat int
		// var aircraft_long string
		// var model_title string
		// var model_desc string
		// var aircraft_category string
		// var aircraft_callsign string
		// var mission_date string
		// var agency_name string
		// var agency_area_code string
		// var agency_phone string
		// var adderss_street string
		// var address_city string
		// var address_state string
		// var address_zip string
		// err := rows.Scan(
		// 	&aircraft_lat,
		// 	&aircraft_long,
		// 	&model_title,
		// 	&model_desc,
		// 	&aircraft_category,
		// 	&aircraft_callsign,
		// 	&mission_date,
		// 	&agency_name,
		// 	&agency_area_code,
		// 	&address_state,
		// 	&address_zip,
		// 	&address_city,
		// 	&address_state,
		// 	&address_zip
		// )
		// if err != nil {
		// 	fmt.Printf("Error parsing MySQL rows: %v", err)
		// 	os.Exit(1)
		// }
		// fmt.Println(uid)
		// fmt.Println(username)
		// fmt.Println(department)
		// fmt.Println(created)
	}

	// Create a new mux for the web server.
	mux := http.NewServeMux()

	//Wrap this new mux with CORS middleware handler and add that
	//to the main server mux.
	wrappedMux := handlers.NewCORSHandler(mux)

	// Tell the mux to call your handlers

	// Start a web server listening on the address you read from
	// the environment variable, using the mux you created as
	// the root handler. Use log.Fatal() to report any errors
	// that occur when trying to start the web server.
	log.Printf("server is listening at %s...", addr)
	log.Fatal(http.ListenAndServeTLS(addr, tlsCertPath, tlsKeyPath, wrappedMux))
}