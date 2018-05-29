// Adaptation of https://raw.githubusercontent.com/jxnblk/rebass/master/src/Input.js

import sys from "system-components";

const TextInput = sys(
    {
        is: "input",
        type: "text",
        fontSize: "inherit",
        lineHeight: "inherit",
        py: 2,
        m: 0,
        width: 1,
        color: "inherit",
        bg: "transparent"
    },
    "space",
    "borders",
    "borderRadius",
    props => ({
        fontFamily: "inherit",
        display: "inline-block",
        verticalAlign: "middle",
        appearance: "none",
        "&:focus": {
            outline: "none"
        },
        "&:disabled": {
            opacity: 3 / 4
        }
    })
);

TextInput.displayName = "TextInput";

export default TextInput;
