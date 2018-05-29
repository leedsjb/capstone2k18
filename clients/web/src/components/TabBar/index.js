import React from "react";
import { Flex } from "grid-styled";
import Media from "react-media";
import { withTheme } from "styled-components";

import TabBarItem from "../TabBarItem";
import Box from "../Box";

const TabBar = ({ theme: { colors, breakpoints } }) => {
    return (
        <Media
            query={`(max-width: ${breakpoints[1]})`}
            render={() => {
                return (
                    <Box bg="white" borderTop={`1px solid ${colors.border}`}>
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
                                pathSecond="/groups"
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
                    </Box>
                );
            }}
        />
    );
};

export default withTheme(TabBar);
