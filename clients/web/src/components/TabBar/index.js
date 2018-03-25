import React from "react";
import { Flex } from "grid-styled";
import Media from "react-media";
import { withTheme } from "styled-components";

import TabBarItem from "../TabBarItem";

const TabBar = ({ theme: { breakpoints } }) => {
    return (
        <Media
            query={`(max-width: ${breakpoints[1]})`}
            render={() => {
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
            }}
        />
    );
};

export default withTheme(TabBar);
