// Adaptation of https://raw.githubusercontent.com/jxnblk/rebass/master/src/Input.js

import sys from "system-components";
import { theme } from "styled-system";

const TextInput = sys(
    {
        is: "input",
        type: "text",
        fontSize: "inherit",
        lineHeight: "inherit",
        py: 2,
        m: 0,
        width: 1,
        border: 0,
        color: "inherit",
        bg: "transparent"
    },
    "space",
    props => ({
        fontFamily: "inherit",
        display: "inline-block",
        verticalAlign: "middle",
        border: 0,
        appearance: "none",
        "&:focus": {
            outline: "none",
            boxShadow: `inset 0 0 0 1px ${theme("colors.blue")(props)}`
        },
        "&:disabled": {
            opacity: 1 / 2
        }
    })
);

TextInput.displayName = "TextInput";

export default TextInput;
