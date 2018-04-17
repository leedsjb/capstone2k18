import React from "react";
import { Flex } from "grid-styled";
import { NavLink } from "react-router-dom";

import Box from "../Box";
import Icon from "../Icon";
import Link from "../Link";
import Tab from "../../components/Tab";

import RouterProvider from "../../containers/RouterProvider";

import matchPath from "../../utils/matchPath";

// TODO: Consider cleaning up this component and
// revisiting the technique used to detect
// the active route
const NavBarItem = ({ title, path }) => {
    return (
        <RouterProvider
            render={({ router: { location } }) => {
                const { pathname } = location;

                return (
                    <Link is={NavLink} to={path}>
                        <Tab active={matchPath(pathname, path) ? true : false}>
                            {title}
                        </Tab>
                    </Link>
                );
            }}
        />
    );
};

export default NavBarItem;
