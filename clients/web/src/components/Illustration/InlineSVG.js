import sys from "system-components";

const InlineSVG = sys({
    is: "svg",
    position: "absolute",
    top: 0,
    right: 0,
    bottom: 0,
    left: 0,
    size: "100%",
    color: "inherit",
    fill: "red"
});

InlineSVG.displayName = "InlineSVG";

export default InlineSVG;
