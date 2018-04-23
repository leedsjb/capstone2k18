// Adaptation of https://raw.githubusercontent.com/jxnblk/rebass/master/src/Tab.js

import sys from "system-components";

const Tab = sys(
    {
        is: "a",
        fontSize: 1,
        fontWeight: "bold",
        py: 2,
        color: "inherit",
        flex: 1,
        borderBottom: 2,
        textAlign: "center",
        hover: {
            color: "blue"
        },
        display: "block"
    },
    {
        textDecoration: "none"
    },
    props => ({
        borderColor: props.active ? "blue" : "transparent"
    })
);

Tab.displayName = "Tab";

export default Tab;
