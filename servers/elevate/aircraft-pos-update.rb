#!/usr/bin/env ruby

# Global Constants

# areaA = "ALNW - Yakima"
# latA = 46.573515
# longA = -120.552011

# areaB = "ALNW - Seattle"
# latB = 47.543445
# longB = -122.309785

# areaC = "Boeing Field Base"
# latC = 45.52870
# longC = -122.636300

areaA = "Snoqualmie Pass Ski Area"
latA = 47.405333
longA = -121.416667

areaB = "Oregon Health Science University"
latB = 45.499833
longB = -122.685333

areaC = "Boeing Field - Airlift 2"
latC = 47.544000
longC = -122.311333


# current state
curLat = latC
curLong = longC
friendlyName = areaA

msg = "{\"id\":\"6\", \"posLat\":\"#{curLat}\",\"posLong\":\"#{curLong}\",\"posFriendlyName\":\"#{friendlyName}\"}"

# A -> B
# Yakima to Seattle

distA=Math.sqrt(latA**2 + longA**2) 
numDotsA = distA / 0.1

latDiffA = latB - latA
longDiffA = longB - longA

latIncrementA = latDiffA / numDotsA
longIncrementA = longDiffA / numDotsA

puts numDotsA

# B -> C
# Seattle to Portland

distB=Math.sqrt(latB**2 + longB**2) 
numDotsB = distB / 0.1

puts numDotsB

latDiffB = latC - latB
longDiffB = longC - longB

latIncrementB = latDiffB / numDotsB
longIncrementB = longDiffB / numDotsB

# C -> A
# Portland to Yakima

distC=Math.sqrt(latC**2 + longC**2) 
numDotsC = distC / 0.1

puts numDotsC

latDiffC = latA - latC
longDiffC = longA - longC

latIncrementC = latDiffC / numDotsC
longIncrementC = longDiffC / numDotsC

while true
    # A to B
    while curLat.abs < latB.abs && curLong.abs > longB.abs
        puts "Portland to Yakima"
        curLat += latIncrementA
        curLong += longIncrementA
        msg = "{\"id\":\"6\", \"posLat\":\"#{curLat}\",\"posLong\":\"#{curLong}\",\"posFriendlyName\":\"#{friendlyName}\"}"
        puts "[#{curLat}, #{curLong}]"
        puts "gcloud pubsub topics publish test_ac_position_update --message '#{msg}'"
        %x(gcloud pubsub topics publish test_ac_position_update --message '#{msg}')
        sleep 0.5
    end

    curLat = latA
    curLong = longA
    friendlyName = areaA

    # B to C
    while curLat.abs < latC.abs && curLong.abs < longC.abs
        puts "Yakima to Seattle"
        curLat += latIncrementB
        curLong += longIncrementB
        msg = "{\"id\":\"6\", \"posLat\":\"#{curLat}\",\"posLong\":\"#{curLong}\",\"posFriendlyName\":\"#{friendlyName}\"}"
        puts "[#{curLat}, #{curLong}]"
        puts "gcloud pubsub topics publish test_ac_position_update --message '#{msg}'"
        %x(gcloud pubsub topics publish test_ac_position_update --message '#{msg}')
        sleep 0.5
    end

    curLat = latC
    curLong = longC
    friendlyName = areaC

    # C to A
    while curLat.abs > latA.abs && curLong.abs < longA.abs
        puts "Seattle to Portland"
        curLat += latIncrementC
        curLong += longIncrementC
        msg = "{\"id\":\"6\", \"posLat\":\"#{curLat}\",\"posLong\":\"#{curLong}\",\"posFriendlyName\":\"#{friendlyName}\"}"
        puts "[#{curLat}, #{curLong}]"
        puts "gcloud pubsub topics publish test_ac_position_update --message '#{msg}'"
        %x(gcloud pubsub topics publish test_ac_position_update --message '#{msg}')
        sleep 0.5
    end
end