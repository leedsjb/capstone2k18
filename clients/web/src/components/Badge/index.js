// Adaptation of https://github.com/jxnblk/rebass/blob/master/src/Badge.js

import sys from "system-components";

const Badge = sys(
    {
        fontSize: 0,
        px: 2,
        py: 1,
        mx: 0,
        fontWeight: "bold",
        borderRadius: 2
    },
    {
        WebkitFontSmoothing: "antialiased",
        display: "inline-block",
        verticalAlign: "middle"
    },
    "color",
    "borders"
);

Badge.displayName = "Badge";

export default Badge;
