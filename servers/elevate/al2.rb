#!/usr/bin/env ruby

# msg = "{\"id\":\"6\", \"posLat\":\"#{curLat}\",\"posLong\":\"#{curLong}\",\"posFriendlyName\":\"#{friendlyName}\"}"
# %x(gcloud pubsub topics publish test_ac_position_update --message '#{msg}')
# puts "[#{curLat}, #{curLong}]"

# waypoints =File.readlines('waypoints/yakima_to_seattle.txt')

# lat, long = waypoints[0].strip.split(' ')


  File.open("waypoints/snoqualmie_to_oregon.txt").each do |line|
    curLat, curLong = line.strip.split(' ')
    friendlyName = "Snoqualmie Pass Ski Area"
    msg = "{\"id\":\"1\", \"posLat\":\"#{curLat}\",\"posLong\":\"#{curLong}\",\"posFriendlyName\":\"#{friendlyName}\"}"
    puts "[#{curLat}, #{curLong}] #{friendlyName}"
    %x(gcloud pubsub topics publish test_ac_position_update --message '#{msg}')
    sleep 1
  end

  File.open("waypoints/oregon_to_boeing.txt").each do |line|
    curLat, curLong = line.strip.split(' ')
    friendlyName = "Oregon Health Science University"
    msg = "{\"id\":\"1\", \"posLat\":\"#{curLat}\",\"posLong\":\"#{curLong}\",\"posFriendlyName\":\"#{friendlyName}\"}"
    puts "[#{curLat}, #{curLong}] #{friendlyName}"
    %x(gcloud pubsub topics publish test_ac_position_update --message '#{msg}')
    sleep 1
  end

  File.open("waypoints/boeing_to_snoqualmie.txt").each do |line|
    curLat, curLong = line.strip.split(' ')
    friendlyName = "Boeing Field - Airlift 2"
    msg = "{\"id\":\"1\", \"posLat\":\"#{curLat}\",\"posLong\":\"#{curLong}\",\"posFriendlyName\":\"#{friendlyName}\"}"
    puts "[#{curLat}, #{curLong}] #{friendlyName}"
    %x(gcloud pubsub topics publish test_ac_position_update --message '#{msg}')
    sleep 1
  end

