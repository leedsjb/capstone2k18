// Adaptation of https://github.com/jxnblk/rebass/blob/master/src/Position.js

import sys from "system-components";

const Absolute = sys(
    {
        position: "absolute"
    },
    "space",
    "color",
    "zIndex",
    "top",
    "right",
    "bottom",
    "left"
);

Absolute.displayName = "Absolute";

export default Absolute;
