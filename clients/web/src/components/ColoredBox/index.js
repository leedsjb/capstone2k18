import React from "react";

import GradientBox from "../GradientBox";
import Box from "../Box";

const ColoredBox = ({ word, ...props }) => {
    return <Box bg="blue" {...props} />;
};

export default ColoredBox;
