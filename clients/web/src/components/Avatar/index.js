// Adaptation of https://raw.githubusercontent.com/jxnblk/rebass/master/src/Avatar.js

import sys from "system-components";

const Avatar = sys(
    {
        is: "img",
        size: 48,
        borderRadius: "99999px"
    },
    {
        display: "inline-block"
    },
    "space",
    "color",
    "size"
);

Avatar.displayName = "Avatar";

export default Avatar;
