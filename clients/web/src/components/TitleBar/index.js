import React, { Component } from "react";
import { Flex } from "grid-styled";
import { Link } from "react-router-dom";
import Media from "react-media";
import { withTheme } from "styled-components";

import Box from "../Box";
import Icon from "../Icon";
import Heading from "../Heading";
import ProfileAvatar from "../ProfileAvatar";

class TitleBar extends Component {
    renderIconLeft = () => {
        if (this.props.back) {
            return (
                <Link to={this.props.backPath}>
                    <Icon glyph="chevronLeft" size={16} />
                </Link>
            );
        }

        return <Box w={24} />;
    };

    renderIconRight = () => {
        return (
            <Flex alignItems="center">
                <Icon glyph="grid" size={16} />
                <ProfileAvatar fName="Dave" />
            </Flex>
        );
    };

    render() {
        return (
            <Media query={`(min-width: ${this.props.theme.breakpoints[1]})`}>
                {matches =>
                    matches ? null : (
                        <Box bg="wireframe" p={3}>
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
