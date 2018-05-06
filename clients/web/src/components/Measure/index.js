// Adaptation of https://github.com/jxnblk/rebass/blob/master/src/Measure.js

import sys from "system-components";

const Measure = sys(
    {
        is: "p",
        maxWidth: "30em"
    },
    "maxWidth",
    "space",
    "color",
    "fontSize",
    "fontWeight",
    "textAlign",
    "lineHeight"
);

Measure.displayName = "Measure";

export default Measure;
