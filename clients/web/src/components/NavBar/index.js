import React from "react";
import Media from "react-media";
import { Flex } from "grid-styled";
import { Link } from "react-router-dom";
import { withTheme } from "styled-components";

import Box from "../Box";
import Icon from "../Icon";
import ProfileAvatar from "../ProfileAvatar";
import NavBarItem from "../NavBarItem";

const NavBar = ({ theme: { breakpoints } }) => {
    return (
        <Media query={`(min-width: ${breakpoints[1]})`}>
            {matches =>
                matches ? (
                    <Box bg="wireframe" py={3} px={3}>
                        <Flex
                            justifyContent="space-between"
                            alignItems="center"
                        >
                            <Flex alignItems="center">
                                <Link to="/missions">AirliftNW Elevate</Link>
                                <Flex ml={4}>
                                    <NavBarItem
                                        title="Missions"
                                        path="/missions"
                                        glyph="medicalCross"
                                    />
                                    <Box ml={3}>
                                        <NavBarItem
                                            title="Aircraft"
                                            path="/aircraft"
                                            glyph="airplaneFlight"
                                        />
                                    </Box>
                                    <Box ml={3}>
                                        <NavBarItem
                                            title="People"
                                            path="/people"
                                            glyph="accountGroup"
                                        />
                                    </Box>
                                </Flex>
                            </Flex>
                            <Flex alignItems="center">
                                <Icon glyph="grid" />
                                <ProfileAvatar />
                            </Flex>
                        </Flex>
                    </Box>
                ) : null
            }
        </Media>
    );
};

export default withTheme(NavBar);
