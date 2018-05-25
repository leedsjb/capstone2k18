// Adaptation of https://raw.githubusercontent.com/jxnblk/rebass/master/src/Container.js

import sys from "system-components";
import { Box } from "grid-styled";

const Container = sys(
    {
        is: Box,
        px: 4,
        mx: "auto",
        maxWidth: 1024,
        w: 1
    },
    "maxWidth"
);

Container.displayName = "Container";

export default Container;
