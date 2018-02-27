// Adaptation of https://raw.githubusercontent.com/jxnblk/rebass/master/src/Position.js

import sys from "system-components";

const Sticky = sys(
    {
        position: "sticky"
    },
    "space",
    "color",
    "zIndex",
    "top",
    "right",
    "bottom",
    "left"
);

Sticky.displayName = "Sticky";

export default Sticky;
