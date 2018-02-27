// Adaptation of https://github.com/jxnblk/rebass/blob/master/src/Measure.js

import sys from "system-components";
import Text from "./Text";

export const Measure = sys(
    {
        is: Text,
        maxWidth: "30em"
    },
    "maxWidth"
);

Measure.displayName = "Measure";

export default Measure;
