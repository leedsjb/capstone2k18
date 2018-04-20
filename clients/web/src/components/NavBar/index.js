import React from "react";
import Media from "react-media";
import { Flex } from "grid-styled";
import { Link } from "react-router-dom";
import { withTheme } from "styled-components";

import Box from "../Box";
import Logo from "../Logo";
import ProfileDropdown from "../ProfileDropdown";
import ResourcesDropdown from "../ResourcesDropdown";
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
                                <Link to="/aircraft">
                                    <Logo />
                                </Link>
                                <Box ml={4}>
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
                            <Flex alignItems="center">
                                <ProfileDropdown />
                                <ResourcesDropdown />
                            </Flex>
                        </Flex>
                    </Box>
                ) : null
            }
        </Media>
    );
};

export default withTheme(NavBar);
