import React, { Component } from "react";
import { Flex } from "grid-styled";
import { Link } from "react-router-dom";
import Media from "react-media";
import { withTheme } from "styled-components";

import Box from "../Box";
import Icon from "../Icon";
import Heading from "../Heading";
import Fixed from "../Fixed";
import Circle from "../Circle";
import Text from "../Text";

class TitleBar extends Component {
    renderIconLeft = () => {
        if (this.props.back) {
            return (
                <Link to={this.props.backPath}>
                    <Icon glyph="chevronLeft" />
                </Link>
            );
        }

        return <Box w={24} />;
    };

    renderIconRight = () => {
        return (
            <Flex>
                <Icon glyph="grid" />
                <Circle size={32} p={0}>
                    <Flex
                        flexDirection="column"
                        alignItems="center"
                        justifyContent="center"
                    >
                        <Text>V</Text>
                    </Flex>
                </Circle>
            </Flex>
        );
    };

    render() {
        return (
            <Media query={`(min-width: ${this.props.theme.breakpoints[1]})`}>
                {matches =>
                    matches ? null : (
                        <Box bg="wireframe" py={3} px={3}>
                            <Flex justifyContent="space-between" align="center">
                                {this.renderIconLeft()}
                                <Heading is="h3" fontSize={2}>
                                    {this.props.title}
                                </Heading>
                                {this.renderIconRight()}
                            </Flex>
                        </Box>
                    )
                }
            </Media>
        );
    }
}

export default withTheme(TitleBar);
