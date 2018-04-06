// Adaptation of https://raw.githubusercontent.com/jxnblk/rebass/master/src/Embed.js

import sys from "system-components";

const Embed = sys(
    {
        ratio: 9 / 16
    },
    "ratio",
    "space",
    "color",
    props => ({
        position: "relative",
        overflow: "hidden",
        "& > iframe": {
            position: "absolute",
            width: "100%",
            height: "100%",
            top: 0,
            bottom: 0,
            left: 0,
            border: 0
        }
    })
);

Embed.displayName = "Embed";

export default Embed;