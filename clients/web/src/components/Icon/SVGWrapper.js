import sys from "system-components";

const SVGWrapper = sys(
    {
        display: "inline-block",
        size: 24,
        minWidth: 24,
        minHeight: 24,
        position: "relative"
    },
    "size"
);

SVGWrapper.displayName = "SVGWrapper";

export default SVGWrapper;
