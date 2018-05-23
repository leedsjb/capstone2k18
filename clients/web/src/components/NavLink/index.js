// Adaptation of https://raw.githubusercontent.com/jxnblk/rebass/master/src/NavLink.js

import sys from "system-components";
import { themeGet } from "styled-system";

const NavLink = sys(
    {
        is: "a",
        color: "inherit",
        bg: "transparent",
        fontSize: 1,
        fontWeight: "bold",
        p: 2
    },
    props => ({
        display: "inline-flex",
        alignItems: "center",
        alignSelf: "stretch",
        textDecoration: "none",
        whiteSpace: "nowrap",
        cursor: "pointer",
        "&:disabled": {
            opacity: 1 / 4
        }
    }),
    "width"
);

NavLink.displayName = "NavLink";

export default NavLink;
