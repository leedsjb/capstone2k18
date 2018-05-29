import React, { Component } from "react";
import { Flex } from "grid-styled";
import { withTheme } from "styled-components";
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

// TODO: implement alt path
class TabBarItem extends Component {
    render() {
        const { title, glyph, path, pathSecond } = this.props;

        return (
            <RouterProvider
                render={({ router: { location } }) => {
                    const { pathname } = location;

                    return (
                        <Box
                            flex={1}
                            borderTop={`1px solid ${
                                this.props.theme.colors.gray4
                            }`}
                            py={1}
                        >
                            <Link is={NavLink} to={path}>
                                <Flex
                                    flexDirection="column"
                                    justifyContent="center"
                                    alignItems="center"
                                >
                                    <Icon
                                        glyph={
                                            matchPath(pathname, path) ||
                                            (pathSecond &&
                                                matchPath(pathname, pathSecond))
                                                ? `${glyph}Filled`
                                                : `${glyph}Line`
                                        }
                                        size={24}
                                        color={
                                            matchPath(pathname, path) ||
                                            (pathSecond &&
                                                matchPath(pathname, pathSecond))
                                                ? "airlift1"
                                                : "black3"
                                        }
                                    />
                                    <Span
                                        color={
                                            matchPath(pathname, path) ||
                                            (pathSecond &&
                                                matchPath(pathname, pathSecond))
                                                ? "airlift"
                                                : "black3"
                                        }
                                    >
                                        {title}
                                    </Span>
                                </Flex>
                            </Link>
                        </Box>
                    );
                }}
            />
        );
    }
}

export default withTheme(TabBarItem);
