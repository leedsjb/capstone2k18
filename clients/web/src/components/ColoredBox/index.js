import React from "react";

import Box from "../Box";

const ColoredBox = ({ word, ...props }) => {
    return <Box bg="blue" {...props} />;
};

export default ColoredBox;
