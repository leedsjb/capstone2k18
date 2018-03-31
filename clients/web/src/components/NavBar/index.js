import React from "react";
import Media from "react-media";
import { Flex } from "grid-styled";
import { withTheme } from "styled-components";

import Box from "../Box";
import Icon from "../Icon";
import ProfileAvatar from "../ProfileAvatar";

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
                            <Flex>
                                AirliftNW Elevate
                                <Flex ml={4}>Missions Aircraft People</Flex>
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
