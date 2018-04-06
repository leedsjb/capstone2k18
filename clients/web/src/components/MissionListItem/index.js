import React from "react";
import { Link } from "react-router-dom";

import Box from "../Box";

const MissionListItem = () => {
    return (
        <Box py={2}>
            <Link to="/missions/test">Mission List item</Link>
            <Link to="/missions/4">Mission List item</Link>
        </Box>
    );
};

export default MissionListItem;
