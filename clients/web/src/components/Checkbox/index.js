// Adaptation of https://raw.githubusercontent.com/jxnblk/rebass/master/src/Checkbox.js

import sys from "system-components";

const Checkbox = sys(
    {
        is: "input",
        type: "checkbox",
        m: 0,
        mr: 2
    },
    "space",
    "color"
);

Checkbox.displayName = "Checkbox";

export default Checkbox;
