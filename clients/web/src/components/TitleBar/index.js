import React, { Component } from "react";
import { Flex } from "grid-styled";
import { Link } from "react-router-dom";

import Box from "../Box";
import Icon from "../Icon";
import Heading from "../Heading";
import Fixed from "../Fixed";

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
        if (this.props.icon) {
            return (
                <Icon
                    glyph={this.props.icon}
                    onClick={this.props.iconOnClick}
                />
            );
        }

        return <Box w={32} />;
    };

    render() {
        return (
            <Box bg="wireframe" py={3} px={3}>
                <Flex justifyContent="space-between" align="center">
                    {this.renderIconLeft()}
                    <Heading is="h3" fontSize={2}>
                        {this.props.title}
                    </Heading>
                    {this.renderIconRight()}
                </Flex>
            </Box>
        );
    }
}

export default TitleBar;
