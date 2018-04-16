import sys from "system-components";

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
    "lineHeight"
);

Span.displayName = "Span";

export default Span;
