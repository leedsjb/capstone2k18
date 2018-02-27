// Adaptation of https://raw.githubusercontent.com/jxnblk/rebass/master/src/Radio.js

import sys from "system-components";

const Radio = sys(
    {
        is: "input",
        type: "radio",
        m: 0,
        mr: 2
    },
    "space",
    "color"
);

Radio.displayName = "Radio";

export default Radio;
