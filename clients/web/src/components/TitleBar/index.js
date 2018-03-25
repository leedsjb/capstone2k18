import React, { Component } from "react";
import { Flex } from "grid-styled";
import { Link } from "react-router-dom";

import Box from "../Box";
import Icon from "../Icon";
import Heading from "../Heading";
import Fixed from "../Fixed";
import Drawer from "../Drawer";

class TitleBar extends Component {
    constructor(props) {
        super(props);
        this.state = {
            drawer: false
        };
    }

    toggleDrawer = () => {
        this.setState({
            drawer: !this.state.drawer
        });
    };

    renderIconLeft = () => {
        if (this.props.back) {
            return (
                <Link to={this.props.backPath}>
                    <Icon glyph="chevronLeft" />
                </Link>
            );
        }

        return (
            <Icon glyph="navigationDrawerFilled" onClick={this.toggleDrawer} />
        );
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
                {this.state.drawer && (
                    <Fixed
                        top={0}
                        right={0}
                        bottom={0}
                        left={0}
                        onClick={this.toggleDrawer}
                    />
                )}
                <Drawer
                    open={this.state.drawer}
                    position="left"
                    p={3}
                    color="white"
                    bg="black"
                >
                    TODO: Change this drawer component
                </Drawer>
            </Box>
        );
    }
}

export default TitleBar;
