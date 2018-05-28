import sys from "system-components";
import { themeGet } from "styled-system";

export const Span = sys(
    {
        is: "span",
        m: 0
    },
    "space",
    "color",
    "fontSize",
    "fontWeight",
    "textAlign",
    "lineHeight",
    "display",
    props => ({
        colors: themeGet("colors.black1")(props)
    })
);

Span.displayName = "Span";

export default Span;
