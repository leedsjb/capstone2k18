import React from "react";
import { Flex } from "grid-styled";

import TabBarItem from "../TabBarItem";

const TabBar = () => {
    return (
        <Flex>
            <TabBarItem title="Missions" />
            <TabBarItem title="Aircraft" />
            <TabBarItem title="Personnel" />
        </Flex>
    );
};

export default TabBar;