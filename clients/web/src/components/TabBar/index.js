import React from "react";
import { Flex } from "grid-styled";
import Media from "react-media";
import { withTheme } from "styled-components";

import TabBarItem from "../TabBarItem";
import GradientBox from "../GradientBox";

const TabBar = ({ theme: { breakpoints } }) => {
    return (
        <Media
            query={`(max-width: ${breakpoints[1]})`}
            render={() => {
                return (
                    <GradientBox>
                        <Flex>
                            <TabBarItem
                                title="Aircraft"
                                glyph="airplaneFlight"
                                path="/aircraft"
                            />
                            <TabBarItem
                                title="People"
                                glyph="accountGroup"
                                path="/people"
                            />
                            <TabBarItem
                                title="Resources"
                                glyph="accountGroup"
                                path="/resources"
                            />
                            <TabBarItem
                                title="Profile"
                                glyph="accountGroup"
                                path="/profile"
                            />
                        </Flex>
                    </GradientBox>
                );
            }}
        />
    );
};

export default withTheme(TabBar);
