import React from "react";

import RouterProvider from "../../containers/RouterProvider";

const MissionListItem = () => {
    return (
        <RouterProvider
            render={({ push }) => {
                return (
                    <div onClick={() => push("/missions/test")}>
                        Mission List item
                    </div>
                );
            }}
        />
    );
};

export default MissionListItem;
