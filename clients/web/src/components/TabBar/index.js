import React from "react";
import { Flex } from "grid-styled";

import TabBarItem from "../TabBarItem";

const TabBar = () => {
    return (
        <Flex>
            <TabBarItem title="Missions" glyph="chevronLeft" path="/missions" />
            <TabBarItem
                title="Aircraft"
                glyph="navigationDrawerFilled"
                path="/aircraft"
            />
            <TabBarItem title="Personnel" path="/personnel" />
        </Flex>
    );
};

export default TabBar;
