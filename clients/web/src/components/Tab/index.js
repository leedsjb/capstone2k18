// Adaptation of https://raw.githubusercontent.com/jxnblk/rebass/master/src/Tab.js

import sys from "system-components";

const Tab = sys(
    {
        is: "a",
        fontSize: 1,
        fontWeight: "bold",
        mr: 3,
        py: 2,
        color: "inherit",
        borderBottom: 2,
        borderColor: "transparent",
        hover: {
            color: "blue"
        }
    },
    {
        textDecoration: "none"
    }
);

Tab.displayName = "Tab";

export default Tab;
