// Adaptation of https://raw.githubusercontent.com/jxnblk/rebass/master/src/Text.js

import sys from "system-components";

export const Text = sys(
    {
        m: 0
    },
    "space",
    "color",
    "fontSize",
    "fontWeight",
    "textAlign",
    "lineHeight"
);

Text.displayName = "Text";

export default Text;
