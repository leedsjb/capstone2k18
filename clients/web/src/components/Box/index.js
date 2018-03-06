import sys from "system-components";
import { Box as GSBox } from "grid-styled";

const Box = sys(
    {
        is: GSBox
    },
    "borders"
);

Box.displayName = "Box";

export default Box;
