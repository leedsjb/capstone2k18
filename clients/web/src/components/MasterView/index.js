import React from "react";
import { Flex } from "grid-styled";

import Box from "../Box";

const MasterView = ({ children }) => {
    return (
        <Box
            w={[1, 1, 1 / 2]}
            maxWidth={[null, null, 400]}
            style={{
                overflowY: "hidden",
                display: "flex",
                flexDirection: "column",
                flex: 1
            }}
        >
            {children}
        </Box>
    );
};

export default MasterView;
