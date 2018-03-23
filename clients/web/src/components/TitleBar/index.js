import React, { Component } from "react";
import { Flex } from "grid-styled";

import Box from "../Box";
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

    render() {
        return (
            <Box bg="wireframe" py={3} px={3}>
                <Flex justifyContent="space-between" align="center">
                    <Box size={32} bg="black" onClick={this.toggleDrawer} />
                    <Heading is="h3" fontSize={2}>
                        {this.props.title}
                    </Heading>
                    <Box size={32} bg="black" />
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
