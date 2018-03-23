import React from "react";

import Box from "../../components/Box";

import RouterProvider from "../../containers/RouterProvider";

const MissionListItem = () => {
    return (
        <RouterProvider
            render={({ push }) => {
                return (
                    <Box py={3} onClick={() => push("/missions/test")}>
                        Mission List item
                    </Box>
                );
            }}
        />
    );
};

export default MissionListItem;
