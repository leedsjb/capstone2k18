// Adaptation of https://raw.githubusercontent.com/jxnblk/rebass/master/src/Label.js

import sys from "system-components";

const Label = sys(
    {
        is: "label",
        fontSize: 1,
        mb: 1,
        align: "center"
    },
    {
        display: "flex"
    },
    "alignItems",
    "space",
    "color"
);

Label.displayName = "Label";

export default Label;