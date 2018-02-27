// Adaptation of https://raw.githubusercontent.com/jxnblk/rebass/master/src/Tabs.js

import sys from "system-components";
import { Flex } from "grid-styled";

const Tabs = sys({
    is: Flex,
    borderBottom: 1,
    borderColor: "gray"
});

Tabs.displayName = "Tabs";

export default Tabs;
