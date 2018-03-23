import React from "react";
import { Flex } from "grid-styled";
import { NavLink } from "react-router-dom";

import Box from "../Box";
import Icon from "../Icon";

const TabBarItem = ({ title, glyph, path }) => {
    return (
        <Box bg="wireframe" flex={1}>
            <NavLink to={path}>
                <Flex
                    flexDirection="column"
                    justifyContent="center"
                    alignItems="center"
                >
                    <Icon glyph={glyph} />
                    {title}
                </Flex>
            </NavLink>
        </Box>
    );
};

export default TabBarItem;
