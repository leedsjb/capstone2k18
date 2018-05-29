// Adaptation of https://github.com/jxnblk/rebass/blob/master/src/Border.js

import sys from "system-components";

const Border = sys(
    {
        border: 1,
        borderColor: "border"
    },
    "space",
    "width",
    "color"
);

Border.displayName = "Border";

export default Border;
