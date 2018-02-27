// Adaptation of https://raw.githubusercontent.com/jxnblk/rebass/master/src/Position.js

import sys from "system-components";

export const Position = sys(
    {
        position: "relative"
    },
    "space",
    "color",
    "zIndex",
    "top",
    "right",
    "bottom",
    "left"
);

Relative.displayName = "Relative";

export default Relative;
