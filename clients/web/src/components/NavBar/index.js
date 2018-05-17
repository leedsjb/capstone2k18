import React from "react";
import Media from "react-media";
import { Flex } from "grid-styled";
import { Link } from "react-router-dom";
import { withTheme } from "styled-components";

import Box from "../Box";
import ProfileDropdown from "../ProfileDropdown";
import ResourcesDropdown from "../ResourcesDropdown";
import NavBarItem from "../NavBarItem";
import GradientBox from "./GradientBox";
import Image from "../Image";

import logo from "../../images/logo.svg";

const NavBar = ({ theme: { breakpoints } }) => {
    return (
        <Media query={`(min-width: ${breakpoints[1]})`}>
            {matches =>
                matches ? (
                    <GradientBox py={3} px={4}>
                        <Flex
                            justifyContent="space-between"
                            alignItems="center"
                        >
                            <Flex alignItems="center">
                                <Link to="/aircraft">
                                    <Image src={logo} alt="Elevate" w={49} />
                                </Link>
                                <Box ml={8}>
                                    <NavBarItem
                                        title="Aircraft"
                                        path="/aircraft"
                                        glyph="airplaneFlight"
                                    />
                                </Box>
                                <Box ml={6}>
                                    <NavBarItem
                                        title="People"
                                        path="/people"
                                        glyph="accountGroup"
                                    />
                                </Box>
                            </Flex>
                            <Flex alignItems="center">
                                <ResourcesDropdown />
                                <Box ml={3}>
                                    <ProfileDropdown />
                                </Box>
                            </Flex>
                        </Flex>
                    </GradientBox>
                ) : null
            }
        </Media>
    );
};

export default withTheme(NavBar);
