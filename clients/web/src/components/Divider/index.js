// Adaptation of https://raw.githubusercontent.com/jxnblk/rebass/master/src/Divider.js

import sys from "system-components";

const Divider = sys(
    {
        is: "hr",
        mx: 0,
        my: 0,
        border: 0,
        borderBottom: 1,
        borderColor: "wireframe"
    },
    "space",
    "color"
);

Divider.displayName = "Divider";

export default Divider;
