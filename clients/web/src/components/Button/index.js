// Adaptation of https://raw.githubusercontent.com/jxnblk/rebass/master/src/Button.js

import sys from "system-components";
import { themeGet } from "styled-system";

const Button = sys(
    {
        is: "button",
        fontSize: 1,
        fontWeight: "bold",
        lineHeight: 16 / 14,
        m: 0,
        px: 3,
        py: 2,
        color: "white",
        borderRadius: 2,
        border: 0
    },
    props => ({
        fontFamily: "inherit",
        WebkitFontSmoothing: "antialiased",
        display: "inline-block",
        verticalAlign: "middle",
        textAlign: "center",
        textDecoration: "none",
        appearance: "none",
        backgroundColor: `${themeGet("colors.black2")(props)}`,
        "&:hover": {
            cursor: "pointer"
        },
        "&:focus": {
            outline: 0
        },
        "&:active": {},
        "&:disabled": {
            opacity: 1 / 4
        }
    })
);

Button.displayName = "Button";

export default Button;
