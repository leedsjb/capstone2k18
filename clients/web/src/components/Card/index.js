// Adaptation of https://raw.githubusercontent.com/jxnblk/rebass/master/src/Card.js

import sys from "system-components";
import Box from "../Box";

const Card = sys(
    {
        p: 4,
        bg: "white",
        borderRadius: 4,
        border: "1px solid",
        borderColor: "gray",
        is: Box
    },
    {
        overflow: "hidden"
    },
    "space",
    "color"
);

Card.displayName = "Card";

export default Card;
