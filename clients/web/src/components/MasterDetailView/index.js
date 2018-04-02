import React from "react";
import { Flex } from "grid-styled";

const MasterDetailView = ({ children }) => {
    return <Flex style={{ flex: 1, overflowY: "hidden" }}>{children}</Flex>;
};

export default MasterDetailView;
