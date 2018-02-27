// Adaptation of https://raw.githubusercontent.com/jxnblk/rebass/master/src/Container.js

import sys from "system-components";
import { Box } from "grid-styled";

export const Container = sys(
    {
        is: Box,
        px: 3,
        mx: "auto",
        maxWidth: 1024
    },
    "maxWidth"
);

Container.displayName = "Container";

export default Container;
