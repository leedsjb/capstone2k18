#!/usr/bin/env ruby

# Global Constants

areaA = "ALNW - Yakima"
latA = 46.573515
longA = -120.552011

areaB = "ALNW - Seattle"
latB = 47.543445
longB = -122.309785

areaC = "Portland"
latC = 45.513499
longC = -122.676489

# current state
curLat = latA
curLong = longA
friendlyName = areaA

msg = "{\"id\":\"2\", \"posLat\":\"#{curLat}\",\"posLong\":\"#{curLong}\",\"posFriendlyName\":\"#{friendlyName}\"}"

# A -> B
# Yakima to Seattle

distA=Math.sqrt(latA**2 + longA**2) 
numDotsA = distA / 0.005

latDiffA = latB - latA
longDiffA = longB - longA

latIncrementA = latDiffA / numDotsA
longIncrementA = longDiffA / numDotsA

puts numDotsA

# B -> C
# Seattle to Portland

distB=Math.sqrt(latB**2 + longB**2) 
numDotsB = distB / 0.005

puts numDotsB

latDiffB = latC - latB
longDiffB = longC - longB

latIncrementB = latDiffB / numDotsB
longIncrementB = longDiffB / numDotsB

# C -> A
# Portland to Yakima

distC=Math.sqrt(latC**2 + longC**2) 
numDotsC = distC / 0.005

puts numDotsC

latDiffC = latA - latC
longDiffC = longA - longC

latIncrementC = latDiffC / numDotsC
longIncrementC = longDiffC / numDotsC

while true
    while curLat < latA && curLong < longA
        curLat += latIncrementA
        curLong += longIncrementA
        msg = "{\"id\":\"2\", \"posLat\":\"#{curLat}\",\"posLong\":\"#{curLong}\",\"posFriendlyName\":\"#{friendlyName}\"}"
        # %x(gcloud pubsub topics publish test_ac_position_update --message '#{msg}')
        puts "[#{curLat}, #{curLong}]"
        sleep 0.5
    end

    curLat = latB
    curLong = longB
    friendlyName = areaB

    while curLat < latB && curLong < longB
        curLat += latIncrementB
        curLong += longIncrementB
        msg = "{\"id\":\"2\", \"posLat\":\"#{curLat}\",\"posLong\":\"#{curLong}\",\"posFriendlyName\":\"#{friendlyName}\"}"
        puts "[#{curLat}, #{curLong}]"
        # %x(gcloud pubsub topics publish test_ac_position_update --message '#{msg}')
        sleep 0.5
    end

    curLat = latC
    curLong = longC
    friendlyName = areaC

    while curLat < latC && curLong < longC
        curLat += latIncrementC
        curLong += longIncrementC
        msg = "{\"id\":\"2\", \"posLat\":\"#{curLat}\",\"posLong\":\"#{curLong}\",\"posFriendlyName\":\"#{friendlyName}\"}"
        # %x(gcloud pubsub topics publish test_ac_position_update --message '#{msg}')
        puts "[#{curLat}, #{curLong}]"
        sleep 0.5
    end
end