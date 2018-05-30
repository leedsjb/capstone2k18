// Adaptation of https://raw.githubusercontent.com/jxnblk/rebass/master/src/Label.js

import sys from "system-components";

const Label = sys(
    {
        is: "label",
        fontSize: 2,
        mt: 6,
        mb: 3,
        align: "center",
        fontWeight: "bold"
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
