import React, { Component } from "react";
import { Flex } from "grid-styled";
import { Link } from "react-router-dom";
import Media from "react-media";
import { withTheme } from "styled-components";

import Box from "../Box";
import Icon from "../Icon";
import Heading from "../Heading";

class TitleBar extends Component {
    renderIconLeft = () => {
        if (this.props.back) {
            return (
                <Link to={this.props.backPath}>
                    <Icon glyph="chevronLeft" size={16} color="black" />
                </Link>
            );
        }
        return <Box w={16} />;
    };

    renderMapIcon = () => {
        return (
            <Flex alignItems="center">
                <Icon glyph="map" size={16} color="black" />
            </Flex>
        );
    };

    render() {
        return (
            <Media query={`(min-width: ${this.props.theme.breakpoints[1]})`}>
                {matches =>
                    matches ? null : this.props.showMap ? (
                        <Box bg="white" p={3} borderBottom="1px solid gray">
                            <Flex justifyContent="flex-start" align="center">
                                <Box width={1 / 3}>
                                    <Link to={this.props.link}>
                                        {this.renderMapIcon()}
                                    </Link>
                                </Box>
                                <Box width={1 / 3}>
                                    <Heading
                                        is="h3"
                                        fontSize={2}
                                        textAlign="center"
                                        color="black"
                                    >
                                        {this.props.title}
                                    </Heading>
                                </Box>
                            </Flex>
                        </Box>
                    ) : (
                        <Box bg="white" p={3} borderBottom="1px solid gray">
                            <Flex justifyContent="space-between" align="center">
                                {this.renderIconLeft()}
                                <Heading is="h3" fontSize={2} color="black">
                                    {this.props.title}
                                </Heading>
                                <Box w={16} />
                            </Flex>
                        </Box>
                    )
                }
            </Media>
        );
    }
}

export default withTheme(TitleBar);
