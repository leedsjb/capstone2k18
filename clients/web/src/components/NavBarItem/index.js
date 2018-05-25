import React from "react";
import { Flex } from "grid-styled";
import { NavLink } from "react-router-dom";

import Box from "../Box";
import Icon from "../Icon";
import Link from "../Link";
import Span from "../Span";

import RouterProvider from "../../containers/RouterProvider";

import matchPath from "../../utils/matchPath";

// TODO: Consider cleaning up this component and
// revisiting the technique used to detect
// the active route
const NavBarItem = ({ title, glyph, path }) => {
    return (
        <RouterProvider
            render={({ router: { location } }) => {
                const { pathname } = location;

                return (
                    <Link is={NavLink} to={path}>
                        <Flex alignItems="center">
                            <Icon
                                glyph={
                                    matchPath(pathname, path)
                                        ? `${glyph}Filled`
                                        : `${glyph}Line`
                                }
                                size={24}
                                color="black"
                            />
                            <Box ml={3}>
                                <Span color="black">{title}</Span>
                            </Box>
                        </Flex>
                    </Link>
                );
            }}
        />
    );
};

export default NavBarItem;
