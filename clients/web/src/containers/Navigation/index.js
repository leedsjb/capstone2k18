import React, { Component } from "react";
import { Flex, Box } from "grid-styled";
import { Link as RouterLink } from "react-router-dom";

import Link from "../../components/Link";
import Heading from "../../components/Heading";
import Icon from "../../components/Icon";

class Navigation extends Component {
    render() {
        return (
            <Box bg="wireframe" px={2} py={3}>
                <Flex justifyContent="center">
                    <Heading is="h3" children="Test" fontSize={2} />
                    <Icon glyph="facebook" size={24} />
                </Flex>
            </Box>
        );
    }
}

export default Navigation;
