#!/bin/bash

while true; do
	gcloud pubsub topics publish test_mission_create --message "{\"type\":\"FETCH_AIRCRAFTDETAIL_SUCCESS\",\"payload\":{\"status\":\"on a mission\",\"callsign\":\"AL3\",\"mission\":{\"missionDetails\":\"blahblahblah\"}}}"
	echo $i
	sleep 1s
done
