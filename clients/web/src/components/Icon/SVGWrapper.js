import sys from "system-components";

const SVGWrapper = sys(
    {
        display: "inline-block",
        size: 24,
        position: "relative"
    },
    "size",
    "color"
);

SVGWrapper.displayName = "SVGWrapper";

export default SVGWrapper;
