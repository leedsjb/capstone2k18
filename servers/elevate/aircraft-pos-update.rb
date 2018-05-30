#!/usr/bin/env ruby

require "Math"

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

msg = "{\"id\":\"2\", \"posLat\":\"#{curLat}\",\"posLong\":\"#{curLong}\",\"posFriendlyName\":\"#{friendlyName}\"}"

# A -> B
# Yakima to Seattle

distA=Math.sqrt(latA**2 + longA**2) 
numDotsA = distA / 0.005

latDiffA = latB - latA
longDiffA = longB - longA

latIncrementA = latDiffA / numDotsA
longIncrementA = longDiffA / numDotsA

# B -> C
# Seattle to Portland

distB=Math.sqrt(latB**2 + longB**2) 
numDotsB = distB / 0.005

latDiffB = latC - latB
longDiffB = longC - longB

latIncrementB = latDiffB / numDotsB
longIncrementB = longDiffB / numDotsB

# C -> A
# Portland to Yakima

distC=Math.sqrt(latC**2 + longC**2) 
numDotsC = distC / 0.0025

latDiffC = latA - latC
longDiffC = longA - longC

latIncrementC = latDiffC / numDotsC
longIncrementC = longDiffC / numDotsC

while true
    while curLat < latA && curLong < longA
        curLat += latIncrement
        curLong += longIncrement
        %x(gcloud pubsub topics publish test_ac_position_update --message '#{msg}')
        sleep 0.5
    end

    while curLat < latB && curLong < longB
        curLat += latIncrement
        curLong += longIncrement
        %x(gcloud pubsub topics publish test_ac_position_update --message '#{msg}')
        sleep 0.5
    end

    while curLat < latC && curLong < longC
        curLat += latIncrement
        curLong += longIncrement
        %x(gcloud pubsub topics publish test_ac_position_update --message '#{msg}')
        sleep 0.5
    end
end