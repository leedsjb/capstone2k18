// Adaptation of https://raw.githubusercontent.com/jxnblk/rebass/master/src/Card.js

import sys from "system-components";
import Box from "../Box";

const Card = sys(
    {
        bg: "white",
        borderRadius: 4,
        border: "1px solid",
        borderColor: "border",
        is: Box
    },
    {
        overflow: "hidden"
    },
    "space",
    "color",
    "height"
);

Card.displayName = "Card";

export default Card;
