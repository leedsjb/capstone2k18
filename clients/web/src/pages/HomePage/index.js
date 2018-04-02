import React from "react";
import { Flex } from "grid-styled";

import Box from "../../components/Box";
import Heading from "../../components/Heading";
import Container from "../../components/Container";
import Logo from "../../components/Logo";
import Measure from "../../components/Measure";

const HomePage = () => {
    return (
        <div>
            <Box p={3}>
                <Flex justifyContent="space-between">
                    <Logo />
                    <span>Sign in</span>
                </Flex>
            </Box>
            <Container>
                <Heading is="h1">AirliftNW Elevate</Heading>

                <Heading is="h2" fontSize={4}>
                    Feature One
                </Heading>
                <Measure>
                    Donec sed odio dui. Vivamus sagittis lacus vel augue laoreet
                    rutrum faucibus dolor auctor. Integer posuere erat a ante
                    venenatis dapibus posuere velit aliquet. Sed posuere
                    consectetur est at lobortis.
                </Measure>

                <Heading is="h2" fontSize={4}>
                    Feature Two
                </Heading>
                <Measure>
                    Maecenas sed diam eget risus varius blandit sit amet non
                    magna. Nulla vitae elit libero, a pharetra augue. Donec
                    ullamcorper nulla non metus auctor fringilla. Nulla vitae
                    elit libero, a pharetra augue.
                </Measure>
            </Container>
        </div>
    );
};

export default HomePage;
