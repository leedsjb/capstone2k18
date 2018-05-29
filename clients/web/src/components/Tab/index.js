// Adaptation of https://raw.githubusercontent.com/jxnblk/rebass/master/src/Tab.js

import sys from "system-components";

const Tab = sys(
    {
        is: "a",
        fontSize: 1,
        fontWeight: "bold",
        py: 3,
        flex: 1,
        borderBottom: 2,
        textAlign: "center",
        hover: {
            color: "purple"
        },
        display: "block"
    },
    {
        textDecoration: "none"
    },
    props => ({
        borderColor: props.active ? props.theme.colors.black1 : "transparent",
        color: props.active
            ? props.theme.colors.black1
            : props.theme.colors.gray3
    })
);

Tab.displayName = "Tab";

export default Tab;
