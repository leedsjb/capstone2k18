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
        hover: {
            color: "blue"
        }
    },
    {
        textDecoration: "none",
        display: "inline-block"
    },
    props => ({
        borderColor: props.active ? "blue" : "transparent"
    })
);

Tab.displayName = "Tab";

export default Tab;
