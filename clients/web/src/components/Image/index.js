// Adaptation of https://raw.githubusercontent.com/jxnblk/rebass/master/src/Image.js

import sys from "system-components";

const Image = sys(
    {
        is: "img"
    },
    {
        display: "block",
        maxWidth: "100%",
        height: "auto"
    },
    "space",
    "width",
    "color"
);

Image.displayName = "Image";

export default Image;
