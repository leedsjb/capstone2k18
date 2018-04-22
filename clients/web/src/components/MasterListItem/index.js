import React from "react";
import StyledBox from "./StyledBox";

const MasterListItem = ({ children, active }) => {
    console.log("WAHOO", active);
    return (
        <StyledBox p={3} active={active}>
            {children}
        </StyledBox>
    );
};

export default MasterListItem;
