import React from "react";
import { Flex } from "grid-styled";

import TabBarItem from "../TabBarItem";

const TabBar = () => {
    return (
        <Flex>
            <TabBarItem
                title="Missions"
                path="/missions"
                glyph="medicalCross"
            />
            <TabBarItem
                title="Aircraft"
                glyph="airplaneFlight"
                path="/aircraft"
            />
            <TabBarItem
                title="Personnel"
                glyph="accountGroup"
                path="/personnel"
            />
        </Flex>
    );
};

export default TabBar;
