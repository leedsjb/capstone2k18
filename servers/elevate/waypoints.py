
from geographiclib.geodesic import Geodesic
import math
geod = Geodesic.WGS84


# lines between locations
snoqualmie_to_oregon = geod.InverseLine(47.405333, -121.416667, 45.499833, -122.685333)
oregon_to_boeing = geod.InverseLine(45.499833, -122.685333, 45.513499, -122.676489)
boeing_to_snoqualmie = geod.InverseLine(47.544000, -122.311333, 47.405333, -121.416667)
 
# I should have fixed this to loop, but there were only three so I did it manually
l = portland_to_yakima
 
# ds is in meters (this makes it effectively meters/tick, so here it
# is meters/second if you update to a new waypoint each second)
ds = 223; n = int(math.ceil(l.s13 / ds))
print n
with open("portland_to_yakima.txt", "a") as myfile:
    for i in range(n + 1):
        if i == 0:
            print "distance latitude longitude"
        s = min(ds * i, l.s13)
        g = l.Position(s, Geodesic.STANDARD | Geodesic.LONG_UNROLL)
        myfile.write("{:.5f} {:.5f}".format(g['lat2'], g['lon2']) + "\n")
        print "{:.5f} {:.5f}".format(g['lat2'], g['lon2'])