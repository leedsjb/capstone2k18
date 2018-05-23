// Adaptation of https://raw.githubusercontent.com/jxnblk/rebass/master/src/Truncate.js

import sys from "system-components";
import Text from "../Text";

const Truncate = sys(
    {
        is: Text
    },
    {
        overflow: "hidden",
        whiteSpace: "nowrap",
        textOverflow: "ellipsis"
    }
);

Truncate.displayName = "Truncate";

export default Truncate;
