// Adaptation of https://github.com/jxnblk/rebass/blob/master/src/Position.js

import sys from "system-components";

const Fixed = sys(
    {
        position: "fixed"
    },
    "space",
    "color",
    "zIndex",
    "top",
    "right",
    "bottom",
    "left"
);

Fixed.displayName = "Fixed";

export default Fixed;
