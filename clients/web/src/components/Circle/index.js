// Adaptation of https://raw.githubusercontent.com/jxnblk/rebass/master/src/Circle.js

import sys from "system-components";
import Badge from "../Badge";

const Circle = sys(
    {
        is: Badge,
        size: 24,
        align: "center",
        borderRadius: "99999px"
    },
    "textAlign",
    "size",
    "space"
);

Circle.displayName = "Circle";

export default Circle;
