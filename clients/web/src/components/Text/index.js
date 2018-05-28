// Adaptation of https://raw.githubusercontent.com/jxnblk/rebass/master/src/Text.js

import sys from "system-components";
import { themeGet } from "styled-system";

const Text = sys(
    {
        m: 0
    },
    "space",
    "color",
    "fontSize",
    "fontWeight",
    "textAlign",
    "lineHeight",
    props => ({
        colors: themeGet("colors.black1")(props)
    })
);

Text.displayName = "Text";

export default Text;
