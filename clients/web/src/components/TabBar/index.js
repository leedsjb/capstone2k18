import React from "react";
import { Flex } from "grid-styled";

import TabBarItem from "../TabBarItem";

import RouterProvider from "../../containers/RouterProvider";

const TabBar = () => {
    return (
        <RouterProvider
            render={({ push }) => {
                return (
                    <Flex>
                        <TabBarItem
                            title="Missions"
                            onClick={() => push("/missions")}
                        />
                        <TabBarItem
                            title="Aircraft"
                            onClick={() => push("/aircraft")}
                        />
                        <TabBarItem
                            title="Personnel"
                            onClick={() => push("/personnel")}
                        />
                    </Flex>
                );
            }}
        />
    );
};

export default TabBar;
