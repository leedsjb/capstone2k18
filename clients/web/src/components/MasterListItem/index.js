import React from "react";
import StyledBox from "./StyledBox";

const MasterListItem = ({ children }) => {
    return <StyledBox p={3}>{children}</StyledBox>;
};

export default MasterListItem;
