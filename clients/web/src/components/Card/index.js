// Adaptation of https://raw.githubusercontent.com/jxnblk/rebass/master/src/Card.js

import sys from "system-components";

const Card = sys(
    {
        p: 2,
        bg: "white",
        borderRadius: 2,
        boxShadow: 2
    },
    {
        overflow: "hidden"
    },
    "space",
    "color"
);

Card.displayName = "Card";

export default Card;