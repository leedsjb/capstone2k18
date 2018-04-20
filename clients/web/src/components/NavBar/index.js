import React from "react";
import Media from "react-media";
import { Flex } from "grid-styled";
import { Link } from "react-router-dom";
import { withTheme } from "styled-components";

import Box from "../Box";
import Logo from "../Logo";
import ProfileDropdown from "../ProfileDropdown";
import ResourcesDropdown from "../ResourcesDropdown";

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
