// Adaptation of https://github.com/jxnblk/rebass/blob/master/src/Border.js

import sys from "system-components";

export const Border = sys(
    {
        border: 1,
        borderColor: "gray"
    },
    "space",
    "width",
    "color"
);

Border.displayName = "Border";

export default Border;
