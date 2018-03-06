import React from "react";
import Box from "../Box";

const Toolbar = props => {
    return (
        <Box py={1} px={2} borderBottom="1px solid">
            {props.children}
        </Box>
    );
};

export default Toolbar;
