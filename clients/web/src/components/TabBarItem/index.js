import React from "react";
import { Flex } from "grid-styled";
import { NavLink } from "react-router-dom";

import Box from "../Box";
import Icon from "../Icon";
import Link from "../Link";

import RouterProvider from "../../containers/RouterProvider";

import matchPath from "../../utils/matchPath";

// TODO: Consider cleaning up this component and
// revisiting the technique used to detect
// the active route
const TabBarItem = ({ title, glyph, path }) => {
    return (
        <RouterProvider
            render={({ router: { location } }) => {
                const { pathname } = location;

                return (
                    <Box bg="wireframe" flex={1}>
                        <Link is={NavLink} to={path}>
                            <Flex
                                flexDirection="column"
                                justifyContent="center"
                                alignItems="center"
                            >
                                <Icon
                                    glyph={
                                        matchPath(pathname, path)
                                            ? `${glyph}Filled`
                                            : `${glyph}Line`
                                    }
                                />
                                {title}
                            </Flex>
                        </Link>
                    </Box>
                );
            }}
        />
    );
};

export default TabBarItem;
