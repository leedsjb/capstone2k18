// Adaptation of https://raw.githubusercontent.com/jxnblk/rebass/master/src/Divider.js

import sys from "system-components";

const Divider = sys(
    {
        is: "hr",
        mx: 0,
        my: 3,
        border: 0,
        borderBottom: 1,
        borderColor: "gray"
    },
    "space",
    "color"
);

Divider.displayName = "Divider";

export default Divider;
