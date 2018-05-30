#!/usr/bin/env ruby

# msg = "{\"id\":\"6\", \"posLat\":\"#{curLat}\",\"posLong\":\"#{curLong}\",\"posFriendlyName\":\"#{friendlyName}\"}"
# %x(gcloud pubsub topics publish test_ac_position_update --message '#{msg}')
# puts "[#{curLat}, #{curLong}]"

# waypoints =File.readlines('waypoints/yakima_to_seattle.txt')

# lat, long = waypoints[0].strip.split(' ')



  File.open("waypoints/yakima_to_seattle.txt").each do |line|
    curLat, curLong = line.strip.split(' ')
    friendlyName = "ALNW - Yakima"
    msg = "{\"id\":\"6\", \"posLat\":\"#{curLat}\",\"posLong\":\"#{curLong}\",\"posFriendlyName\":\"#{friendlyName}\"}"
    %x(gcloud pubsub topics publish test_ac_position_update --message '#{msg}')
    sleep 1
  end

  File.open("waypoints/seattle_to_portland.txt").each do |line|
    curLat, curLong = line.strip.split(' ')
    friendlyName = "ALNW - Seattle"
    msg = "{\"id\":\"6\", \"posLat\":\"#{curLat}\",\"posLong\":\"#{curLong}\",\"posFriendlyName\":\"#{friendlyName}\"}"
    %x(gcloud pubsub topics publish test_ac_position_update --message '#{msg}')
    sleep 1
  end

  File.open("waypoints/portland_to_yakima.txt").each do |line|
    curLat, curLong = line.strip.split(' ')
    friendlyName = "Boeing Field Base"
    msg = "{\"id\":\"6\", \"posLat\":\"#{curLat}\",\"posLong\":\"#{curLong}\",\"posFriendlyName\":\"#{friendlyName}\"}"
    %x(gcloud pubsub topics publish test_ac_position_update --message '#{msg}')
    sleep 1
  end

