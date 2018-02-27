// Adaptation of https://raw.githubusercontent.com/jxnblk/rebass/master/src/Link.js

import sys from "system-components";

const Link = sys(
    {
        is: "a",
        color: "blue"
    },
    "space"
);

Link.displayName = "Link";

export default Link;
